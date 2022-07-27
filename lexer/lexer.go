package lexer

import (
	"github.com/dev001hajipro/json_parser/token"
)

type Lexer struct {
	input        string
	position     int // current position
	readPosition int // next position
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	// set position to first char.
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()

	switch l.ch {
	case ':':
		tok = newToken(token.COLON, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '[':
		tok = newToken(token.LBRACKET, l.ch)
	case ']':
		tok = newToken(token.RBRACKET, l.ch)

	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifer()
			tok.Type = token.LookupIndent(tok.Literal)
			return tok
		} else if isNumber(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.NUMBER
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

// gross specification.
func isNumber(ch byte) bool {
	return ('0' <= ch && ch <= '9') || ch == 'e' || ch == '-' || ch == '.'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func (l *Lexer) readNumber() string {
	pos := l.position
	for isNumber(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func newToken(t token.TokenType, b byte) token.Token {
	return token.Token{Type: t, Literal: string(b)}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) readString() string {
	// when this function called, position is '"'. so, move to next char.
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\b' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readIdentifer() string {
	pos := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

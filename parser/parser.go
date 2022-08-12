package parser

import (
	"fmt"
	"strconv"

	"github.com/dev001hajipro/json_parser/lexer"
	"github.com/dev001hajipro/json_parser/token"
)

type Parser struct {
	l         *lexer.Lexer
	errors    []string
	curToken  token.Token
	peekToken token.Token // one ahead token.
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// call nextToken twice to initialize curToken and peekToken.
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken() // read next token.
}

func (p *Parser) Parse() any {
	for p.curToken.Type != token.EOF {
		switch p.curToken.Type {
		case token.LBRACKET:
			return p.ParseArray(token.RBRACKET)
		case token.LBRACE:
			return p.ParseObject(token.RBRACE)
		case token.STRING:
			return p.curToken.Literal
		case token.NUMBER:
			var err error
			i, err := strconv.ParseInt(p.curToken.Literal, 10, 0)
			if err  == nil {
				return i;
			}

			var f float64
			f, err = strconv.ParseFloat(p.curToken.Literal, 64)
			if err  != nil {
				return err;
			}
			return f
			
		default:
		}

		p.nextToken()
	}
	return nil
}

func (p *Parser) ParseArray(end token.TokenType) []any {
	var list []any

	// if next token is ], then read ].
	if p.peekTokenIs(end) {
		p.nextToken()
		return list
	}

	// read first element.
	p.nextToken()
	list = append(list, p.Parse())

	for p.peekTokenIs(token.COMMA) {
		p.nextToken() // skip comma
		p.nextToken() // cursor position set to target element.
		
		list = append(list, p.Parse())
	}

	// read ]
	if !p.expectPeek(end) {
		return nil
	}
	return list
}

func (p *Parser) ParseObject(end token.TokenType) map[any]any {
	var kv = make(map[any]any)

	for !p.peekTokenIs(token.RBRACE) {
		p.nextToken()
		key := p.Parse()

		if !p.expectPeek(token.COLON) {
			return nil
		}

		p.nextToken() // skip colon

		value := p.Parse()

		kv[key] = value

		if !p.peekTokenIs(end) && !p.expectPeek(token.COMMA) {
			return nil
		}
	}

	if !p.expectPeek(end) {
		return nil
	}

	return kv
}

func (p *Parser)peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t

}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expeceted next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
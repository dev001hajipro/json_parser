package parser

import (
	"fmt"

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

func (p *Parser) Parse() {
	for p.curToken.Type != token.EOF {
		switch p.curToken.Type {
		case token.LBRACKET:
			println("lbracket")
			p.ParseArray(token.RBRACKET)
		case token.LBRACE:
			println("lbrace")
		default:
		}

		p.nextToken()
	}
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
	//list = append(list, p.parseExpression(LOWEST))

	for p.peekTokenIs(token.COMMA) {
		p.nextToken() // skip comma
		p.nextToken() // cursor position set to target element.
	//	list = append(list, p.parseExpression(LOWEST))
	}

	if !p.expectPeek(end) {
		return nil
	}
	return list
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
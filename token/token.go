package token

const (
	COMMA = ","
	COLON = ":"
	LBRACKET = "["
	RBRACKET = "]"
	LBRACE = "{"
	RBRACE = "}"
	QUOTE = "\""

	STRING = "STRING"
	NUMBER = "NUMBER"
)

type TokenType string
type Token struct {
	Type TokenType
	Literal string
}
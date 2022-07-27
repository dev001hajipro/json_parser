package token

const (
	ILLEGAL  = "ILLEGAL"
	EOF      = "EOF"
	IDENT    = "IDENT" // add, foobar, x, y...
	COMMA    = ","
	COLON    = ":"
	LBRACKET = "["
	RBRACKET = "]"
	LBRACE   = "{"
	RBRACE   = "}"
	QUOTE    = "\""
	STRING   = "STRING"
	NUMBER   = "NUMBER"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	NULL	 = "NULL"
)

type TokenType string
type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"true":  TRUE,
	"false": FALSE,
	"null": NULL,
}

func LookupIndent(s string) TokenType {
	if tok, ok := keywords[s]; ok {
		return tok
	}
	return IDENT
}

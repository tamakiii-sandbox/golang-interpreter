package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifier + Literal
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 12345

	// Calculator
	ASSIGN = "="
	PLUS   = "+"
	MINUS = "-"
	BANG = "!"
	ASTERISK = "*"
	SLASH = "/"
	LT = "<"
	GT = ">"

	// Delimiter
	COMMA     = ","
	SEMICOLON = ";"

	// Parenthesis and Brace
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keyword
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType {
	"fn": FUNCTION,
	"let": LET,
}

func LookupIdent(indent string) TokenType {
  if tok, ok := keywords[indent]; ok {
		return tok
	}
	return IDENT
}


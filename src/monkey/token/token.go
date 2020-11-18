package token

type TokenType string

type Token struct {
  Type TokenType
  Literal string
}

const (
  ILLEGAL = "ILLEGAL"
  EOF = "EOF"

  // Identifier + Literal
  IDENT = "IDENT" // add, foobar, x, y, ...
  INT = "INT" // 12345

  // Calculator
  ASSIGN = "="
  PLUS = "+"

  // Delimiter
  COMMA = ","
  SEMICOLON = ";"

  // Parenthesis and Brace
  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"

  // Keyword
  FUNCTION = "FUNCTION"
  LET = "LET"
)

package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// Identifier and literal
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT = "INT" //1343456

	// Operator
	ASSIGN = "="
	PLUS = "+"

	// Decimeter
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = ")"

	// Keyword
	FUNCTION = "FUNCTION"
	LET = "LET"
)
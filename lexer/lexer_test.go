package lexer

import (
	"testing"

	"github.com/iamrishavj/kavi/token"
)

type expectedToken struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	tests := []expectedToken{
		{token.ASSIGN, "="},

		{token.PLUS, "+"},
		{token.L_PAREN, "{"},
		{token.R_PAREN, "}"},
		{token.L_BRACE, "("},
		{token.R_PAREN, ")"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	l := New(input)

	for i, tt := range tests {
		token := l.NextToken()

		if token.Type != tt.expectedType {
			t.Fatalf("test[%d] - literal wrong. \nExpected=%q, got=%q", i, tt.expectedType, token.Literal)
		}
	}

}

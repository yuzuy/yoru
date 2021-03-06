package lexer

import (
	"testing"

	"github.com/yuzuy/yoru/token"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
  x + y;
};

let result = add(five, ten);
!-/*5;
5 < 10 > 5;

if (5 < 10) {
	ret true;
} else {
	ret false;
}

10 == 10;
10 != 9;
"hogehoge";
"nya nya nya";
[1, 2];
5 % 2;
{"foo":"bar"};
switch foo {
case "foo":
  ret 1;
case "bar":
  ret 2;
};
`

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.Let, "let"},
		{token.Ident, "five"},
		{token.Assign, "="},
		{token.Int, "5"},
		{token.Semicolon, ";"},
		{token.Let, "let"},
		{token.Ident, "ten"},
		{token.Assign, "="},
		{token.Int, "10"},
		{token.Semicolon, ";"},
		{token.Let, "let"},
		{token.Ident, "add"},
		{token.Assign, "="},
		{token.Function, "fn"},
		{token.Lparen, "("},
		{token.Ident, "x"},
		{token.Comma, ","},
		{token.Ident, "y"},
		{token.Rparen, ")"},
		{token.Lbrace, "{"},
		{token.Ident, "x"},
		{token.Plus, "+"},
		{token.Ident, "y"},
		{token.Semicolon, ";"},
		{token.Rbrace, "}"},
		{token.Semicolon, ";"},
		{token.Let, "let"},
		{token.Ident, "result"},
		{token.Assign, "="},
		{token.Ident, "add"},
		{token.Lparen, "("},
		{token.Ident, "five"},
		{token.Comma, ","},
		{token.Ident, "ten"},
		{token.Rparen, ")"},
		{token.Semicolon, ";"},
		{token.Bang, "!"},
		{token.Minus, "-"},
		{token.Slash, "/"},
		{token.Asterisk, "*"},
		{token.Int, "5"},
		{token.Semicolon, ";"},
		{token.Int, "5"},
		{token.LT, "<"},
		{token.Int, "10"},
		{token.GT, ">"},
		{token.Int, "5"},
		{token.Semicolon, ";"},
		{token.If, "if"},
		{token.Lparen, "("},
		{token.Int, "5"},
		{token.LT, "<"},
		{token.Int, "10"},
		{token.Rparen, ")"},
		{token.Lbrace, "{"},
		{token.Return, "ret"},
		{token.True, "true"},
		{token.Semicolon, ";"},
		{token.Rbrace, "}"},
		{token.Else, "else"},
		{token.Lbrace, "{"},
		{token.Return, "ret"},
		{token.False, "false"},
		{token.Semicolon, ";"},
		{token.Rbrace, "}"},
		{token.Int, "10"},
		{token.EQ, "=="},
		{token.Int, "10"},
		{token.Semicolon, ";"},
		{token.Int, "10"},
		{token.NotEQ, "!="},
		{token.Int, "9"},
		{token.Semicolon, ";"},
		{token.String, "hogehoge"},
		{token.Semicolon, ";"},
		{token.String, "nya nya nya"},
		{token.Semicolon, ";"},
		{token.LBracket, "["},
		{token.Int, "1"},
		{token.Comma, ","},
		{token.Int, "2"},
		{token.RBracket, "]"},
		{token.Semicolon, ";"},
		{token.Int, "5"},
		{token.Mod, "%"},
		{token.Int, "2"},
		{token.Semicolon, ";"},
		{token.Lbrace, "{"},
		{token.String, "foo"},
		{token.Colon, ":"},
		{token.String, "bar"},
		{token.Rbrace, "}"},
		{token.Semicolon, ";"},
		{token.Switch, "switch"},
		{token.Ident, "foo"},
		{token.Lbrace, "{"},
		{token.Case, "case"},
		{token.String, "foo"},
		{token.Colon, ":"},
		{token.Return, "ret"},
		{token.Int, "1"},
		{token.Semicolon, ";"},
		{token.Case, "case"},
		{token.String, "bar"},
		{token.Colon, ":"},
		{token.Return, "ret"},
		{token.Int, "2"},
		{token.Semicolon, ";"},
		{token.Rbrace, "}"},
		{token.Semicolon, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestStringLiteralError(t *testing.T) {
	input := `"hoge`
	l := New(input)

	tok := l.NextToken()
	if tok.Type != token.Illegal {
		t.Fatalf("token type wrong. expected=%s, got=%s", token.Illegal, tok.Type)
	}
	errMsg := "string literal not terminated"
	if tok.Literal != errMsg {
		t.Fatalf("error message wrong. expected=%s, got=%s", errMsg, tok.Literal)
	}
}

func TestEscapeCharacters(t *testing.T) {
	input := `"\n\r\t\"\'\\"`
	l := New(input)

	tok := l.NextToken()
	if tok.Type != token.String {
		t.Fatalf("token type wrong. expected=%s, got=%s", token.String, tok.Type)
	}
	if tok.Literal != "\n\r\t\"'\\" {
		t.Fatalf("tok.Literal wrong. got=%s", tok.Literal)
	}
}

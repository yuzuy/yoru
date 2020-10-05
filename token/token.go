package token

type Type string

type Token struct {
	Type    Type
	Literal string
}

const (
	Illegal = "ILLEGAL"
	EOF     = "EOF"

	Ident  = "IDENT"
	Int    = "INT"
	String = "STRING"

	Assign   = "="
	Plus     = "+"
	Minus    = "-"
	Bang     = "!"
	Asterisk = "*"
	Slash    = "/"
	Mod      = "%"

	LT = "<"
	GT = ">"

	EQ    = "=="
	NotEQ = "!="

	Comma     = ","
	Colon     = ":"
	Semicolon = ";"

	Lparen   = "("
	Rparen   = ")"
	Lbrace   = "{"
	Rbrace   = "}"
	LBracket = "["
	RBracket = "]"

	Function = "FUNCTION"
	Let      = "LET"
	True     = "TRUE"
	False    = "FALSE"
	If       = "IF"
	Else     = "ELSE"
	Switch   = "SWITCH"
	Case     = "case"
	Default  = "DEFAULT"
	Return   = "RETURN"
	Null     = "NULL"
)

var keywords = map[string]Type{
	"fn":      Function,
	"let":     Let,
	"true":    True,
	"false":   False,
	"if":      If,
	"else":    Else,
	"switch":  Switch,
	"case":    Case,
	"default": Default,
	"return":  Return,
	"null":    Null,
}

func LookUpIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return Ident
}

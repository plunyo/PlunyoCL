package frontend

import "fmt"

type TokenType int

const (
	// literals / identifiers
	Identifier TokenType = iota
	Number
	String

	// operators
	Plus
	Minus
	Star
	Slash
	Percent
	Equal
	DoubleEqual
	Not
	NotEqual
	LessThan
	LessEqual
	GreaterThan
	GreaterEqual
	Or
	And

	// punctuation
	LParen
	RParen
	LBrace
	RBrace
	LBracket
	RBracket
	Semicolon
	Comma
	Dot

	// keywords
	Var
	Const
	If
	Else
	For
	While
	Func
	Return

	// whitespace/comments
	Comment

	// special
	EOF
)

type Token struct {
	Type  TokenType
	Value string
}

func (t TokenType) String() string {
	names := [...]string{
		// literals / identifiers
		"Identifier",
		"Number",
		"String",

		// operators
		"Plus",
		"Minus",
		"Star",
		"Slash",
		"Percent",
		"Equal",
		"DoubleEqual",
		"Not",
		"NotEqual",
		"LessThan",
		"LessEqual",
		"GreaterThan",
		"GreaterEqual",
		"Or",
		"And",

		// punctuation
		"LParen",
		"RParen",
		"LBrace",
		"RBrace",
		"LBracket",
		"RBracket",
		"Semicolon",
		"Comma",
		"Dot",

		// keywords
		"Var",
		"Const",
		"If",
		"Else",
		"For",
		"While",
		"Func",
		"Return",

		// whitespace/comments
		"Comment",

		// special
		"EOF",
	}

	if int(t) < 0 || int(t) >= len(names) {
		return "Unknown"
	}
	
	return names[t]
}

func (t Token) String() string {
	return fmt.Sprintf("Token { Type: %-10s | Value: '%s' }", t.Type, t.Value)
}


// token.go
package frontend

import "fmt"

type TokenType int

const (
	// literals / identifiers
	IdentifierToken TokenType = iota
	NumberToken
	StringToken

	// operators
	PlusToken
	MinusToken
	StarToken
	SlashToken
	PercentToken
	EqualToken
	DoubleEqualToken
	NotEqualToken
	LessThanToken
	LessEqualToken
	GreaterThanToken
	GreaterEqualToken
	LogicalOrToken
	LogicalAndToken
	LogicalNotToken
	BitwiseOrToken
	BitwiseAndToken
	BitwiseNotToken

	// punctuation
	LParenToken
	RParenToken
	LBraceToken
	RBraceToken
	LBracketToken
	RBracketToken
	SemicolonToken
	CommaToken
	DotToken

	// keywords
	VarToken
	IfToken
	ElseToken
	ForToken
	WhileToken
	FuncToken
	ReturnToken

	// whitespace/comments
	CommentToken

	// special
	EOFToken
)

type Token struct {
	Type  TokenType
	Value string
}

func (token TokenType) String() string {
	names := [...]string{
		// literals / identifiers
		"IdentifierToken",
		"NumberToken",
		"StringToken",

		// operators
		"PlusToken",
		"MinusToken",
		"StarToken",
		"SlashToken",
		"PercentToken",
		"EqualToken",
		"DoubleEqualToken",
		"NotEqualToken",
		"LessThanToken",
		"LessEqualToken",
		"GreaterThanToken",
		"GreaterEqualToken",
		"LogicalOrToken",
		"LogicalAndToken",
		"LogicalNotToken",
		"BitwiseOrToken",
		"BitwiseAndToken",
		"BitwiseNotToken",

		// punctuation
		"LParenToken",
		"RParenToken",
		"LBraceToken",
		"RBraceToken",
		"LBracketToken",
		"RBracketToken",
		"SemicolonToken",
		"CommaToken",
		"DotToken",

		// keywords
		"VarToken",
		"IfToken",
		"ElseToken",
		"ForToken",
		"WhileToken",
		"FuncToken",
		"ReturnToken",

		// whitespace/comments
		"CommentToken",

		// special
		"EOFToken",
	}

	if int(token) < 0 || int(token) >= len(names) {
		return "UnknownToken"
	}

	return names[token]
}

func (token Token) String() string {
	return fmt.Sprintf("Token { Type: %-15s | Value: '%s' }", token.Type, token.Value)
}

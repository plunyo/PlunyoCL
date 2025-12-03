package frontend

import (
	"strconv"
)

type Parser struct {
	tokens []Token
	pos    int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{tokens: tokens, pos: 0}
}

func (parser *Parser) peek() *Token {
	if parser.pos >= len(parser.tokens) {
		return nil
	}
	return &parser.tokens[parser.pos]
}

func (parser *Parser) eat() *Token {
	if parser.pos >= len(parser.tokens) {
		return nil
	}

	t := &parser.tokens[parser.pos]
	parser.pos++
	return t
}

func (parser *Parser) expect(tType TokenType, msg string) *Token {
	if t := parser.peek(); t != nil && t.Type == tType {
		return parser.eat()
	}
	panic("unexpected token: " + msg)
}

func (parser *Parser) GenerateAST() ASTNode {
	root := &ProgramNode{Statements: []ASTNode{}}

	for t := parser.peek(); t != nil && t.Type != EOFToken; t = parser.peek() {
		root.Statements = append(root.Statements, parser.parseStatement())
	}

	return root
}

func (parser *Parser) parseStatement() ASTNode {
	tok := parser.peek()
	
	if tok == nil {
		panic("unexpected end of input")
	}


	switch tok.Type { 
		case VarToken: return parser.parseVarDecl()
	}

	panic("unknown statement starting at token: " + tok.Value)
}

func (parser *Parser) parseVarDecl() ASTNode {
	parser.eat()

	name := parser.expect(IdentifierToken, "expected identifier after 'var'")
	parser.expect(EqualToken, "expected '=' after identifier")

	value := parser.parseExpression()
	parser.expect(SemicolonToken, "expected ';' after expression")

	return &VarDeclNode{
		Name:    name.Value,
		Value:   &value,
	}
}

func (parser *Parser) parseExpression() ASTNode {
	return parser.parseAdditive()
}

func (parser *Parser) parseAdditive() ASTNode {
	left := parser.parseMultiplicative()

	for {
		token := parser.peek()
		if token == nil {
			break
		}

		if token.Type == PlusToken || token.Type == MinusToken {
			parser.eat()
			right := parser.parseMultiplicative()

			left = &BinaryOpNode{
				Left:     &left,
				Operator: *token,
				Right:    &right,
			}
		}
	}

	return left
}

func (parser *Parser) parseMultiplicative() ASTNode {
	left := parser.parseLiteral()
	
	for {
		token := parser.peek()
		if token == nil {
			break
		}

		if token.Type == StarToken || token.Type == SlashToken {
			parser.eat()
			right := parser.parseLiteral()

			left = &BinaryOpNode{
				Left:     &left,
				Operator: *token,
				Right:    &right,
			}
		} else {
			break
		}
	}

	return left
}

func (parser *Parser) parseLiteral() ASTNode {
	token := parser.peek()
	if token == nil {
		panic("unexpected end of input in term")
	}

	switch token.Type {
		case NumberToken:
			parser.eat()
			val, err := strconv.ParseFloat(token.Value, 64)

			if err != nil {
				panic("invalid number format")
			}

			return &NumberLiteralNode{Value: val}

		case StringToken:
			parser.eat()
			return &StringLiteralNode{Value: token.Value}

	}

	panic("unknown term: " + token.Value)
}

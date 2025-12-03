package frontend

import (
	"strconv"
	"strings"
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
		case IdentifierToken: return parser.parseAssignment()
	}

	panic("unknown statement starting at token: " + tok.String())
}

func (parser *Parser) parseAssignment() ASTNode {
	nameToken := parser.expect(IdentifierToken, "expected identifier at start of assignment")
	parser.expect(EqualToken, "expected '=' after identifier")

	value := parser.parseExpression()
	parser.expect(SemicolonToken, "expected ';' after expression")

	return &AssignmentNode{
		Name:  nameToken.Value,
		Value: value,
	}
}

func (parser *Parser) parseVarDecl() ASTNode {
	parser.eat()

	name := parser.expect(IdentifierToken, "expected identifier after 'var'")

	currentToken := parser.peek()
	switch currentToken.Type {
		case EqualToken:
			parser.eat() // eat =

			value := parser.parseExpression()
			parser.expect(SemicolonToken, "expected ';' after expression")

			return &VarDeclNode{
				Name:    name.Value,
				Value:   value,
			}
		case SemicolonToken :
			parser.expect(SemicolonToken, "expected ';' after variable declaration")
			return &VarDeclNode{
				Name:	name.Value,
				Value:  nil,
			}
	}

	panic("unexpected token: " + currentToken.String())
}

func (parser *Parser) parseExpression() ASTNode {
	return parser.parseComparison()
}

func (parser *Parser) parseComparison() ASTNode {
	left := parser.parseAdditive()

	for {
		token := parser.peek()
		if token == nil {
			break
		}

		if token.Type == LessThanToken || token.Type == GreaterThanToken ||
			token.Type == DoubleEqualToken || token.Type == NotEqualToken ||
			token.Type == LessEqualToken || token.Type == GreaterEqualToken {

			parser.eat()
			right := parser.parseExpression()

			left = &BinaryOpNode{
				Left:     left,
				Operator: token.Value,
				Right:    right,
			}
		} else {
			break
		}
	}

	return left
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
				Left:     left,
				Operator: token.Value,
				Right:    right,
			}
		} else {
			break
		}
	}

	return left
}

func (parser *Parser) parseMultiplicative() ASTNode {
	left := parser.parseUnary()
	
	for {
		token := parser.peek()
		if token == nil {
			break
		}

		if token.Type == StarToken || token.Type == SlashToken || token.Type == PercentToken {
			parser.eat()
			right := parser.parseUnary()

			left = &BinaryOpNode{
				Left:     left,
				Operator: token.Value,
				Right:    right,
			}
		} else {
			break
		}
	}

	return left
}

func (parser *Parser) parseUnary() ASTNode {
	token := parser.peek()
	if token == nil {
		panic("unexpected end of input in unary")
	}

	if token.Type == PlusToken || token.Type == MinusToken {
		parser.eat()
		operand := parser.parseUnary()

		return &UnaryOpNode{
			Operator: token.Value,
			Operand:  operand,
		}
	}

	return parser.parsePrimary()
}

func (parser *Parser) parsePrimary() ASTNode {
	token := parser.peek()
	if token == nil {
		panic("unexpected end of input in term")
	}

	switch token.Type {
		case NumberToken:
			parser.eat()
			val := token.Value

			if strings.Contains(val, ".") || strings.ContainsAny(val, "eE") {
				num, err := strconv.ParseFloat(val, 64)
				if err != nil {
					panic("invalid float literal: " + val)
				}

				return &LiteralNode[float64]{Value: num}
			} else {
				num, err := strconv.Atoi(val)
				if err != nil {
					panic("invalid int literal: " + val)
				}

				return &LiteralNode[int]{Value: num}
			}

		case StringToken:
			parser.eat()
			return &LiteralNode[string]{Value: token.Value}
		case LParenToken:
			parser.eat()
			expr := parser.parseExpression()
			parser.expect(RParenToken, "expected ')' after expression")
			return expr
		case IdentifierToken:
			parser.eat()
			return &IdentifierNode{Name: token.Value}

	}

	panic("unknown term: " + token.Value)
}

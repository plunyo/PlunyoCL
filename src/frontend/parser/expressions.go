package parser

import (
	"pcl/src/frontend/ast"
	"pcl/src/frontend/lexer"
	"strconv"
	"strings"
)

func (parser *Parser) parseExpression() ast.ASTNode {
	return parser.parseComparison()
}

func (parser *Parser) parseComparison() ast.ASTNode {
	left := parser.parseAdditive()

	for {
		token := parser.peek()
		if token == nil {
			break
		}

		if token.Type == lexer.LessThanToken || token.Type == lexer.GreaterThanToken ||
			token.Type == lexer.DoubleEqualToken || token.Type == lexer.NotEqualToken ||
			token.Type == lexer.LessEqualToken || token.Type == lexer.GreaterEqualToken ||
			token.Type == lexer.LogicalAndToken || token.Type == lexer.LogicalOrToken {

			parser.eat()
			right := parser.parseAdditive()

			left = &ast.BinaryOpNode{
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

func (parser *Parser) parseAdditive() ast.ASTNode {
	left := parser.parseMultiplicative()

	for {
		token := parser.peek()
		if token == nil {
			break
		}

		if token.Type == lexer.PlusToken || token.Type == lexer.MinusToken {
			parser.eat()
			right := parser.parseMultiplicative()
			left = &ast.BinaryOpNode{Left: left, Operator: token.Value, Right: right}
		} else {
			break
		}
	}

	return left
}

func (parser *Parser) parseMultiplicative() ast.ASTNode {
	left := parser.parseUnary()

	for {
		token := parser.peek()
		if token == nil {
			break
		}

		if token.Type == lexer.StarToken || token.Type == lexer.SlashToken || token.Type == lexer.PercentToken {
			parser.eat()
			right := parser.parseUnary()
			left = &ast.BinaryOpNode{Left: left, Operator: token.Value, Right: right}
		} else {
			break
		}
	}

	return left
}

func (parser *Parser) parseUnary() ast.ASTNode {
	token := parser.peek()
	if token == nil {
		panic("unexpected end of input in unary")
	}

	if token.Type == lexer.PlusToken || token.Type == lexer.MinusToken {
		parser.eat()
		return &ast.UnaryOpNode{Operator: token.Value, Operand: parser.parseUnary()}
	}

	return parser.parsePrimary()
}

func (parser *Parser) parsePrimary() ast.ASTNode {
	token := parser.peek()
	if token == nil {
		panic("unexpected end of input in term")
	}

	switch token.Type {

	case lexer.FuncToken:
		return parser.parseFunctionLiteral()

	case lexer.NumberToken:
		parser.eat()
		val := token.Value
		if strings.Contains(val, ".") || strings.ContainsAny(val, "eE") {
			num, err := strconv.ParseFloat(val, 64)
			if err != nil {
				panic("invalid float literal: " + val)
			}
			return &ast.LiteralNode[float64]{Value: num}
		} else {
			num, err := strconv.Atoi(val)
			if err != nil {
				panic("invalid int literal: " + val)
			}
			return &ast.LiteralNode[int]{Value: num}
		}

	case lexer.StringToken:
		parser.eat()
		return &ast.LiteralNode[string]{Value: token.Value}

	case lexer.LParenToken:
		parser.eat()
		expr := parser.parseExpression()
		parser.expect(lexer.RParenToken, "expected ')' after expression")
		return expr

	case lexer.IdentifierToken:
		return parser.parseFunctionCall()

	default:
		panic("unknown term: " + token.Value)
	}
}

// ---------- Identifier ----------

func (parser *Parser) parseIdentifier() *ast.IdentifierNode {
	identToken := parser.eat() // eat identifier
	return &ast.IdentifierNode{Name: identToken.Value}
}

// ---------- Function Call ----------

func (parser *Parser) parseFunctionCall() ast.ASTNode {
	identifier := parser.parseIdentifier()

	if parser.peek() != nil && parser.peek().Type == lexer.LParenToken {
		parser.eat() // eat '('

		functionCallNode := &ast.FunctionCallNode{
			Callee:    identifier,
			Arguments: []ast.ASTNode{},
		}

		for parser.peek() != nil && parser.peek().Type != lexer.RParenToken {
			arg := parser.parseExpression()
			functionCallNode.Arguments = append(functionCallNode.Arguments, arg)

			if parser.peek() != nil && parser.peek().Type == lexer.CommaToken {
				parser.eat() // eat ','
			}
		}

		parser.expect(lexer.RParenToken, "expected ')' after arguments")

		return functionCallNode
	}

	return identifier
}

func (parser *Parser) parseFunctionLiteral() ast.ASTNode {
	parser.expect(lexer.FuncToken, "expected 'func'")

	parser.expect(lexer.LParenToken, "expected '(' after func")

	var params []string
	for t := parser.peek(); t != nil && t.Type != lexer.RParenToken; t = parser.peek() {
		paramName := parser.expect(lexer.IdentifierToken, "expected parameter name")
		params = append(params, paramName.Value)

		if parser.peek().Type == lexer.CommaToken {
			parser.eat()
		}

	}

	parser.expect(lexer.RParenToken, "expected ')' after parameters")
	parser.expect(lexer.LBraceToken, "expected '{' before function body")

	var body []ast.ASTNode

	for t := parser.peek(); t != nil && t.Type != lexer.RBraceToken; t = parser.peek() {
		body = append(body, parser.parseStatement())
	}

	parser.expect(lexer.RBraceToken, "expected '}' after function body")

	return &ast.FunctionLiteralNode{
		Arguments: params,
		Body:      &ast.BodyNode{Statements: body},
	}
}

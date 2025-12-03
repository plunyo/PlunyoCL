package parser

import (
	"pcl/src/frontend/ast"
	"pcl/src/frontend/lexer"
	"strconv"
	"strings"
)

func (p *Parser) parseExpression() ast.ASTNode {
	return p.parseComparison()
}

func (p *Parser) parseComparison() ast.ASTNode {
	left := p.parseAdditive()

	for {
		token := p.peek()
		if token == nil {
			break
		}

		if token.Type == lexer.LessThanToken || token.Type == lexer.GreaterThanToken ||
			token.Type == lexer.DoubleEqualToken || token.Type == lexer.NotEqualToken ||
			token.Type == lexer.LessEqualToken || token.Type == lexer.GreaterEqualToken ||
			token.Type == lexer.LogicalAndToken || token.Type == lexer.LogicalOrToken ||
			token.Type == lexer.LogicalNotToken {

			p.eat()
			right := p.parseExpression()

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

func (p *Parser) parseAdditive() ast.ASTNode {
	left := p.parseMultiplicative()

	for {
		token := p.peek()
		if token == nil {
			break
		}

		if token.Type == lexer.PlusToken || token.Type == lexer.MinusToken {
			p.eat()
			right := p.parseMultiplicative()
			left = &ast.BinaryOpNode{Left: left, Operator: token.Value, Right: right}
		} else {
			break
		}
	}

	return left
}

func (p *Parser) parseMultiplicative() ast.ASTNode {
	left := p.parseUnary()

	for {
		token := p.peek()
		if token == nil {
			break
		}

		if token.Type == lexer.StarToken || token.Type == lexer.SlashToken || token.Type == lexer.PercentToken {
			p.eat()
			right := p.parseUnary()
			left = &ast.BinaryOpNode{Left: left, Operator: token.Value, Right: right}
		} else {
			break
		}
	}

	return left
}

func (p *Parser) parseUnary() ast.ASTNode {
	token := p.peek()
	if token == nil {
		panic("unexpected end of input in unary")
	}

	if token.Type == lexer.PlusToken || token.Type == lexer.MinusToken {
		p.eat()
		return &ast.UnaryOpNode{Operator: token.Value, Operand: p.parseUnary()}
	}

	return p.parsePrimary()
}

func (p *Parser) parsePrimary() ast.ASTNode {
	token := p.peek()
	if token == nil {
		panic("unexpected end of input in term")
	}

	switch token.Type {
	case lexer.NumberToken:
		p.eat()
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
		p.eat()
		return &ast.LiteralNode[string]{Value: token.Value}

	case lexer.LParenToken:
		p.eat()
		expr := p.parseExpression()
		p.expect(lexer.RParenToken, "expected ')' after expression")
		return expr

	case lexer.IdentifierToken:
		p.eat()
		return &ast.IdentifierNode{Name: token.Value}
	}

	panic("unknown term: " + token.Value)
}

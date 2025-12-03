package parser

import (
	"pcl/src/frontend/ast"
	"pcl/src/frontend/lexer"
)

func (p *Parser) parseAssignment() ast.ASTNode {
	nameToken := p.expect(lexer.IdentifierToken, "expected identifier at start of assignment")
	p.expect(lexer.EqualToken, "expected '=' after identifier")

	value := p.parseExpression()
	p.expect(lexer.SemicolonToken, "expected ';' after expression")

	return &ast.AssignmentNode{
		Name:  nameToken.Value,
		Value: value,
	}
}

func (p *Parser) parseVarDecl() ast.ASTNode {
	p.eat() // eat 'var'
	name := p.expect(lexer.IdentifierToken, "expected identifier after 'var'")

	currentToken := p.peek()
	switch currentToken.Type {
	case lexer.EqualToken:
		p.eat() // eat '='
		value := p.parseExpression()
		p.expect(lexer.SemicolonToken, "expected ';' after expression")
		return &ast.VarDeclNode{Name: name.Value, Value: value}

	case lexer.SemicolonToken:
		p.expect(lexer.SemicolonToken, "expected ';' after variable declaration")
		return &ast.VarDeclNode{Name: name.Value, Value: nil}
	}

	panic("unexpected token: " + currentToken.String())
}

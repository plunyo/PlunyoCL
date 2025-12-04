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

func (p *Parser) parseFuncDecl() ast.ASTNode {
	p.eat() // eat 'func'
	name := p.expect(lexer.IdentifierToken, "expected function name after 'func'")
	p.expect(lexer.LParenToken, "expected '(' after function name")

	// Parse parameters
	var params []string
	for t := p.peek(); t != nil && t.Type != lexer.RParenToken; t = p.peek() {
		paramName := p.expect(lexer.IdentifierToken, "expected parameter name")
		params = append(params, paramName.Value)

		if p.peek().Type == lexer.CommaToken {
			p.eat() // eat ','
		}
	}

	p.expect(lexer.RParenToken, "expected ')' after parameters")
	p.expect(lexer.LBraceToken, "expected '{' before function body")

	// Parse body as statements
	var statements []ast.ASTNode
	for t := p.peek(); t != nil && t.Type != lexer.RBraceToken; t = p.peek() {
		statements = append(statements, p.parseStatement())
	}

	p.expect(lexer.RBraceToken, "expected '}' after function body")

	return &ast.FunctionDeclNode{
		Name:       name.Value,
		Arguments:  params,
		Statements: statements,
	}
}

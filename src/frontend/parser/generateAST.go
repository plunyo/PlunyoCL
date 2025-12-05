package parser

import (
	"pcl/src/frontend/ast"
	"pcl/src/frontend/lexer"
)

func (p *Parser) GenerateAST() ast.ASTNode {
	program := &ast.ProgramNode{Statements: []ast.ASTNode{}}

	for t := p.peek(); t != nil && t.Type != lexer.EOFToken; t = p.peek() {
		program.Statements = append(program.Statements, p.parseStatement())
	}

	return program
}

func (p *Parser) parseStatement() ast.ASTNode {
	tok := p.peek()
	if tok == nil {
		panic("unexpected end of input")
	}

	switch tok.Type {
	case lexer.VarToken:
		return p.parseVarDecl()
	case lexer.FuncToken:
		return p.parseFuncDecl()
	case lexer.LBraceToken:
		return p.parseBody()
	case lexer.IdentifierToken:
		// Look ahead to see if it's an assignment or an expression statement
		if p.peekAhead(1) != nil && p.peekAhead(1).Type == lexer.EqualToken {
			return p.parseAssignment()
		}
		// Otherwise it's an expression statement
		expr := p.parseExpression()
		if p.peek() != nil && p.peek().Type == lexer.SemicolonToken {
			p.eat() // eat ';'
		}
		return expr
	default:
		panic("unknown statement starting at token: " + tok.String())
	}
}

func (p *Parser) parseBody() ast.ASTNode {
	p.eat() // eat '{'
	block := &ast.BodyNode{Statements: []ast.ASTNode{}}

	for t := p.peek(); t != nil && t.Type != lexer.RBraceToken; t = p.peek() {
		block.Statements = append(block.Statements, p.parseStatement())
	}

	p.eat() // eat '}'
	return block
}

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
	case lexer.IdentifierToken:
		return p.parseAssignment()
	case lexer.LBraceToken:
		return p.parseBlock()
	}

	panic("unknown statement starting at token: " + tok.String())
}

func (p *Parser) parseBlock() ast.ASTNode {
	p.eat() // eat '{'
	block := &ast.BlockNode{Statements: []ast.ASTNode{}}

	for t := p.peek(); t != nil && t.Type != lexer.RBraceToken; t = p.peek() {
		block.Statements = append(block.Statements, p.parseStatement())
	}

	p.eat() // eat '}'
	return block
}

package parser

import (
	"pcl/src/frontend/ast"
	"pcl/src/frontend/lexer"
)

func (parser *Parser) GenerateAST() ast.ASTNode {
	program := &ast.ProgramNode{Statements: []ast.ASTNode{}}

	for t := parser.peek(); t != nil && t.Type != lexer.EOFToken; t = parser.peek() {
		program.Statements = append(program.Statements, parser.parseStatement())
	}

	return program
}

func (parser *Parser) parseStatement() ast.ASTNode {
	tok := parser.peek()
	if tok == nil {
		panic("unexpected end of input")
	}

	switch tok.Type {
	case lexer.VarToken:
		return parser.parseVarDecl()
	case lexer.FuncToken:
		return parser.parseFuncDecl()
	case lexer.ReturnToken:
		return parser.parseReturnStatement()
	case lexer.LBraceToken:
		return parser.parseBody()
	case lexer.IdentifierToken:
		if parser.peekAhead(1) != nil && parser.peekAhead(1).Type == lexer.EqualToken {
			return parser.parseAssignment()
		}

		expr := parser.parseExpression()
		if parser.peek() != nil && parser.peek().Type == lexer.SemicolonToken {
			parser.eat() // eat ';'
		}
		
		return expr
	default:
		panic("unknown statement starting at token: " + tok.String())
	}
}

func (parser *Parser) parseBody() ast.ASTNode {
	parser.eat() // eat '{'
	block := &ast.BodyNode{Statements: []ast.ASTNode{}}

	for t := parser.peek(); t != nil && t.Type != lexer.RBraceToken; t = parser.peek() {
		block.Statements = append(block.Statements, parser.parseStatement())
	}

	parser.eat() // eat '}'
	return block
}

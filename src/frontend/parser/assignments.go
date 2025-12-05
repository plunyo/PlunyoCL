package parser

import (
	"pcl/src/frontend/ast"
	"pcl/src/frontend/lexer"
)

func (parser *Parser) parseAssignment() ast.ASTNode {
	nameToken := parser.expect(lexer.IdentifierToken, "expected identifier at start of assignment")
	parser.expect(lexer.EqualToken, "expected '=' after identifier")

	value := parser.parseExpression()
	parser.expect(lexer.SemicolonToken, "expected ';' after expression")

	return &ast.AssignmentNode{
		Name:  nameToken.Value,
		Value: value,
	}
}

func (parser *Parser) parseVarDecl() ast.ASTNode {
	parser.eat() // eat 'var'
	name := parser.expect(lexer.IdentifierToken, "expected identifier after 'var'")

	currentToken := parser.peek()
	switch currentToken.Type {
	case lexer.EqualToken:
		parser.eat() // eat '='
		value := parser.parseExpression()
		parser.expect(lexer.SemicolonToken, "expected ';' after expression")
		return &ast.VarDeclNode{Name: name.Value, Value: value}

	case lexer.SemicolonToken:
		parser.expect(lexer.SemicolonToken, "expected ';' after variable declaration")
		return &ast.VarDeclNode{Name: name.Value, Value: nil}
	}

	panic("unexpected token: " + currentToken.String())
}

func (parser *Parser) parseFuncDecl() ast.ASTNode {
    parser.eat() // eat 'func'
    name := parser.expect(lexer.IdentifierToken, "expected function name after 'func'")

    // parse the literal starting from '('
    literal := parser.parseFunctionLiteral()

    // wrap it inside a var decl node
    return &ast.VarDeclNode{
        Name:  name.Value,
        Value: literal,
    }
}


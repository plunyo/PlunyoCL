package parser

import (
	"pcl/src/frontend/ast"
	"pcl/src/frontend/lexer"
)

func (p *Parser) parseReturnStatement() ast.ASTNode {
    p.expect(lexer.ReturnToken, "expected 'return'")

    var value ast.ASTNode = nil
    if p.peek().Type != lexer.SemicolonToken {
        value = p.parseExpression()
    }

    p.expect(lexer.SemicolonToken, "expected ';' after return")

    return &ast.ReturnNode{Value: value}
}

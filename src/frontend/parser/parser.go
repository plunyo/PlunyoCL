package parser

import "pcl/src/frontend/lexer"

type Parser struct {
	tokens []lexer.Token
	pos    int
}

func NewParser(tokens []lexer.Token) *Parser {
	return &Parser{tokens: tokens, pos: 0}
}

// --- token helpers ---
func (p *Parser) peek() *lexer.Token {
	if p.pos >= len(p.tokens) {
		return nil
	}
	return &p.tokens[p.pos]
}

func (p *Parser) peekAhead(n int) *lexer.Token {
	pos := p.pos + n
	if pos >= len(p.tokens) {
		return nil
	}
	return &p.tokens[pos]
}

func (p *Parser) eat() *lexer.Token {
	if p.pos >= len(p.tokens) {
		return nil
	}
	t := &p.tokens[p.pos]
	p.pos++
	return t
}

func (p *Parser) expect(tType lexer.TokenType, msg string) *lexer.Token {
	if t := p.peek(); t != nil && t.Type == tType {
		return p.eat()
	}
	panic("unexpected token: " + msg)
}

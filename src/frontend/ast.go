package frontend

type NodeType int

const (
	ProgramNodeType NodeType = iota
	VarDeclNodeType
	BinaryOpNodeType
	UnaryOpNodeType
	NumberLiteralNodeType
	StringLiteralNodeType
	IdentifierNodeType
)

type ASTNode interface {
	Type() NodeType
}

// program
type ProgramNode struct {
	Statements []ASTNode
}

func (p *ProgramNode) Type() NodeType {
	return ProgramNodeType
}

// var declaration
type VarDeclNode struct {
	Name  string
	Value ASTNode
}

func (v *VarDeclNode) Type() NodeType {
	return VarDeclNodeType
}

// binary operation
type BinaryOpNode struct {
	Left     ASTNode
	Operator Token
	Right    ASTNode
}

func (b *BinaryOpNode) Type() NodeType {
	return BinaryOpNodeType
}

// unary operation
type UnaryOpNode struct {
	Operator Token
	Operand  ASTNode
}

func (u *UnaryOpNode) Type() NodeType {
	return UnaryOpNodeType
}

// number literal
type NumberLiteralNode struct {
	Value string
}

func (n *NumberLiteralNode) Type() NodeType {
	return NumberLiteralNodeType
}

// string literal
type StringLiteralNode struct {
	Value string
}

func (s *StringLiteralNode) Type() NodeType {
	return StringLiteralNodeType
}

// identifier
type IdentifierNode struct {
	Name string
}

func (i *IdentifierNode) Type() NodeType {
	return IdentifierNodeType
}

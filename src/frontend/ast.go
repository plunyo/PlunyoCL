package frontend

import (
	"fmt"
	"strings"
)

// ---------- AST Node Types ----------

type NodeType int

const (
	ProgramNodeType NodeType = iota
	BlockNodeType
	VarDeclNodeType
	AssignmentNodeType
	BinaryOpNodeType
	UnaryOpNodeType
	IntLiteralNodeType
	FloatLiteralNodeType
	StringLiteralNodeType
	BooleanLiteralNodeType
	IdentifierNodeType
)

type ASTNode interface {
	Type() NodeType
	String() string
}

// ---------- Pretty printing helpers ----------

func indentStr(level int) string {
	return strings.Repeat("  ", level)
}

func pretty(node ASTNode, level int) string {
	if node == nil {
		return indentStr(level) + "nil"
	}

	switch n := node.(type) {
		case *ProgramNode:
			sb := &strings.Builder{}
			sb.WriteString(indentStr(level) + "ProgramNode {\n")
			for _, stmt := range n.Statements {
				sb.WriteString(pretty(stmt, level+1) + ",\n")
			}
			sb.WriteString(indentStr(level) + "}")
			return sb.String()

		case *BlockNode:
			sb := &strings.Builder{}
			sb.WriteString(indentStr(level) + "BlockNode {\n")
			for _, stmt := range n.Statements {
				sb.WriteString(pretty(stmt, level+1) + ",\n")
			}
			sb.WriteString(indentStr(level) + "}")
			return sb.String()

		case *VarDeclNode:
			return formatNode("VarDeclNode", level, map[string]ASTNode{
				"Name":  &IdentifierNode{Name: n.Name},
				"Value": n.Value,
			})

		case *AssignmentNode:
			return formatNode("AssignmentNode", level, map[string]ASTNode{
				"Name":  &IdentifierNode{Name: n.Name},
				"Value": n.Value,
			})

		case *BinaryOpNode:
			return formatNode("BinaryOpNode", level, map[string]ASTNode{
				"Operator": &LiteralNode[string]{Value: n.Operator},
				"Left":     n.Left,
				"Right":    n.Right,
			})

		case *UnaryOpNode:
			return formatNode("UnaryOpNode", level, map[string]ASTNode{
				"Operator": &LiteralNode[string]{Value: n.Operator},
				"Operand":  n.Operand,
			})

		case *LiteralNode[int]:
			return indentStr(level) + fmt.Sprintf("IntLiteralNode { Value: %v }", n.Value)

		case *LiteralNode[float64]:
			return indentStr(level) + fmt.Sprintf("FloatLiteralNode { Value: %v }", n.Value)

		case *LiteralNode[string]:
			return indentStr(level) + `StringLiteralNode { Value: "` + n.Value + `" }`

		case *IdentifierNode:
			return indentStr(level) + "IdentifierNode { Name: " + n.Name + " }"

		default:
			return indentStr(level) + node.String()
	}
}

func formatNode(name string, level int, fields map[string]ASTNode) string {
	sb := &strings.Builder{}
	sb.WriteString(indentStr(level) + name + " {\n")
	for k, v := range fields {
		sb.WriteString(indentStr(level+1) + k + ":\n")
		sb.WriteString(pretty(v, level+2) + "\n")
	}
	sb.WriteString(indentStr(level) + "}")
	return sb.String()
}

// ---------- Concrete AST Nodes ----------

type ProgramNode struct{ Statements []ASTNode }
func (p *ProgramNode) Type() NodeType { return ProgramNodeType }
func (p *ProgramNode) String() string { return pretty(p, 0) }

type BlockNode struct { Statements []ASTNode }
func (b *BlockNode) Type() NodeType { return BlockNodeType }
func (b *BlockNode) String() string { return pretty(b, 0) }

type VarDeclNode struct {
	Name  string
	Value ASTNode
}
func (v *VarDeclNode) Type() NodeType { return VarDeclNodeType }
func (v *VarDeclNode) String() string { return pretty(v, 0) }

type AssignmentNode struct {
	Name  string
	Value ASTNode
}
func (a *AssignmentNode) Type() NodeType { return AssignmentNodeType }
func (a *AssignmentNode) String() string { return pretty(a, 0) }

type BinaryOpNode struct {
	Left, Right ASTNode
	Operator    string
}
func (b *BinaryOpNode) Type() NodeType { return BinaryOpNodeType }
func (b *BinaryOpNode) String() string { return pretty(b, 0) }

type UnaryOpNode struct {
	Operator string
	Operand  ASTNode
}
func (u *UnaryOpNode) Type() NodeType { return UnaryOpNodeType }
func (u *UnaryOpNode) String() string { return pretty(u, 0) }

type IdentifierNode struct{ Name string }
func (i *IdentifierNode) Type() NodeType { return IdentifierNodeType }
func (i *IdentifierNode) String() string { return pretty(i, 0) }

type LiteralNode[T int | float64 | string] struct { Value T }
func (l *LiteralNode[T]) Type() NodeType {
	switch any(l.Value).(type) {
	case int:
		return IntLiteralNodeType
	case float64:
		return FloatLiteralNodeType
	case string:
		return StringLiteralNodeType
	default:
		panic("unknown literal type")
	}
}
func (l *LiteralNode[T]) String() string { return pretty(l, 0) }
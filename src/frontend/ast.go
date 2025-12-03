package frontend

import (
	"fmt"
	"strings"
)

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
	String() string
}

// helper: produce indentation
func indentStr(level int) string {
	return strings.Repeat("    ", level)
}

// generic pretty-printer (uses a type switch so concrete nodes are printed with indentation)
func pretty(node ASTNode, level int) string {
	if node == nil {
		return indentStr(level) + "nil"
	}

	switch n := node.(type) {
	case *ProgramNode:
		var sb strings.Builder
		sb.WriteString(indentStr(level) + "ProgramNode {\n")
		for _, stmt := range n.Statements {
			sb.WriteString(pretty(stmt, level+1))
			sb.WriteString("\n")
		}
		sb.WriteString(indentStr(level) + "}")
		return sb.String()

	case *VarDeclNode:
		var sb strings.Builder

		sb.WriteString(indentStr(level) + "VarDeclNode {\n")
		sb.WriteString(indentStr(level+1) + "Name: " + n.Name + "\n")

		// Value is *ASTNode (pointer to interface) in your definition — handle nil safely
		if n.Value == nil {
			sb.WriteString(indentStr(level+1) + "Value: nil\n")
		} else {
			sb.WriteString(indentStr(level+1) + "Value:\n")
			sb.WriteString(pretty(*n.Value, level+2))
			sb.WriteString("\n")
		}
		
		sb.WriteString(indentStr(level) + "}")
		return sb.String()

	case *BinaryOpNode:
		var sb strings.Builder
		sb.WriteString(indentStr(level) + "BinaryOpNode {\n")
		sb.WriteString(indentStr(level+1) + "Operator: " + fmt.Sprintf("%v", n.Operator) + "\n")

		if n.Left == nil {
			sb.WriteString(indentStr(level+1) + "Left: nil\n")
		} else {
			sb.WriteString(indentStr(level+1) + "Left:\n")
			sb.WriteString(pretty(*n.Left, level+2))
			sb.WriteString("\n")
		}

		if n.Right == nil {
			sb.WriteString(indentStr(level+1) + "Right: nil\n")
// generic pretty-printer (uses a type switch so concrete nodes are printed with indentation)r(level+1) + "Right: nil\n")
		} else {
			sb.WriteString(indentStr(level+1) + "Right:\n")
			sb.WriteString(pretty(*n.Right, level+2))
			sb.WriteString("\n")
		}

		sb.WriteString(indentStr(level) + "}")
		return sb.String()

	case *UnaryOpNode:
		var sb strings.Builder
		sb.WriteString(indentStr(level) + "UnaryOpNode {\n")
		sb.WriteString(indentStr(level+1) + "Operator: " + fmt.Sprintf("%v", n.Operator) + "\n")

		if n.Operand == nil {
			sb.WriteString(indentStr(level+1) + "Operand: nil\n")
		} else {
			sb.WriteString(indentStr(level+1) + "Operand:\n")
			sb.WriteString(pretty(*n.Operand, level+2))
			sb.WriteString("\n")
		}

		sb.WriteString(indentStr(level) + "}")
		return sb.String()

	case *NumberLiteralNode:
		return indentStr(level) + "NumberLiteralNode { Value: " + fmt.Sprintf("%v", n.Value) + " }"

	case *StringLiteralNode:
		return indentStr(level) + `StringLiteralNode { Value: "` + n.Value + `" }`

	case *IdentifierNode:
		return indentStr(level) + "IdentifierNode { Name: " + n.Name + " }"

	default:
		// fallback to the node's String() if it's some other type
		return indentStr(level) + node.String()
	}
}

// ---------- concrete types (String() delegates to pretty with level 0) ----------

// program
type ProgramNode struct {
	Statements []ASTNode
}

func (p *ProgramNode) Type() NodeType { return ProgramNodeType }

func (p *ProgramNode) String() string { return pretty(p, 0) }

// var declaration
type VarDeclNode struct {
	Name    string
	Value   *ASTNode
}

func (v *VarDeclNode) Type() NodeType { return VarDeclNodeType }

func (v *VarDeclNode) String() string { return pretty(v, 0) }

// binary operation
type BinaryOpNode struct {
	Left     *ASTNode
	Operator Token
	Right    *ASTNode
}

func (b *BinaryOpNode) Type() NodeType { return BinaryOpNodeType }

func (b *BinaryOpNode) String() string { return pretty(b, 0) }

// unary operation
type UnaryOpNode struct {
	Operator Token
	Operand  *ASTNode
}

func (u *UnaryOpNode) Type() NodeType { return UnaryOpNodeType }

func (u *UnaryOpNode) String() string { return pretty(u, 0) }

// number literal
type NumberLiteralNode struct {
	Value float64
}

func (n *NumberLiteralNode) Type() NodeType { return NumberLiteralNodeType }

func (n *NumberLiteralNode) String() string { return pretty(n, 0) }

// string literal
type StringLiteralNode struct {
	Value string
}

func (s *StringLiteralNode) Type() NodeType { return StringLiteralNodeType }

func (s *StringLiteralNode) String() string { return pretty(s, 0) }

// identifier
type IdentifierNode struct {
	Name string
}

func (i *IdentifierNode) Type() NodeType { return IdentifierNodeType }

func (i *IdentifierNode) String() string { return pretty(i, 0) }

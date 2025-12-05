package interpreter

import (
	"pcl/src/frontend/ast"
	"pcl/src/runtime"
	"strconv"
)

func (interpreter *Interpreter) Evaluate(node ast.ASTNode) runtime.RuntimeValue {
	switch node := node.(type) {
	case *ast.ProgramNode:
		return interpreter.evalProgram(node)
	case *ast.BlockNode:
		return interpreter.evalBlock(node)
	case *ast.BinaryOpNode:
		return interpreter.evalBinOp(node)
	case *ast.VarDeclNode:
		return interpreter.evalVarDecl(node)
	case *ast.AssignmentNode:
		return interpreter.evalAssignment(node)
	case *ast.UnaryOpNode:
		return interpreter.evalUnary(node)
	case *ast.IdentifierNode:
		return interpreter.evalIdentifier(node)
	case *ast.FunctionCallNode:
		return interpreter.evalFuncCall(node)
	case *ast.FunctionLiteralNode:
		return interpreter.
	case *ast.LiteralNode[float64]:
		return &runtime.FloatValue{Value: node.Value}
	case *ast.LiteralNode[int]:
		return &runtime.IntValue{Value: node.Value}
	case *ast.LiteralNode[string]:
		return &runtime.StringValue{Value: node.Value}
	default:
		panic("unsupported AST node type: " + strconv.Itoa(int(node.Type())))
	}
}

func (interpreter *Interpreter) evalProgram(programNode *ast.ProgramNode) runtime.RuntimeValue {
	var lastVal runtime.RuntimeValue = &runtime.NilValue{}

	for _, stmt := range programNode.Statements {
		lastVal = interpreter.Evaluate(stmt)
	}

	return lastVal
}

func (interpreter *Interpreter) evalBlock(blockNode *ast.BlockNode) runtime.RuntimeValue {
	var lastVal runtime.RuntimeValue = &runtime.NilValue{}
	interpreter.EnterScope()

	for _, stmt := range blockNode.Statements {
		lastVal = interpreter.Evaluate(stmt)
	}

	interpreter.ExitScope()
	return lastVal
}

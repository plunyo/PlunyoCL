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
		case *ast.BodyNode:
			return interpreter.evalBody(node)
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
			return &runtime.FunctionValue{Arguments: node.Arguments, Body: node.Body}
		case *ast.ReturnNode:
			return &runtime.ReturnValue{Value: interpreter.Evaluate(node.Value)}
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

		if _, ok := lastVal.(*runtime.ReturnValue); ok {
            panic("return outside function")
        }
	}

	return lastVal
}

func (interpreter *Interpreter) evalBody(blockNode *ast.BodyNode) runtime.RuntimeValue {
    interpreter.EnterScope()
    defer interpreter.ExitScope()

	var result runtime.RuntimeValue

    for _, stmt := range blockNode.Statements {
        result = interpreter.Evaluate(stmt)

        if _, ok := result.(*runtime.ReturnValue); ok {
            return result
        }
    }

    return result
}

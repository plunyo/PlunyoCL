package interpreter

import (
	"pcl/src/frontend/ast"
	"pcl/src/runtime"
)

func (interpreter *Interpreter) evalVarDecl(node *ast.VarDeclNode) runtime.RuntimeValue {
	var val runtime.RuntimeValue = &runtime.NilValue{}

	if node.Value != nil {
		val = interpreter.Evaluate(node.Value)
	}

	return interpreter.currentScope.SetVariable(node.Name, val)
}

func (interpreter *Interpreter) evalAssignment(node *ast.AssignmentNode) runtime.RuntimeValue {
	if !interpreter.currentScope.HasVariable(node.Name) {
		return interpreter.currentScope.SetVariable(node.Name, interpreter.Evaluate(node.Value))
	}

	panic("variable not found: " + node.Name)
}

func (interpreter *Interpreter) evalIdentifier(node *ast.IdentifierNode) runtime.RuntimeValue {
	if interpreter.currentScope.HasVariable(node.Name) {
		return interpreter.currentScope.GetVariable(node.Name)
	} 
	
	panic("variable not found: " + node.Name)
}

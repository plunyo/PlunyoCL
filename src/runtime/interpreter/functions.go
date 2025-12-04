package interpreter

import (
	"pcl/src/frontend/ast"
	"pcl/src/runtime"
)

func (interpreter *Interpreter) evalFuncDecl(node *ast.FunctionDeclNode) runtime.RuntimeValue {
	funcValue := &runtime.FunctionValue{
		Name:       node.Name,
		Parameters: node.Arguments,
		Body:       &ast.BlockNode{Statements: node.Statements},
		Closure:    interpreter.currentScope,
	}

	// Store the function in current scope
	interpreter.currentScope.SetVariable(node.Name, funcValue)
	return &runtime.NilValue{}
}

func (interpreter *Interpreter) evalFuncCall(node *ast.FunctionCallNode) runtime.RuntimeValue {
	funcVal := interpreter.Evaluate(node.Callee)

	// Check if it's a function value
	function, ok := funcVal.(*runtime.FunctionValue)
	if !ok {
		panic("cannot call non-function value")
	}

	// Evaluate arguments
	args := make([]runtime.RuntimeValue, len(node.Arguments))
	for i, arg := range node.Arguments {
		args[i] = interpreter.Evaluate(arg)
	}

	// Check argument count
	if len(args) != len(function.Parameters) {
		panic("function expects " + string(rune(len(function.Parameters))) + " arguments, got " + string(rune(len(args))))
	}

	// Save current scope and switch to function's closure
	prevScope := interpreter.currentScope
	interpreter.currentScope = runtime.NewScope(function.Closure)

	// Bind parameters to arguments
	for i, param := range function.Parameters {
		interpreter.currentScope.SetVariable(param, args[i])
	}

	// Execute function body
	result := interpreter.evalBlock(function.Body)

	// Restore scope
	interpreter.currentScope = prevScope

	return result
}

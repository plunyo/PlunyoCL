package interpreter

import (
	"pcl/src/frontend/ast"
	"pcl/src/runtime"
)


func (interpreter *Interpreter) evalFuncCall(node *ast.FunctionCallNode) runtime.RuntimeValue {
	funcVal := interpreter.Evaluate(node.Callee)

	function, ok := funcVal.(*runtime.FunctionValue)
	if !ok {
		panic("cannot call non-function value")
	}

	args := make([]runtime.RuntimeValue, len(node.Arguments))
	for i, arg := range node.Arguments {
		args[i] = interpreter.Evaluate(arg)
	}

	if len(args) != len(function.Arguments) {
		panic("function expects " + string(rune(len(function.Arguments))) + " arguments, got " + string(rune(len(args))))
	}

	prevScope := interpreter.currentScope
	interpreter.currentScope = runtime.NewScope(prevScope)

	// geniusely make the params into variables
	for i, param := range function.Arguments {
		interpreter.currentScope.SetVariable(param, args[i])
	}

	result := interpreter.evalBody(function.Body)

	interpreter.currentScope = prevScope

	return result
}

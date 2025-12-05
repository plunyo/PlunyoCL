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

	if len(args) != len(function.Parameters) {
		panic("function expects " + string(rune(len(function.Parameters))) + " arguments, got " + string(rune(len(args))))
	}

	prevScope := interpreter.currentScope
	interpreter.currentScope = runtime.NewScope(function.Closure)

	// geniusely make the params into variables
	for i, param := range function.Parameters {
		interpreter.currentScope.SetVariable(param, args[i])
	}

	result := interpreter.evalBlock(function.Body)

	interpreter.currentScope = prevScope

	return result
}

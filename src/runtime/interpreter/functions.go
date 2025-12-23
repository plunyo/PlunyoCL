package interpreter

import (
	"pcl/src/frontend/ast"
	"pcl/src/runtime"
	"strconv"
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
        panic("function expects " +
            strconv.Itoa(len(function.Arguments)) +
            " arguments, got " +
            strconv.Itoa(len(args)))
    }

    prevScope := interpreter.currentScope
    interpreter.currentScope = runtime.NewScope(prevScope)

    for i, param := range function.Arguments {
        interpreter.currentScope.SetVariable(param, args[i])
    }

    result := interpreter.evalBody(function.Body)

    interpreter.currentScope = prevScope

    // this part is the money shot brochacho
    if ret, ok := result.(*runtime.ReturnValue); ok {
        return ret.Value
    }

    return result
}


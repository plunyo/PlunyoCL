package interpreter

import (
	"pcl/src/frontend/ast"
	"pcl/src/runtime"
)

func (interpreter *Interpreter) evalUnary(node *ast.UnaryOpNode) runtime.RuntimeValue {
	operand := interpreter.Evaluate(node.Operand)
	switch node.Operator {
	case "-":
		switch num := operand.(type) {
		case *runtime.IntValue:
			return &runtime.IntValue{Value: -num.Value}
		case *runtime.FloatValue:
			return &runtime.FloatValue{Value: -num.Value}
		default:
			panic("cannot negate non-number")
		}
	case "~":
		if num, ok := operand.(*runtime.IntValue); ok {
			return &runtime.IntValue{Value: ^num.Value}
		}
		panic("bitwise not only works on ints")
	default:
		panic("unknown unary operator: " + node.Operator)
	}
}

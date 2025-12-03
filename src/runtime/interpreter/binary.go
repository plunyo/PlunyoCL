package interpreter

import (
	"pcl/src/frontend/ast"
	"pcl/src/runtime"
)

func (interpreter *Interpreter) evalBinOp(binOpNode *ast.BinaryOpNode) runtime.RuntimeValue {
	op := binOpNode.Operator

	// logical
	if op == "&&" || op == "||" {
		return interpreter.evalLogical(binOpNode, op)
	}

	left := interpreter.Evaluate(binOpNode.Left)
	right := interpreter.Evaluate(binOpNode.Right)

	// bitwise
	if op == "&" || op == "|" || op == "^" || op == "<<" || op == ">>" {
		return interpreter.evalBitwise(left, right, op)
	}

	// arithmetic
	if op == "+" || op == "-" || op == "*" || op == "/" || op == "%" {
		return interpreter.evalArithmetic(left, right, op)
	}

	// comparison
	if op == "==" || op == "!=" || op == "<" || op == ">" || op == "<=" || op == ">=" {
		return interpreter.evalComparison(left, right, op)
	}

	panic("unknown binary operator: " + op)
}

func (interpreter *Interpreter) evalLogical(binOpNode *ast.BinaryOpNode, op string) runtime.RuntimeValue {
	left := interpreter.Evaluate(binOpNode.Left)
	lb, ok := left.(*runtime.BooleanValue)
	if !ok {
		panic("left operand of logical operator is not bool")
	}

	if op == "&&" {
		if !lb.Value {
			return &runtime.BooleanValue{Value: false}
		}
		right := interpreter.Evaluate(binOpNode.Right)
		rb, ok := right.(*runtime.BooleanValue)
		if !ok {
			panic("right operand of logical operator is not bool")
		}
		return &runtime.BooleanValue{Value: lb.Value && rb.Value}
	}

	if lb.Value {
		return &runtime.BooleanValue{Value: true}
	}

	right := interpreter.Evaluate(binOpNode.Right)
	rb, ok := right.(*runtime.BooleanValue)
	if !ok {
		panic("right operand of logical operator is not bool")
	}
	return &runtime.BooleanValue{Value: lb.Value || rb.Value}
}

func (interpreter *Interpreter) evalBitwise(left, right runtime.RuntimeValue, op string) runtime.RuntimeValue {
	lInt, lok := left.(*runtime.IntValue)
	rInt, rok := right.(*runtime.IntValue)
	if !lok || !rok {
		panic("bitwise operators only work on ints")
	}

	switch op {
	case "&":
		return &runtime.IntValue{Value: lInt.Value & rInt.Value}
	case "|":
		return &runtime.IntValue{Value: lInt.Value | rInt.Value}
	case "^":
		return &runtime.IntValue{Value: lInt.Value ^ rInt.Value}
	default:
		panic("unsupported bitwise op: " + op)
	}
}

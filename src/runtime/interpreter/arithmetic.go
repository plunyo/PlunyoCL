package interpreter

import (
	"math"
	"pcl/src/runtime"
)

func (interpreter *Interpreter) evalArithmetic(left, right runtime.RuntimeValue, op string) runtime.RuntimeValue {
	lInt, lIsInt := left.(*runtime.IntValue)
	rInt, rIsInt := right.(*runtime.IntValue)

	// int modulo
	if op == "%" && lIsInt && rIsInt {
		if rInt.Value == 0 {
			panic("modulo by zero")
		}
		return &runtime.IntValue{Value: lInt.Value % rInt.Value}
	}

	if op == "/" {
		if rIsInt && rInt.Value == 0 {
			panic("division by zero")
		}
		if rf, ok := right.(*runtime.FloatValue); ok && rf.Value == 0 {
			panic("division by zero")
		}
	}

	lf := interpreter.asFloat(left)
	rf := interpreter.asFloat(right)
	var res float64

	switch op {
	case "+":
		if ls, lok := left.(*runtime.StringValue); lok {
			if rs, rok := right.(*runtime.StringValue); rok {
				return &runtime.StringValue{Value: ls.Value + rs.Value}
			}
		}
		res = lf + rf
	case "-":
		res = lf - rf
	case "*":
		res = lf * rf
	case "/":
		res = lf / rf
	case "%":
		res = math.Mod(lf, rf)
	default:
		panic("unsupported arithmetic op: " + op)
	}

	if lIsInt && rIsInt && res == math.Trunc(res) {
		return &runtime.IntValue{Value: int(res)}
	}
	return &runtime.FloatValue{Value: res}
}

func (interpreter *Interpreter) evalComparison(left, right runtime.RuntimeValue, op string) runtime.RuntimeValue {
	if isNumber(left) && isNumber(right) {
		lf := interpreter.asFloat(left)
		rf := interpreter.asFloat(right)
		switch op {
		case "==":
			return &runtime.BooleanValue{Value: lf == rf}
		case "!=":
			return &runtime.BooleanValue{Value: lf != rf}
		case "<":
			return &runtime.BooleanValue{Value: lf < rf}
		case ">":
			return &runtime.BooleanValue{Value: lf > rf}
		case "<=":
			return &runtime.BooleanValue{Value: lf <= rf}
		case ">=":
			return &runtime.BooleanValue{Value: lf >= rf}
		}
	}

	if ls, lok := left.(*runtime.StringValue); lok {
		if rs, rok := right.(*runtime.StringValue); rok {
			switch op {
			case "==":
				return &runtime.BooleanValue{Value: ls.Value == rs.Value}
			case "!=":
				return &runtime.BooleanValue{Value: ls.Value != rs.Value}
			case "<":
				return &runtime.BooleanValue{Value: ls.Value < rs.Value}
			case ">":
				return &runtime.BooleanValue{Value: ls.Value > rs.Value}
			case "<=":
				return &runtime.BooleanValue{Value: ls.Value <= rs.Value}
			case ">=":
				return &runtime.BooleanValue{Value: ls.Value >= rs.Value}
			}
		}
	}

	switch op {
	case "==":
		return &runtime.BooleanValue{Value: runtimeEqual(left, right)}
	case "!=":
		return &runtime.BooleanValue{Value: !runtimeEqual(left, right)}
	}

	panic("invalid comparison between types for operator: " + op)
}

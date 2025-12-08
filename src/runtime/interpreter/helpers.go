package interpreter

import "pcl/src/runtime"


func (interpreter *Interpreter) asFloat(val runtime.RuntimeValue) float64 {
	switch v := val.(type) {
	case *runtime.IntValue:
		return float64(v.Value)
	case *runtime.FloatValue:
		return v.Value
	default:
		panic("operand is not a number")
	}
}

func isNumber(val runtime.RuntimeValue) bool {
	switch val.(type) {
		case *runtime.IntValue, *runtime.FloatValue:
			return true
		default:
			return false
	}
}

func runtimeEqual(a, b runtime.RuntimeValue) bool {
	switch aa := a.(type) {
	case *runtime.IntValue:
		switch bb := b.(type) {
		case *runtime.IntValue:
			return aa.Value == bb.Value
		case *runtime.FloatValue:
			return float64(aa.Value) == bb.Value
		}
	case *runtime.FloatValue:
		switch bb := b.(type) {
		case *runtime.FloatValue:
			return aa.Value == bb.Value
		case *runtime.IntValue:
			return aa.Value == float64(bb.Value)
		}
	case *runtime.StringValue:
		if bb, ok := b.(*runtime.StringValue); ok {
			return aa.Value == bb.Value
		}
	case *runtime.BooleanValue:
		if bb, ok := b.(*runtime.BooleanValue); ok {
			return aa.Value == bb.Value
		}
	case *runtime.NilValue:
		_, ok := b.(*runtime.NilValue)
		return ok
	}
	return false
}

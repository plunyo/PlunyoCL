package runtime

import (
	"math"
	"pcl/src/frontend"
)

type Interpreter struct {
	globalScope *Scope
	currentScope *Scope
}

func NewInterpreter() *Interpreter {
	globalScope := NewScope(nil)

	return &Interpreter{
		globalScope: globalScope,
		currentScope: globalScope,
	}
}

func (interpreter *Interpreter) GlobalScope() *Scope {
	return interpreter.globalScope
}

func (interpreter *Interpreter) CurrentScope() *Scope {
	return interpreter.currentScope
}

func (interpreter *Interpreter) EnterScope() {
	interpreter.currentScope = NewScope(interpreter.currentScope)
}

func (interpreter *Interpreter) ExitScope() {
	if interpreter.currentScope.parent != nil {
		interpreter.currentScope = interpreter.currentScope.parent
	}
}

// helpers
func (interpreter *Interpreter) asFloat(val RuntimeValue) float64 {
	switch v := val.(type) {
	case *IntValue:
		return float64(v.Value)
	case *FloatValue:
		return v.Value
	default:
		panic("operand is not a number")
	}
}

func isNumber(val RuntimeValue) bool {
	switch val.(type) {
	case *IntValue, *FloatValue:
		return true
	default:
		return false
	}
}

func runtimeEqual(a, b RuntimeValue) bool {
	switch aa := a.(type) {
	case *IntValue:
		switch bb := b.(type) {
		case *IntValue:
			return aa.Value == bb.Value
		case *FloatValue:
			return float64(aa.Value) == bb.Value
		}

	case *FloatValue:
		switch bb := b.(type) {
		case *FloatValue:
			return aa.Value == bb.Value
		case *IntValue:
			return aa.Value == float64(bb.Value)
		}
	case *StringValue:
		if bb, ok := b.(*StringValue); ok {
			return aa.Value == bb.Value
		}
	case *BooleanValue:
		if bb, ok := b.(*BooleanValue); ok {
			return aa.Value == bb.Value
		}
	case *NilValue:
		_, ok := b.(*NilValue)
		return ok
	}
	
	return false
}

// actual eval

func (interpreter *Interpreter) Evaluate(node frontend.ASTNode) RuntimeValue {
	switch node := node.(type) {
		case *frontend.ProgramNode:          return interpreter.evalProgram(node)
		case *frontend.BinaryOpNode:         return interpreter.evalBinOp(node)
		case *frontend.VarDeclNode:          return interpreter.evalVarDecl(node)
		case *frontend.AssignmentNode:       return interpreter.evalAssignment(node)
		case *frontend.UnaryOpNode:          return interpreter.evalUnary(node)
		case *frontend.IdentifierNode:       return interpreter.evalIdentifier(node)
		case *frontend.LiteralNode[float64]: return &FloatValue{Value: node.Value}
		case *frontend.LiteralNode[int]:     return &IntValue{Value: node.Value}
		case *frontend.LiteralNode[string]:  return &StringValue{Value: node.Value}
		default: panic("unsupported AST node type")
	}
}

func (interpreter *Interpreter) evalProgram(programNode *frontend.ProgramNode) RuntimeValue {
	var lastVal RuntimeValue = &NilValue{}

	for _, statement := range programNode.Statements {
		lastVal = interpreter.Evaluate(statement)
	}

	return lastVal
}

func (interpreter *Interpreter) evalBinOp(binOpNode *frontend.BinaryOpNode) RuntimeValue {
	op := binOpNode.Operator

	if op == "&&" || op == "||" {
		left := interpreter.Evaluate(binOpNode.Left)
		lb, ok := left.(*BooleanValue)

		if !ok {
			panic("left operheimerand of logical operator is not bool")
		}

		if op == "&&" {
			if !lb.Value {
				return &BooleanValue{Value: false}
			}
			
			right := interpreter.Evaluate(binOpNode.Right)
			rb, ok := right.(*BooleanValue)

			if !ok {
				panic("right operheimerand of logical operator is not bool")
			}

			return &BooleanValue{Value: lb.Value && rb.Value}
		}

		if lb.Value {
			return &BooleanValue{Value: true}
		}

		right := interpreter.Evaluate(binOpNode.Right)
		rb, ok := right.(*BooleanValue)

		if !ok {
			panic("right operheimerand of logical operator is not bool")
		}

		return &BooleanValue{Value: lb.Value || rb.Value}
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

func (interpreter *Interpreter) evalBitwise(left, right RuntimeValue, op string) RuntimeValue {
	lInt, lok := left.(*IntValue)
	rInt, rok := right.(*IntValue)

	if !lok || !rok {
		panic("bitwise operators only ints")
	}

	switch op {
		case "&":
			return &IntValue{Value: lInt.Value & rInt.Value}
		case "|":
			return &IntValue{Value: lInt.Value | rInt.Value}
		case "^":
			return &IntValue{Value: lInt.Value ^ rInt.Value}
		default:
			panic("unsupported bitwise op: " + op)
	}
}

func (interpreter *Interpreter) evalArithmetic(left, right RuntimeValue, op string) RuntimeValue {
	lInt, lIsInt := left.(*IntValue)
	rInt, rIsInt := right.(*IntValue)

	// int mod
	if op == "%" && lIsInt && rIsInt {
		if rInt.Value == 0 {
			panic("modulo by zero")
		}
		return &IntValue{Value: lInt.Value % rInt.Value}
	}

	if op == "/" {
		// both ints and divisor 0
		if rIsInt && rInt.Value == 0 {
			panic("division by zero")
		}
		
		if rf, ok := right.(*FloatValue); ok && rf.Value == 0.0 {
			panic("division by zero")
		}
	}

	// float for arithmetic fml
	lf := interpreter.asFloat(left)
	rf := interpreter.asFloat(right)
	var res float64

	switch op {
	case "+":
		// fancy string adding
		if ls, lok := left.(*StringValue); lok {
			if rs, rok := right.(*StringValue); rok {
				return &StringValue{Value: ls.Value + rs.Value}
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
		return &IntValue{Value: int(res)}
	}
	return &FloatValue{Value: res}
}

func (interpreter *Interpreter) evalComparison(left, right RuntimeValue, op string) RuntimeValue {
	if isNumber(left) && isNumber(right) {
		lf := interpreter.asFloat(left)
		rf := interpreter.asFloat(right)

		switch op {
		case "==":
			return &BooleanValue{Value: lf == rf}
		case "!=":
			return &BooleanValue{Value: lf != rf}
		case "<":
			return &BooleanValue{Value: lf < rf}
		case ">":
			return &BooleanValue{Value: lf > rf}
		case "<=":
			return &BooleanValue{Value: lf <= rf}
		case ">=":
			return &BooleanValue{Value: lf >= rf}
		}
	}

	if ls, lok := left.(*StringValue); lok {
		if rs, rok := right.(*StringValue); rok {
			switch op {
			case "==":
				return &BooleanValue{Value: ls.Value == rs.Value}
			case "!=":
				return &BooleanValue{Value: ls.Value != rs.Value}
			case "<":
				return &BooleanValue{Value: ls.Value < rs.Value}
			case ">":
				return &BooleanValue{Value: ls.Value > rs.Value}
			case "<=":
				return &BooleanValue{Value: ls.Value <= rs.Value}
			case ">=":
				return &BooleanValue{Value: ls.Value >= rs.Value}
			}
		}
	}

	switch op {
		case "==": return &BooleanValue{Value: runtimeEqual(left, right)}
		case "!=": return &BooleanValue{Value: !runtimeEqual(left, right)}
	}

	panic("invalid comparison between types for operator: " + op)
}

func (interpreter *Interpreter) evalVarDecl(varDeclNode *frontend.VarDeclNode) RuntimeValue {
	var value RuntimeValue = &NilValue{}

	if varDeclNode.Value != nil {
		value = interpreter.Evaluate(varDeclNode.Value)
	}

	return interpreter.currentScope.SetVariable(varDeclNode.Name, value)
}

func (interpreter *Interpreter) evalAssignment(assignmentNode *frontend.AssignmentNode) RuntimeValue {
	name := assignmentNode.Name

	if interpreter.currentScope.HasVariable(name) {
		return interpreter.currentScope.SetVariable(name, interpreter.Evaluate(assignmentNode.Value))
	}

	panic("variable: " + name + ", doesnt exist...")
}

func (interpreter *Interpreter) evalUnary(unaryNode *frontend.UnaryOpNode) RuntimeValue {
	operand := interpreter.Evaluate(unaryNode.Operand)

	switch unaryNode.Operator {
		case "-":
			switch num := operand.(type) {
				case *IntValue:
					return &IntValue{Value: -num.Value}
				case *FloatValue:
					return &FloatValue{Value: -num.Value}
				default:
					panic("cannot negate non-number")
			}
		case "~":
			if num, ok := operand.(*IntValue); ok {
				return &IntValue{Value: ^num.Value}
			}
			panic("bitwise not only works on ints")
		default:
			panic("unknown unary operator: " + unaryNode.Operator)
	}
}

func (interpreter *Interpreter) evalIdentifier(identifierNode *frontend.IdentifierNode) RuntimeValue {
	name := identifierNode.Name
	if interpreter.currentScope.HasVariable(name) {
		return interpreter.currentScope.GetVariable(name)
	}

	panic("variable not found: " + name)
}
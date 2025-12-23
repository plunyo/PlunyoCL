package runtime

import (
	"fmt"
	"pcl/src/frontend/ast"
)

// types
type ValueType int

const (
	IntValueType ValueType = iota
	FloatValueType
	StringValueType
	BooleanValueType
	NilValueType
	FunctionValueType
	ReturnValueType
)

// interface
type RuntimeValue interface {
	Type() ValueType
	String() string
}

// int
type IntValue struct {
	Value int
}

func (v *IntValue) Type() ValueType { return IntValueType }
func (v *IntValue) String() string {
	return fmt.Sprintf("IntValue { Value: %d }", v.Value)
}

// float
type FloatValue struct {
	Value float64
}

func (v *FloatValue) Type() ValueType { return FloatValueType }
func (v *FloatValue) String() string {
	return fmt.Sprintf("FloatValue { Value: %f }", v.Value)
}

// string
type StringValue struct {
	Value string
}

func (v *StringValue) Type() ValueType { return StringValueType }
func (v *StringValue) String() string {
	return fmt.Sprintf("StringValue { Value: %q }", v.Value)
}

// boolean
type BooleanValue struct {
	Value bool
}

func (v *BooleanValue) Type() ValueType { return BooleanValueType }
func (v *BooleanValue) String() string {
	return fmt.Sprintf("BooleanValue { Value: %t }", v.Value)
}

// nil
type NilValue struct{}

func (v *NilValue) Type() ValueType { return NilValueType }
func (v *NilValue) String() string {
	return "NilValue { Value: nil }"
}

// reutrn value
type ReturnValue struct {
	Value RuntimeValue
}

func (r *ReturnValue) Type() ValueType { return ReturnValueType }
func (r *ReturnValue) String() string {
	return fmt.Sprintf("ReturnValue { %s }", r.Value.String())
}

// function
type FunctionValue struct {
	Arguments []string
	Body       *ast.BodyNode
}

func (f *FunctionValue) Type() ValueType { return FunctionValueType }
func (f *FunctionValue) String() string {
	return fmt.Sprintf("FunctionValue { Arguments: %v }", f.Arguments)
}
package runtime

import "strings"

type Scope struct {
	Parent    *Scope
	variables map[string]RuntimeValue
}

func NewScope(parent *Scope) *Scope {
	return &Scope{
		variables: make(map[string]RuntimeValue),
		Parent:    parent,
	}
}

func (scope *Scope) GetVariable(name string) RuntimeValue {
	if val, ok := scope.variables[name]; ok {
		return val
	}

	if scope.Parent != nil {
		return scope.Parent.GetVariable(name)
	}

	panic("variable not found: " + name)
}

func (scope *Scope) SetVariable(name string, value RuntimeValue) RuntimeValue {
	scope.variables[name] = value
	return value
}

func (scope *Scope) HasVariable(name string) bool {
	if _, ok := scope.variables[name]; ok {
		return true
	}

	if scope.Parent != nil {
		return scope.Parent.HasVariable(name)
	}

	return false
}

func (scope *Scope) String() string {
	var builder strings.Builder

	builder.WriteString("Scope {\n")

	for name, value := range scope.variables {
		builder.WriteString("  ")
		builder.WriteString(name)
		builder.WriteString(": ")
		builder.WriteString(value.String())
		builder.WriteByte('\n')
	}

	builder.WriteByte('}')

	return builder.String()
}
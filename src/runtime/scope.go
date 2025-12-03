package runtime

type Scope struct {
	variables map[string]RuntimeValue
	parent    *Scope
}

func NewScope(parent *Scope) *Scope {
	return &Scope{
		variables: make(map[string]RuntimeValue),
		parent:    parent,
	}
}

func (scope *Scope) GetVariable(name string) RuntimeValue {
	if val, ok := scope.variables[name]; ok {
		return val
	}

	if scope.parent != nil {
		return scope.parent.GetVariable(name)
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

	if scope.parent != nil {
		return scope.parent.HasVariable(name)
	}

	return false
}

func (scope *Scope) String() string {
	result := "Scope {\n"
	for name, value := range scope.variables {
		result += "  " + name + ": " + value.String() + "\n"
	}
	result += "}"
	return result
}
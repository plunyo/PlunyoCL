package runtime

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
	result := "Scope {\n"
	for name, value := range scope.variables {
		result += "  " + name + ": " + value.String() + "\n"
	}
	result += "}"
	return result
}
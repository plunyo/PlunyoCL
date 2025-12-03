package interpreter

import "pcl/src/runtime"

type Interpreter struct {
	globalScope  *runtime.Scope
	currentScope *runtime.Scope
}

func NewInterpreter() *Interpreter {
	globalScope := runtime.NewScope(nil)

	// init globalScope
	globalScope.SetVariable("nil", &runtime.NilValue{})
	globalScope.SetVariable("false", &runtime.BooleanValue{Value: false})
	globalScope.SetVariable("true", &runtime.BooleanValue{Value: true})

	return &Interpreter{
		currentScope: globalScope,
		globalScope:  globalScope,
	}
}

// --- scope management ---
func (interpreter *Interpreter) CurrentScope() *runtime.Scope {
	return interpreter.currentScope
}

func (interpreter *Interpreter) EnterScope() {
	interpreter.currentScope = runtime.NewScope(interpreter.currentScope)
}

func (interpreter *Interpreter) ExitScope() {
	if interpreter.currentScope.Parent != nil {
		interpreter.currentScope = interpreter.currentScope.Parent
	}
}

package object

func NewEnclisedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// NewEnvironment return Environment
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

// Environment is struct
type Environment struct {
	store map[string]Object
	outer *Environment
}

// Get return the environment object and return status
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set set map of a environment
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
package object

// NewEnvironment return Environment
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

// Environment is struct
type Environment struct {
	store map[string]Object
}

// Get return the environment object and return status
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}

// Set set map of a environment
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
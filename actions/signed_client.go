package actions

type object interface {
	// todo
	toString() string
}

type ActionWithSign struct {
	Owner    string
	Endpoint string
	Object   object
}

package model

type Sendable interface {
	// id of itself
	GetID() string

	// id of user object
	GetActor() string
	GetObject() string

	//
	ClearContext()
}

type Acceptable interface{
	Sendable
}

type Rejectable interface{
	Sendable
}
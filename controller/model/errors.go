package model

type Error struct {
	ErrorString string `json:"error"`
}

func (e *Error) Error() string {
	return e.ErrorString
}

var (
	Unauthorized = &Error{"The access token is invalid"}
	NotFound     = &Error{"Record not found"}
)

package model

type Register struct {
	Username string `json:"username"`
	Host     string `json:"host"`
	Email    string `json:"email"`
	Passwd   string `json:"passwd"`
}

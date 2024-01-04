package model

type Register struct {
	Username string `json:"username" form:"username"`
	Host     string `json:"host" form:"host"`
	Email    string `json:"email" form:"email"`
	Passwd   string `json:"passwd" form:"passwd"`
}

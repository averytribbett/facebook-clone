package models

type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	HomeTown string `json:"homeTown"`
	Job      string `json:"job"`
	Username string `json:"username"`
}

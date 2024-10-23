package models

type Friend struct {
	User_id       string `json:"userId"`
	Friend_id     string `json:"friendId"`
	Friend_status string `json:"friendStatus"`
}

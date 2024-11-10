package models

type Reply struct {
	PostId           int    `json:"postId"`
	UserId           string `json:"userId"`
	ReplyText        string `json:"replyText"`
	ReplierUsername  string `json:"replierUsername"`
	ReplierFirstName string `json:"replierFirstName"`
	ReplierLastName  string `json:"replierLastName"`
}

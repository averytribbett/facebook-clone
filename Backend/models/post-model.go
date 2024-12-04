package models

type Post struct {
	Id              int    `json:"id"`
	Text            string `json:"text"`
	AuthorId        int    `json:"authorId"`
	AuthorUsername  string `json:"authorUsername"`
	AuthorFirstName string `json:"authorFirstName"`
	AuthorLastName  string `json:"authorLastName"`
	ReplyCount      int    `json:"replyCount"`
	ReactionCount   int    `json:"reactionCount"`
	ReactionByUser  string `json:"reactionByUser"`
}

package models

type Post struct {
	Id              int    `json:"id"`
	Text            string `json:"text"`
	AuthorId        int    `json:"authorId"`
	AuthorFirstName string `json:"authorFirstName"`
	AuthorLastName  string `json:"authorLastName"`
}

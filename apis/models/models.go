package models

type Post struct {
	Id         int32  `json:"id"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	Created_at string `json:"created_at"`
	Author     string `json:"author"`
}

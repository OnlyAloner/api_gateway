package models

type PostType struct {
	Id         int32  `json:"id"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	Created_at string `json:"created_at"`
	Author     string `json:"author"`
}

type PostesType struct {
	Postes []PostType `json:"postes"`
}

type CreateRequest struct {
	Postes PostesType `json:"postes"`
}

type CreateResponse struct {
	Postes PostesType `json:"postes"`
}

type GetRequest struct {
	Id int32 `json:"id"`
}

type GetResponse struct {
	Post PostType `json:"post"`
}

type GetAllRequest struct {
}

type GetAllResponse struct {
	Postes PostesType `json:"postes"`
}

type UpdateRequest struct {
	Id   int32    `json:"id"`
	Post PostType `json:"post"`
}

type UpdateResponse struct {
	Post PostType `json:"post"`
}

type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Reason  string `json:"reason"`
}

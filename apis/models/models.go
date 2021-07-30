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
	Postes PostType `json:"postes"`
}

type CreateResponse struct {
	Postes PostType `json:"postes"`
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

//ResponseSuccess ...
type EmptyResponse struct {
}
type ResponseSuccess struct {
	Metadata interface{}
	Data     interface{}
}

//ResponseError ...
type ResponseError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

//InternalServerError ...
type InternalServerError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

//ValidationError ...
type ValidationError struct {
	Code        string `json:"code"`
	Message     string `json:"message"`
	UserMessage string `json:"unread_message"`
}

type ResponseOK struct {
	Message interface{}
}

type Response struct {
	ID string `json:"id"`
}

// Find query ...
type FindQueryModel struct {
	Page     int64  `json:"page,string"`
	Search   string `json:"search"`
	Active   bool   `json:"active"`
	Inactive bool   `josn:"inactive"`
	Limit    int64  `json:"limit,string"`
	Sort     string `json:"sort" example:"name|asc"`
	Lang     string `json:"lang"`
}

type AuthorizationModel struct {
	Token string `header:"Authorization"`
}

type UserInfo struct {
	ID   string `json:"id"`
	Role string `json:"role"`
}

// SuccessModel ...
type SuccessModel struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ErrorModel ...
type ErrorModel struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

// FileUploadedModel ...
type FileUploadedModel struct {
	Filename string `json:"filename"`
	URL      string `json:"url"`
}

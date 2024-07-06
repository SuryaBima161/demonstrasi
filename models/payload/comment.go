package payload

import uuid "github.com/satori/go.uuid"

type AddComment struct {
	Id_Galery uuid.UUID `json:"id" form:"id"`
	Name      string    `json:"name" form:"name"`
	Comment   string    `json:"comment" form:"comment"`
	Rating    float64   `json:"rating" form:"rating"`
}
type GetCommentRespone struct {
	Name    string `json:"name" form:"name"`
	Comment string `json:"comment" form:"comment"`
}

type ReplyCommentRequest struct {
	ReplyComment string `json:"reply_comment" form:"reply_comment"`
}

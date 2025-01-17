package payload

import uuid "github.com/satori/go.uuid"

type AddComment struct {
	Id_Galery uuid.UUID `json:"id" form:"id"`
	Name      string    `json:"name" form:"name"`
	Comment   string    `json:"comment" form:"comment"`
	Rating    float64   `json:"rating" form:"rating"`
}
type GetCommentRespone struct {
	ID                         uuid.UUID                  `json:"id" form:"id"`
	Name                       string                     `json:"name" form:"name"`
	Comment                    string                     `json:"comment" form:"comment"`
	Rating                     float64                    `json:"rating"`
	Reply                      string                     `json:"reply"`
	Status                     string                     `json:"status"`
	GetGaleryForCommentRespone GetGaleryForCommentRespone `json:"galery"`
}

type GetGaleryForCommentRespone struct {
	Id_Galery   uuid.UUID `json:"idgalery"`
	Judul_foto  string    `json:"judulfoto"`
	Nama_lokasi string    `json:"namalokasi"`
}
type GetCommentValidateRespone struct {
	ID        uuid.UUID `json:"id" form:"id"`
	Id_Galery uuid.UUID `json:"idgalery"`
	Name      string    `json:"name" form:"name"`
	Comment   string    `json:"comment" form:"comment"`
	Rating    float64   `json:"rating"`
	Reply     string    `json:"reply"`
}

type ReplyCommentRequest struct {
	ReplyComment string `json:"reply_comment" form:"reply_comment"`
}

type ValidateComment struct {
	Status string `json:"status" form:"status"`
}

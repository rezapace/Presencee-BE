package model

import (
	"gorm.io/gorm"
)

type Dosen struct {
	gorm.Model
	Nama   string `json:"nama" form:"nama"`
	Email  string `json:"email" form:"email"`
	NIP    string `json:"nip" form:"nip"`
	Image  string `json:"image" form:"image"`
	UserID uint   `json:"user_id" form:"user_id"`
}

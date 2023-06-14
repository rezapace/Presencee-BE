package model

import (
	"gorm.io/gorm"
)

type Mahasiswa struct {
	gorm.Model
	Name       string `json:"name" form:"name"`
	Email      string `json:"email" form:"email"`
	NIM        string `json:"nim" form:"nim"`
	Image      string `json:"image" form:"image"`
	Phone      string `json:"phone" form:"phone"`
	Jurusan    string `json:"jurusan" form:"jurusan"`
	TahunMasuk string `json:"tahun_masuk" form:"tahun_masuk"`
	IPK        string `json:"ipk" form:"ipk"`
	RoomID     string `json:"room_id" form:"room_id"`
	UserID     uint   `json:"user_id" form:"user_id"`
	Absen      Absen  `json:"-" form:"absen"`
}

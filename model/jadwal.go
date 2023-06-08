package model

import (
	"gorm.io/gorm"
)

type Jadwal struct {
	gorm.Model
	Hari         string `json:"hari" form:"hari"`
	MatakuliahID uint   `json:"matakuliahID" form:"matakuliahID"`
	RoomID       uint   `json:"roomID" form:"roomID"`
	Jam          string `json:"jam" form:"jam"`
	Name         string `json:"name" form:"Name"`
	UserID       uint   `json:"user_id" form:"user_id" validate:"required"`
	Absen        Absen  `json:"-" form:"absen"`
}

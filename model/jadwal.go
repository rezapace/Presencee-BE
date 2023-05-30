package model

import (
	"gorm.io/gorm"
)

type Jadwal struct {
	gorm.Model
	Hari      string `json:"hari" form:"hari"`
	Matkul    string `json:"matkul" form:"matkul"`
	Ruangan   string `json:"ruangan" form:"ruangan"`
	Jam       string `json:"jam" form:"jam"`
	NamaDosen string `json:"nama_dosen" form:"nama_dosen"`
}

package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string    `json:"name" form:"name"`
	Email     string    `json:"email" form:"email"`
	Password  string    `json:"password" form:"password"`
	Token     string    `gorm:"-"`
	Dosen     Dosen     `gorm:"foreignKey:User_id"`
	Mahasiswa Mahasiswa `gorm:"foreignKey:User_id"`
}

type Dosen struct {
	gorm.Model
	nama    string `json:"nama" form:"nama"`
	email   string `json:"email" form:"email"`
	nip     string `json:"nip" form:"nip"`
	image   string `json:"image" form:"image"`
	user_id uint   `json:"user_id" form:"user_id"`
}

type Mahasiswa struct {
	gorm.Model
	mahasiswa   string `json:"mahasiswa" form:"mahasiswa"`
	email       string `json:"email" form:"email"`
	name        string `json:"name" form:"name"`
	nim         string `json:"nim" form:"nim"`
	image       string `json:"image" form:"image"`
	phone       string `json:"phone" form:"phone"`
	jurusan     string `json:"jurusan" form:"jurusan"`
	tahun_masuk string `json:"tahun_masuk" form:"tahun_masuk"`
	ipk         string `json:"ipk" form:"ipk"`
	user_id     uint   `json:"user_id" form:"user_id"`
}

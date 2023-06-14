package model

import (
	"time"

	"gorm.io/gorm"
)

type Matakuliah struct {
	gorm.Model
	Name   string    `json:"name" form:"name"`
	Date   time.Time `json:"date" form:"date"`
	Jadwal Jadwal    `json:"-" form:"jadwal"`
}

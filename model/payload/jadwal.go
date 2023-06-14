package payload

import "presensee_project/model"

type CreateJadwalRequest struct {
	Hari         string `json:"hari" validate:"required"`
	MatakuliahID uint   `json:"matakuliah_id" validate:"required"`
	RoomID       uint   `json:"room_id" validate:"required"`
	Jam          string `json:"jam" validate:"required"`
	Name         string `json:"name" validate:"required"`
	UserID       uint   `json:"user_id" form:"user_id" validate:"required"`
	Sks          uint   `json:"sks" form:"sks" validate:"required"`
	Date         string `json:"date" form:"date"`
}

type CreateJadwalResponse struct {
	JadwalID uint `json:"jadwal_id"`
}

type UpdateJadwalRequest struct {
	Hari         string `json:"hari" validate:"required"`
	MatakuliahID uint   `json:"matakuliah_id" validate:"required"`
	RoomID       uint   `json:"room_id" validate:"required"`
	Jam          string `json:"jam" validate:"required"`
	Name         string `json:"name" validate:"required"`
	UserID       uint   `json:"user_id" form:"user_id" validate:"required"`
	Sks          uint   `json:"sks" form:"sks" validate:"required"`
	Date         string `json:"date" form:"date"`
}

type UpdateJadwalResponse struct {
	JadwalID uint `json:"jadwal_id"`
}

type GetJadwalResponse struct {
	JadwalID   uint             `json:"jadwal_id"`
	Hari       string           `json:"hari"`
	Matkul     string           `json:"matkul"`
	Ruangan    string           `json:"ruangan"`
	Jam        string           `json:"jam"`
	NamaDosen  string           `json:"nama_dosen"`
	Matakuliah model.Matakuliah `json:"matakuliah"`
	Room       model.Room       `json:"room"`
	Sks        uint             `json:"sks"`
}

type GetJadwalsResponse struct {
	Jadwals []GetJadwalResponse `json:"jadwals"`
}

type GetListJadwalsByDateResponse struct {
	Jadwals []GetJadwalResponse `json:"jadwals"`
}

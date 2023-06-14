package payload

import (
	"time"

	"presensee_project/model"
)

type CreateAbsenRequest struct {
	UserID       uint      `json:"user_id" validate:"required"`
	MahasiswaID  uint      `json:"mahasiswa_id" validate:"required"`
	JadwalID     uint      `json:"jadwal_id" validate:"required"`
	Matakuliah   string    `json:"matakuliah" validate:"required"`
	TimeAttemp   time.Time `json:"time_attemp" validate:"required"`
	Status       string    `json:"status" validate:"required"`
	Description  string    `json:"description" validate:"required"`
	Location     string    `json:"location" validate:"required"`
	Image        string    `json:"image" validate:"required"`
	IsKonfirmasi bool      `json:"is_konfirmasi" validate:"required"`
}

func (u *CreateAbsenRequest) ToEntity() *model.Absen {
	return &model.Absen{
		UserID:       u.UserID,
		MahasiswaID:  u.MahasiswaID,
		JadwalID:     u.JadwalID,
		Matakuliah:   u.Matakuliah,
		TimeAttemp:   u.TimeAttemp,
		Status:       u.Status,
		Description:  u.Description,
		Location:     u.Location,
		Image:        u.Image,
		IsKonfirmasi: u.IsKonfirmasi,
	}
}

type UpdateAbsenRequest struct {
	ID           uint      `json:"id"`
	UserID       uint      `json:"user_id"`
	MahasiswaID  uint      `json:"mahasiswa_id"`
	JadwalID     uint      `json:"jadwal_id"`
	Matakuliah   string    `json:"matakuliah"`
	TimeAttemp   time.Time `json:"time_attemp"`
	Status       string    `json:"status"`
	Description  string    `json:"description"`
	Location     string    `json:"location"`
	Image        string    `json:"image"`
	IsKonfirmasi bool      `json:"is_konfirmasi"`
}

func (u *UpdateAbsenRequest) ToEntity() *model.Absen {
	return &model.Absen{
		UserID:       u.UserID,
		MahasiswaID:  u.MahasiswaID,
		JadwalID:     u.JadwalID,
		Matakuliah:   u.Matakuliah,
		TimeAttemp:   u.TimeAttemp,
		Status:       u.Status,
		Description:  u.Description,
		Location:     u.Location,
		Image:        u.Image,
		IsKonfirmasi: u.IsKonfirmasi,
	}
}

type GetSingleAbsenResponse struct {
	ID           uint            `json:"id"`
	UserID       uint            `json:"user_id"`
	MahasiswaID  uint            `json:"mahasiswa_id"`
	JadwalID     uint            `json:"jadwal_id"`
	Matakuliah   string          `json:"matakuliah"`
	TimeAttemp   time.Time       `json:"time_attemp"`
	Status       string          `json:"status"`
	Description  string          `json:"description"`
	Location     string          `json:"location"`
	Image        string          `json:"image"`
	IsKonfirmasi bool            `json:"is_konfirmasi"`
	Mahasiswa    model.Mahasiswa `json:"mahasiswa"`
	Jadwal       model.Jadwal    `json:"jadwal"`
}

func NewGetSingleAbsenResponse(absen *model.Absen) *GetSingleAbsenResponse {
	return &GetSingleAbsenResponse{
		ID:           absen.ID,
		UserID:       absen.UserID,
		MahasiswaID:  absen.MahasiswaID,
		JadwalID:     absen.JadwalID,
		Matakuliah:   absen.Matakuliah,
		TimeAttemp:   absen.TimeAttemp,
		Status:       absen.Status,
		Description:  absen.Description,
		Location:     absen.Location,
		Image:        absen.Image,
		IsKonfirmasi: absen.IsKonfirmasi,
	}
}

type GetPageAbsenResponse struct {
	ID           uint      `json:"id"`
	UserID       uint      `json:"user_id"`
	MahasiswaID  uint      `json:"mahasiswa_id"`
	JadwalID     uint      `json:"jadwal_id"`
	Matakuliah   string    `json:"matakuliah"`
	TimeAttemp   time.Time `json:"time_attemp"`
	Status       string    `json:"status"`
	Description  string    `json:"description"`
	Location     string    `json:"location"`
	Image        string    `json:"image"`
	IsKonfirmasi bool      `json:"is_konfirmasi"`
}

func NewGetPageAbsenResponse(absen *model.Absen) *GetPageAbsenResponse {
	return &GetPageAbsenResponse{
		ID:           absen.ID,
		UserID:       absen.UserID,
		MahasiswaID:  absen.MahasiswaID,
		JadwalID:     absen.JadwalID,
		Matakuliah:   absen.Matakuliah,
		TimeAttemp:   absen.TimeAttemp,
		Status:       absen.Status,
		Description:  absen.Description,
		Location:     absen.Location,
		Image:        absen.Image,
		IsKonfirmasi: absen.IsKonfirmasi,
	}
}

type GetPageAbsensResponse []GetPageAbsenResponse

func NewGetPageAbsensResponse(absens *model.Absens) *GetPageAbsensResponse {
	var briefAbsensResponse GetPageAbsensResponse
	for _, absen := range *absens {
		briefAbsensResponse = append(briefAbsensResponse, *NewGetPageAbsenResponse(&absen))
	}
	return &briefAbsensResponse
}

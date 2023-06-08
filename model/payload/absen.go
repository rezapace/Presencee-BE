package payload

import (
	"time"

	"presensee_project/model"
)

type CreateAbsenRequest struct {
	UserID      uint      `json:"user_id" validate:"required"`
	MahasiswaID uint      `json:"mahasiswa_id" validate:"required"`
	JadwalID    uint      `json:"jadwal_id" validate:"required"`
	TimeAttemp  time.Time `json:"time_attemp" validate:"required"`
	Status      string    `json:"status" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Location    string    `json:"location" validate:"required"`
	Image       string    `json:"image" validate:"required"`
}

func (u *CreateAbsenRequest) ToEntity() *model.Absen {
	return &model.Absen{
		UserID:      u.UserID,
		MahasiswaID: u.MahasiswaID,
		JadwalID:    u.JadwalID,
		TimeAttemp:  u.TimeAttemp,
		Status:      u.Status,
		Description: u.Description,
		Location:    u.Location,
		Image:       u.Image,
	}
}

type UpdateAbsenRequest struct {
	ID          uint      `json:"id"`
	UserID      uint      `json:"user_id"`
	MahasiswaID uint      `json:"mahasiswa_id"`
	JadwalID    uint      `json:"jadwal_id"`
	TimeAttemp  time.Time `json:"time_attemp"`
	Status      string    `json:"status"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Image       string    `json:"image"`
}

func (u *UpdateAbsenRequest) ToEntity() *model.Absen {
	return &model.Absen{
		UserID:      u.UserID,
		MahasiswaID: u.MahasiswaID,
		JadwalID:    u.JadwalID,
		TimeAttemp:  u.TimeAttemp,
		Status:      u.Status,
		Description: u.Description,
		Location:    u.Location,
		Image:       u.Image,
	}
}

type GetSingleAbsenResponse struct {
	ID          uint            `json:"id"`
	UserID      uint            `json:"user_id"`
	MahasiswaID uint            `json:"mahasiswa_id"`
	JadwalID    uint            `json:"jadwal_id"`
	TimeAttemp  time.Time       `json:"time_attemp"`
	Status      string          `json:"status"`
	Description string          `json:"description"`
	Location    string          `json:"location"`
	Image       string          `json:"image"`
	Mahasiswa   model.Mahasiswa `json:"mahasiswa"`
	Jadwal      model.Jadwal    `json:"jadwal"`
}

func NewGetSingleAbsenResponse(absen *model.Absen) *GetSingleAbsenResponse {
	return &GetSingleAbsenResponse{
		ID:          absen.ID,
		UserID:      absen.UserID,
		MahasiswaID: absen.MahasiswaID,
		JadwalID:    absen.JadwalID,
		TimeAttemp:  absen.TimeAttemp,
		Status:      absen.Status,
		Description: absen.Description,
		Location:    absen.Location,
		Image:       absen.Image,
	}
}

type GetPageAbsenResponse struct {
	ID          uint      `json:"id"`
	UserID      uint      `json:"user_id"`
	MahasiswaID uint      `json:"mahasiswa_id"`
	JadwalID    uint      `json:"jadwal_id"`
	TimeAttemp  time.Time `json:"time_attemp"`
	Status      string    `json:"status"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Image       string    `json:"image"`
}

func NewGetPageAbsenResponse(absen *model.Absen) *GetPageAbsenResponse {
	return &GetPageAbsenResponse{
		ID:          absen.ID,
		UserID:      absen.UserID,
		MahasiswaID: absen.MahasiswaID,
		JadwalID:    absen.JadwalID,
		TimeAttemp:  absen.TimeAttemp,
		Status:      absen.Status,
		Description: absen.Description,
		Location:    absen.Location,
		Image:       absen.Image,
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

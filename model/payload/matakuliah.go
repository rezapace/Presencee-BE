package payload

import "time"

type CreateMatakuliahRequest struct {
	Name string    `json:"name" form:"name"`
	Date time.Time `json:"date" form:"date"`
}

type CreateMatakuliahResponse struct {
	MatakuliahID uint `json:"matakuliah_id"`
}

type UpdateMatakuliahRequest struct {
	Name string    `json:"name" form:"name"`
	Date time.Time `json:"date" form:"date"`
}

type UpdateMatakuliahResponse struct {
	MatakuliahID uint `json:"matakuliah_id"`
}

type GetMatakuliahResponse struct {
	MatakuliahID uint      `json:"matakuliah_id"`
	Name         string    `json:"name" form:"name"`
	Date         time.Time `json:"date" form:"date"`
}

type GetMatakuliahsResponse struct {
	Matakuliahs []GetMatakuliahResponse `json:"matakuliahs"`
}

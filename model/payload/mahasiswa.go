package payload

type CreateMahasiswaRequest struct {
	Name   string `json:"name" form:"name" validate:"required"`
	Email  string `json:"email" form:"email" validate:"required,email"`
	NIM    string `json:"nim" form:"nim" validate:"required"`
	Image  string `json:"image" form:"image"`
	UserID uint   `json:"user_id" form:"user_id" validate:"required"`
}

type CreateMahasiswaResponse struct {
	MahasiswaID uint `json:"mahasiswa_id"`
}

type UpdateMahasiswaRequest struct {
	Name   string `json:"name" form:"name" validate:"required"`
	Email  string `json:"email" form:"email" validate:"required,email"`
	NIM    string `json:"nim" form:"nim" validate:"required"`
	Image  string `json:"image" form:"image"`
	UserID uint   `json:"user_id" form:"user_id" validate:"required"`
}

type UpdateMahasiswaResponse struct {
	MahasiswaID uint `json:"mahasiswa"`
}

type GetMahasiswaResponse struct {
	MahasiswaID uint   `json:"mahasiswa_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	NIM         string `json:"nim"`
	Image       string `json:"image"`
	UserID      uint   `json:"user_id"`
}

type GetMahasiswasResponse struct {
	Mahasiswas []GetMahasiswaResponse `json:"mahasiswas"`
}

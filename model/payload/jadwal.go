package payload

type CreateJadwalRequest struct {
	Hari      string `json:"hari" validate:"required"`
	Matkul    string `json:"matkul" validate:"required"`
	Ruangan   string `json:"ruangan" validate:"required"`
	Jam       string `json:"jam" validate:"required"`
	NamaDosen string `json:"nama_dosen" validate:"required"`
}

type CreateJadwalResponse struct {
	JadwalID uint `json:"jadwal_id"`
}

type UpdateJadwalRequest struct {
	ID        uint   `json:"id" validate:"required"`
	Hari      string `json:"hari" validate:"required"`
	Matkul    string `json:"matkul" validate:"required"`
	Ruangan   string `json:"ruangan" validate:"required"`
	Jam       string `json:"jam" validate:"required"`
	NamaDosen string `json:"nama_dosen" validate:"required"`
}

type UpdateJadwalResponse struct {
	JadwalID uint `json:"jadwal_id"`
}

type GetJadwalResponse struct {
	JadwalID  uint   `json:"jadwal_id"`
	Hari      string `json:"hari"`
	Matkul    string `json:"matkul"`
	Ruangan   string `json:"ruangan"`
	Jam       string `json:"jam"`
	NamaDosen string `json:"nama_dosen"`
}

type GetJadwalsResponse struct {
	Jadwals []GetJadwalResponse `json:"jadwals"`
}

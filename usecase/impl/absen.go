package usecase

import (
	"context"

	"presensee_project/model/payload"
	"presensee_project/repository"
	database "presensee_project/repository/impl"
	"presensee_project/usecase"

	"github.com/google/uuid"
)

type (
	AbsenServiceImpl struct {
		absenRepository repository.AbsenRepository
	}
)

func NewAbsenServiceImpl(absenRepository repository.AbsenRepository) usecase.AbsenService {
	return &AbsenServiceImpl{
		absenRepository: absenRepository,
	}
}

func (u *AbsenServiceImpl) CreateAbsen(ctx context.Context, absen *payload.CreateAbsenRequest) error {

	absenEntity := absen.ToEntity()
	absenEntity.ID = uint(uuid.New().ID())

	err := u.absenRepository.CreateAbsen(ctx, absenEntity)
	if err != nil {
		return err
	}

	return nil
}

func (d *AbsenServiceImpl) GetSingleAbsen(ctx context.Context, absenID uint) (*payload.GetSingleAbsenResponse, error) {
	absen, err := d.absenRepository.GetSingleAbsen(ctx, absenID)
	if err != nil {
		return nil, err
	}

	mahasiswa, err := database.GetMahasiswaByID(absen.MahasiswaID)
	if err != nil {
		return nil, err
	}

	jadwal, err := database.GetJadwalByID(absen.JadwalID)
	if err != nil {
		return nil, err
	}
	var absenResponse = payload.NewGetSingleAbsenResponse(absen)
	absenResponse.Mahasiswa = mahasiswa
	absenResponse.Jadwal = jadwal

	return absenResponse, nil
}

func (u *AbsenServiceImpl) GetPageAbsens(ctx context.Context, page int, limit int) (*payload.GetPageAbsensResponse, error) {
	offset := (page - 1) * limit

	absens, err := u.absenRepository.GetPageAbsens(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return payload.NewGetPageAbsensResponse(absens), nil
}

func (u *AbsenServiceImpl) UpdateAbsen(ctx context.Context, absenID uint, updateAbsen *payload.UpdateAbsenRequest) error {
	absen := updateAbsen.ToEntity()
	absen.ID = absenID

	return u.absenRepository.UpdateAbsen(ctx, absen)
}

func (d *AbsenServiceImpl) DeleteAbsen(ctx context.Context, absenID uint) error {

	return d.absenRepository.DeleteAbsen(ctx, absenID)
}

package repository

import (
	"context"

	"presensee_project/model"
)

type AbsenRepository interface {
	CreateAbsen(ctx context.Context, absen *model.Absen) error
	GetSingleAbsen(ctx context.Context, absenID uint) (*model.Absen, error)
	GetPageAbsens(ctx context.Context, limit int, offset int) (*model.Absens, error)
	UpdateAbsen(ctx context.Context, absen *model.Absen) error
	DeleteAbsen(ctx context.Context, absenID uint) error
}

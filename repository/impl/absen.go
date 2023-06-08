package database

import (
	"context"
	"strings"

	"presensee_project/model"
	"presensee_project/repository"
	"presensee_project/utils"

	"gorm.io/gorm"
)

type AbsenRepositoryImpl struct {
	db *gorm.DB
}

func NewAbsenRepositoryImpl(db *gorm.DB) repository.AbsenRepository {
	absenRepository := &AbsenRepositoryImpl{
		db: db,
	}

	return absenRepository
}

func (u *AbsenRepositoryImpl) CreateAbsen(ctx context.Context, absen *model.Absen) error {
	err := u.db.WithContext(ctx).Create(absen).Error
	if err != nil {
		if strings.Contains(err.Error(), "Error 1062: Duplicate entry") {
			switch {
			case strings.Contains(err.Error(), "name"):
				return utils.ErrItemAlreadyExist
			}
		}

		return err
	}

	return nil
}

func (u *AbsenRepositoryImpl) GetSingleAbsen(ctx context.Context, absenID uint) (*model.Absen, error) {
	var absen model.Absen

	err := u.db.WithContext(ctx).
		Where("id = ?", absenID).
		First(&absen).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, utils.ErrAbsenNotFound
		}

		return nil, err
	}

	return &absen, nil
}

func (u *AbsenRepositoryImpl) GetPageAbsens(ctx context.Context, limit int, offset int) (*model.Absens, error) {
	var absens model.Absens
	err := u.db.WithContext(ctx).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&absens).Error
	if err != nil {
		return nil, err
	}

	if len(absens) == 0 {
		return nil, utils.ErrAbsenNotFound
	}

	return &absens, nil
}

func (u *AbsenRepositoryImpl) UpdateAbsen(ctx context.Context, absen *model.Absen) error {
	result := u.db.WithContext(ctx).Model(&model.Absen{}).Where("id = ?", absen.ID).Updates(absen)
	if result.Error != nil {
		errStr := result.Error.Error()
		if strings.Contains(errStr, "Error 1062: Duplicate entry") {
			switch {
			case strings.Contains(errStr, "name"):
				return utils.ErrItemAlreadyExist
			}
		}

		return result.Error
	}

	if result.RowsAffected == 0 {
		return utils.ErrAbsenNotFound
	}

	return nil
}

func (d *AbsenRepositoryImpl) DeleteAbsen(ctx context.Context, absenID uint) error {
	result := d.db.WithContext(ctx).
		Select("Absen").
		Delete(&model.Absen{}, "id = ?", absenID)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return utils.ErrAbsenNotFound
	}

	return nil
}

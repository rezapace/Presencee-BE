package usecase

import (
	"presensee_project/model"
	"presensee_project/repository"
)

func CreateJadwal(jadwal *model.Jadwal) error {
	err := repository.CreateJadwal(jadwal)
	if err != nil {
		return err
	}
	return nil
}

func GetJadwal(id uint) (model.Jadwal, error) {
	jadwal, err := repository.GetJadwalByID(id)
	if err != nil {
		return model.Jadwal{}, err
	}
	return jadwal, nil
}

func GetListJadwals() ([]model.Jadwal, error) {
	jadwals, err := repository.GetJadwals()
	if err != nil {
		return nil, err
	}
	return jadwals, nil
}

func UpdateJadwal(jadwal *model.Jadwal) error {
	err := repository.UpdateJadwal(jadwal)
	if err != nil {
		return err
	}
	return nil
}

func DeleteJadwal(id uint) error {
	err := repository.DeleteJadwal(id)
	if err != nil {
		return err
	}
	return nil
}

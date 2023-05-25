package usecase

import (
	"presensee_project/model"
	"presensee_project/repository/database"
)

func CreateMahasiswa(mahasiswa *model.Mahasiswa) error {
	err := database.CreateMahasiswa(mahasiswa)
	if err != nil {
		return err
	}
	return nil
}

func GetMahasiswa(id uint) (model.Mahasiswa, error) {
	mahasiswa, err := database.GetMahasiswaByID(id)
	if err != nil {
		return model.Mahasiswa{}, err
	}
	return mahasiswa, nil
}

func GetListMahasiswas() ([]model.Mahasiswa, error) {
	mahasiswas, err := database.GetMahasiswas()
	if err != nil {
		return nil, err
	}
	return mahasiswas, nil
}

func UpdateMahasiswa(mahasiswa *model.Mahasiswa) error {
	err := database.UpdateMahasiswa(mahasiswa)
	if err != nil {
		return err
	}
	return nil
}

func DeleteMahasiswa(id uint) error {
	err := database.DeleteMahasiswa(id)
	if err != nil {
		return err
	}
	return nil
}
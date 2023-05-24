package controller

import (
	"net/http"
	"presensee_project/model"
	"presensee_project/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetDosensController mengembalikan semua data dosen
func GetDosensController(c echo.Context) error {
	dosens, err := usecase.GetListDosens()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"dosens": dosens,
	})
}

// GetDosenController mengembalikan data dosen berdasarkan ID
func GetDosenController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	dosen, err := usecase.GetDosen(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"dosen":  dosen,
	})
}

// CreateDosenController membuat data dosen baru
func CreateDosenController(c echo.Context) error {
	payload := new(model.Dosen)
	if err := c.Bind(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err := usecase.CreateDosen(payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": "success",
		"dosen":  payload,
	})
}

// UpdateDosenController mengubah data dosen berdasarkan ID
func UpdateDosenController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	dosen, err := usecase.GetDosen(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	updatedDosen := new(model.Dosen) // Buat variabel pointer untuk menampung data yang di-bind
	if err := c.Bind(updatedDosen); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Perbarui data dosen yang telah di-bind
	dosen.Nama = updatedDosen.Nama
	dosen.Email = updatedDosen.Email
	// Tambahkan perubahan lain sesuai struktur model Dosen

	err = usecase.UpdateDosen(&dosen) // Ubah menjadi &dosen untuk menggunakan pointer
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"dosen":  dosen,
	})
}

// DeleteDosenController menghapus data dosen berdasarkan ID
func DeleteDosenController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	err = usecase.DeleteDosen(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Dosen deleted successfully",
	})
}

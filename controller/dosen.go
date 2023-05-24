package controller

import (
	"net/http"
	"presensee_project/model"
	"presensee_project/repository/database"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetDosenController mengembalikan semua data dosen
func GetDosenController(c echo.Context) error {
	dosens, err := database.GetDosens()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"dosens": dosens,
	})
}

// GetDosenByIDController mengembalikan data dosen berdasarkan ID
func GetDosenByIDController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	dosen, err := database.GetDosenByID(uint(id))
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
	dosen := new(model.Dosen)
	if err := c.Bind(dosen); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err := database.CreateDosen(dosen)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
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
	err = database.DeleteDosen(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Dosen deleted successfully",
	})
}

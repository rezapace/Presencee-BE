package controller

import (
	"net/http"
	"presensee_project/model"
	"presensee_project/model/payload"
	usecase "presensee_project/usecase/impl"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetJadwalsController returns all jadwal data
func GetJadwalsController(c echo.Context) error {
	jadwals, err := usecase.GetListJadwals()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"jadwals": jadwals,
	})
}

// GetJadwalController returns jadwal data based on ID
func GetJadwalController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	jadwal, err := usecase.GetJadwal(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"jadwal": jadwal,
	})
}

// CreateJadwalController creates a new jadwal
func CreateJadwalController(c echo.Context) error {
	requestPayload := new(payload.CreateJadwalRequest)

	// Bind and validate the payload
	if err := c.Bind(requestPayload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(requestPayload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	jadwal := &model.Jadwal{
		Hari:      requestPayload.Hari,
		Matkul:    requestPayload.Matkul,
		Ruangan:   requestPayload.Ruangan,
		Jam:       requestPayload.Jam,
		NamaDosen: requestPayload.NamaDosen,
	}

	err := usecase.CreateJadwal(jadwal)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	responsePayload := &payload.CreateJadwalResponse{
		JadwalID: jadwal.ID,
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": "success",
		"jadwal": responsePayload,
	})
}

// UpdateJadwalController updates jadwal data based on ID
func UpdateJadwalController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	jadwalToUpdate, err := usecase.GetJadwal(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	updatedJadwal := new(payload.UpdateJadwalRequest)

	// Bind and validate the payload
	if err := c.Bind(updatedJadwal); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(updatedJadwal); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Update jadwal data
	jadwalToUpdate.Hari = updatedJadwal.Hari
	jadwalToUpdate.Matkul = updatedJadwal.Matkul
	jadwalToUpdate.Ruangan = updatedJadwal.Ruangan
	jadwalToUpdate.Jam = updatedJadwal.Jam
	jadwalToUpdate.NamaDosen = updatedJadwal.NamaDosen

	err = usecase.UpdateJadwal(&jadwalToUpdate) // Pass the pointer to jadwalToUpdate
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := &payload.UpdateJadwalResponse{
		JadwalID: jadwalToUpdate.ID,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"jadwal": response,
	})
}

// DeleteJadwalController deletes jadwal data based on ID
func DeleteJadwalController(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	err = usecase.DeleteJadwal(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Jadwal deleted successfully",
	})
}
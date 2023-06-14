package controller

import (
	"net/http"
	"presensee_project/model"
	"presensee_project/model/payload"
	usecase "presensee_project/usecase/impl"
	"strconv"
	"time"

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
		Hari:         requestPayload.Hari,
		MatakuliahID: requestPayload.MatakuliahID,
		RoomID:       requestPayload.RoomID,
		Jam:          requestPayload.Jam,
		Name:         requestPayload.Name,
		UserID:       requestPayload.UserID,
		Sks:          requestPayload.Sks,
		Date:         requestPayload.Date,
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
	jadwalToUpdate.MatakuliahID = updatedJadwal.MatakuliahID
	jadwalToUpdate.RoomID = updatedJadwal.RoomID
	jadwalToUpdate.Jam = updatedJadwal.Jam
	jadwalToUpdate.Name = updatedJadwal.Name
	jadwalToUpdate.UserID = updatedJadwal.UserID
	jadwalToUpdate.Sks = updatedJadwal.Sks
	jadwalToUpdate.Date = updatedJadwal.Date

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

// GetListJadwalsByDateController returns jadwal data based on the specified date
func GetListJadwalsByDateController(c echo.Context) error {
	// Get the date parameter from the request query string
	dateStr := c.QueryParam("date")

	// Parse the date string into a time.Time object
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid date format. Please use 'YYYY-MM-DD' format.")
	}

	// Convert the date to string format
	dateString := date.Format("2006-01-02")

	// Get the jadwals for the specified date
	jadwals, err := usecase.GetListJadwalsByDate(dateString)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"jadwals": jadwals,
	})
}

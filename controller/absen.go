package controller

import (
	"net/http"

	"presensee_project/model/payload"
	"presensee_project/usecase"
	"presensee_project/utils"
	"presensee_project/utils/jwt_service"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type AbsenController struct {
	absenService usecase.AbsenService
	jwtService   jwt_service.JWTService
}

func NewAbsenController(absenService usecase.AbsenService, jwtService jwt_service.JWTService) *AbsenController {
	return &AbsenController{
		absenService: absenService,
		jwtService:   jwtService,
	}
}

func (u *AbsenController) CreateAbsen(c echo.Context) error {
	absen := new(payload.CreateAbsenRequest)
	if err := c.Bind(absen); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrBadRequestBody.Error())
	}

	if err := c.Validate(absen); err != nil {
		return err
	}

	err := u.absenService.CreateAbsen(c.Request().Context(), absen)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "success creating absen",
	})
}

func (u *AbsenController) UpdateAbsen(c echo.Context) error {
	claims := u.jwtService.GetClaims(&c)
	role := claims["role"].(string)
	if role != "admin" {
		return echo.NewHTTPError(http.StatusForbidden, utils.ErrDidntHavePermission.Error())
	}
	absen := new(payload.UpdateAbsenRequest)
	if err := c.Bind(absen); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrBadRequestBody.Error())
	}

	if err := c.Validate(absen); err != nil {
		return err
	}

	err := u.absenService.UpdateAbsen(c.Request().Context(), absen.ID, absen)
	if err != nil {
		switch err {
		case utils.ErrAbsenNotFound:
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update absen",
	})
}

func (u *AbsenController) GetSingleAbsen(c echo.Context) error {
	absenIDParam := c.Param("absen_id")
	absenID64, err := strconv.ParseUint(absenIDParam, 10, 0)
	absenID := uint(absenID64)

	absen, err := u.absenService.GetSingleAbsen(c.Request().Context(), absenID)
	if err != nil {
		if err == utils.ErrAbsenNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	claims := u.jwtService.GetClaims(&c)
	role := claims["role"].(string)

	switch {
	case role == "pegawai":
		fallthrough
	case role == "admin" || role == "Mahasiswa" || role == "Dosen":
		return c.JSON(http.StatusOK, echo.Map{
			"message": "success getting user",
			"data":    absen,
		})
	default:
		return echo.NewHTTPError(http.StatusForbidden, utils.ErrDidntHavePermission.Error())
	}
}

func (u *AbsenController) GetPageAbsen(c echo.Context) error {
	page := c.QueryParam("page")
	if page == "" {
		page = "1"
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrInvalidNumber.Error())
	}

	limit := c.QueryParam("limit")
	if limit == "" {
		limit = "20"
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrInvalidNumber.Error())
	}

	userIDStr := c.QueryParam("user_id")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user_id")
	}

	mahasiswaIDStr := c.QueryParam("mahasiswa_id")
	mahasiswaID, err := strconv.ParseUint(mahasiswaIDStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid mahasiswa_id")
	}

	jadwalIDStr := c.QueryParam("jadwal_id")
	jadwalID, err := strconv.ParseUint(jadwalIDStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid jadwal_id")
	}

	absenIDStr := c.QueryParam("absen_id")
	absenID, err := strconv.ParseUint(absenIDStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid absen_id")
	}

	createdAfterStr := c.QueryParam("created_after")
	createdAfterTime, err := time.Parse(time.RFC3339, createdAfterStr)
	if err != nil {
		// Handle error when the provided value is not a valid time
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid created_after")
	}

	createdBeforeStr := c.QueryParam("created_before")
	createdBeforeTime, err := time.Parse(time.RFC3339, createdBeforeStr)
	if err != nil {
		// Handle error when the provided value is not a valid time
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid created_before")
	}

	filter := payload.AbsenFilter{
		UserID:        uint(userID),
		MahasiswaID:   uint(mahasiswaID),
		JadwalID:      uint(jadwalID),
		Status:        c.QueryParam("status"),
		ID:            uint(absenID),
		CreatedAfter:  createdAfterTime,
		CreatedBefore: createdBeforeTime,
	}

	absen, count, err := u.absenService.GetPageAbsens(c.Request().Context(), int(pageInt), int(limitInt), &filter)
	if err != nil {
		if err == utils.ErrAbsenNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success getting absen",
		"data":    absen,
		"meta": echo.Map{
			"total": count,
			"page":  pageInt,
			"limit": limitInt,
		},
	})
}

func (d *AbsenController) DeleteAbsen(c echo.Context) error {
	claims := d.jwtService.GetClaims(&c)
	role := claims["role"].(string)
	if role != "admin" {
		return echo.NewHTTPError(http.StatusForbidden, utils.ErrDidntHavePermission.Error())
	}
	absenIDParam := c.Param("absen_id")
	absenID64, err := strconv.ParseUint(absenIDParam, 10, 0)
	absenID := uint(absenID64)
	err = d.absenService.DeleteAbsen(c.Request().Context(), absenID)
	if err != nil {
		switch err {
		case utils.ErrAbsenNotFound:
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success deleting absen",
	})
}
package controller

import (
	"net/http"

	"presensee_project/model/payload"
	"presensee_project/usecase"
	"presensee_project/utils"
	"presensee_project/utils/jwt_service"
	"strconv"

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
	case role == "admin":
		return c.JSON(http.StatusOK, echo.Map{
			"message": "success getting absen",
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
	pageInt, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrInvalidNumber.Error())
	}

	limit := c.QueryParam("limit")
	if limit == "" {
		limit = "20"
	}
	limitInt, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrInvalidNumber.Error())
	}

	absen, err := u.absenService.GetPageAbsens(c.Request().Context(), int(pageInt), int(limitInt))
	if err != nil {
		if err == utils.ErrAbsenNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success getting document",
		"data":    absen,
		"meta": echo.Map{
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

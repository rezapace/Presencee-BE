package controller

import (
	"net/http"

	"presensee_project/model"
	"presensee_project/usecase"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LoginUserController(c echo.Context) error {
	userRequest := model.User{}
	c.Bind(&userRequest)

	user, err := usecase.LoginUser(&userRequest)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userRequest.Password))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"user":   userRequest,
	})
}

package controller

import (
	"net/http"
	"presensee_project/middleware"
	"presensee_project/model"
	"presensee_project/model/payload"
	"presensee_project/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {
	users, err := usecase.GetListUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func GetUserController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	user, err := usecase.GetUser(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error get user",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	payload := payload.CreateUserRequest{}
	c.Bind(&payload)
	// validasi request body
	if err := c.Validate(payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error create user",
			"errorDescription": err,
		})
	}
	resp, err := usecase.CreateUser(&payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages":         "error create user",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"data":    resp,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	userID := middleware.ExtractTokenUserId(c)
	if userID == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete user",
			"errorDescription": "userId not found",
			"errorMessage":     "Sorry Token is invalid",
		})
	}
	if err := usecase.DeleteUser(uint(userID)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete user",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf user tidak dapat di hapus",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	user := model.User{}
	c.Bind(&user)
	user.ID = uint(id)
	if err := usecase.UpdateUser(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update user",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf user tidak dapat di ubah",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
	})
}
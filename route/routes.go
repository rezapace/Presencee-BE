package route

import (
	"net/http"
	"presensee_project/controller"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func NewRoute(e *echo.Echo, db *gorm.DB) {
	e.Validator = &CustomValidator{validator: validator.New()}

	e.POST("/login", controller.LoginUserController)
	e.POST("/register", controller.CreateUserController)

	// user collection
	user := e.Group("/users")
	user.GET("", controller.GetUserController)
	user.GET("/:id", controller.GetUserController)
	user.POST("", controller.CreateUserController)
	user.PUT("/:id", controller.UpdateUserController)
	user.DELETE("/:id", controller.DeleteUserController)

	// book collection
	mahasiswa := e.Group("/mahasiswa")
	mahasiswa.GET("", controller.GetMahasiswasController)
	mahasiswa.GET("/:id", controller.GetMahasiswaController)
	mahasiswa.POST("", controller.CreateMahasiswaController)
	mahasiswa.PUT("/:id", controller.UpdateMahasiswaController)
	mahasiswa.DELETE("/:id", controller.DeleteMahasiswaController)
	
	// book collection
	// book := e.Group("/books")
	// book.GET("", controller.GetBookController)
	// book.GET("/:id", controller.GetBookController)
	// book.POST("", controller.CreateBookController)
	// book.PUT("/:id", controller.UpdateBookController)
	// book.DELETE("/:id", controller.DeleteBookController)

	// // book collection
	// blog := e.Group("/blogs")
	// blog.GET("", controller.GetBlogController)
	// blog.GET("/:id", controller.GetBlogController)
	// blog.POST("", controller.CreateBlogController)
	// blog.PUT("/:id", controller.UpdateBlogController)
	// blog.DELETE("/:id", controller.DeleteBlogController)
}

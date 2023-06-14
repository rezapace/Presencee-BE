package routes

import (
	"presensee_project/controller"
	"presensee_project/utils/validation"

	ControllerPkg "presensee_project/controller"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Routes struct {
	userController  *ControllerPkg.UserController
	absenController *ControllerPkg.AbsenController
}

func NewRoutes(userController *ControllerPkg.UserController, absenController *ControllerPkg.AbsenController) *Routes {
	return &Routes{
		userController:  userController,
		absenController: absenController,
	}
}

func (r *Routes) Init(e *echo.Echo, conf map[string]string) {
	e.Pre(middleware.AddTrailingSlash())
	e.Use(middleware.Recover())

	e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.Validator = &validation.CustomValidator{Validator: validator.New()}

	jwtMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(conf["JWT_SECRET"]),
	})

	v1 := e.Group("/v1")

	// Users
	users := v1.Group("/users")
	users.POST("/signup/", r.userController.SignUpUser)
	users.POST("/login/", r.userController.LoginUser)

	usersWithAuth := users.Group("", jwtMiddleware)
	usersWithAuth.GET("/", r.userController.GetBriefUsers)
	usersWithAuth.PUT("/", r.userController.UpdateUser)
	usersWithAuth.GET("/:user_id/", r.userController.GetSingleUser, jwtMiddleware)
	usersWithAuth.DELETE("/:user_id/", r.userController.DeleteUser, jwtMiddleware)

	// mahasiswa collection
	mahasiswa := v1.Group("/mahasiswa", jwtMiddleware)
	mahasiswa.GET("/", controller.GetMahasiswasController)
	mahasiswa.GET("/:id/", controller.GetMahasiswaController)
	mahasiswa.POST("/", controller.CreateMahasiswaController)
	mahasiswa.PUT("/:id/", controller.UpdateMahasiswaController)
	mahasiswa.DELETE("/:id/", controller.DeleteMahasiswaController)

	// dosen
	dosen := v1.Group("/dosen", jwtMiddleware)
	dosen.GET("/", controller.GetDosensController)
	dosen.GET("/:id/", controller.GetDosenController)
	dosen.POST("/", controller.CreateDosenController)
	dosen.PUT("/:id/", controller.UpdateDosenController)
	dosen.DELETE("/:id/", controller.DeleteDosenController)

	// room
	room := v1.Group("/room", jwtMiddleware)
	room.GET("/", controller.GetRoomsController)
	room.GET("/:id/", controller.GetRoomController)
	room.POST("/", controller.CreateRoomController)
	room.PUT("/:id/", controller.UpdateRoomController)
	room.DELETE("/:id/", controller.DeleteRoomController)

	// matakuliah
	matakuliah := v1.Group("/matakuliah", jwtMiddleware)
	matakuliah.GET("/", controller.GetMatakuliahsController)
	matakuliah.GET("/:id/", controller.GetMatakuliahController)
	matakuliah.GET("/:name/", controller.GetMatakuliahByNameAndDateController)
	matakuliah.POST("/", controller.CreateMatakuliahController)
	matakuliah.PUT("/:id/", controller.UpdateMatakuliahController)
	matakuliah.DELETE("/:id/", controller.DeleteMatakuliahController)

	// absen
	absens := v1.Group("/absens", jwtMiddleware)
	absens.POST("/", r.absenController.CreateAbsen)
	absens.PUT("/", r.absenController.UpdateAbsen, jwtMiddleware)
	absens.GET("/:absen_id/", r.absenController.GetSingleAbsen, jwtMiddleware)
	absens.GET("/", r.absenController.GetPageAbsen)
	absens.GET("/riwayat/", r.absenController.GetRiwayat)
	absens.GET("/filter/", r.absenController.GetFilterAbsen)
	absens.DELETE("/:absen_id/", r.absenController.DeleteAbsen, jwtMiddleware)

	// Jurusan
	jurusan := v1.Group("/jurusan", jwtMiddleware)
	jurusan.GET("/", controller.GetJurusansController)
	jurusan.GET("/:id/", controller.GetJurusanController)
	jurusan.POST("/", controller.CreateJurusanController)
	jurusan.PUT("/:id/", controller.UpdateJurusanController)
	jurusan.DELETE("/:id/", controller.DeleteJurusanController)

	// jadwal
	jadwal := v1.Group("/jadwal", jwtMiddleware)
	jadwal.GET("/", controller.GetJadwalsController)
	jadwal.GET("/:id/", controller.GetJadwalController)
	jadwal.POST("/", controller.CreateJadwalController)
	jadwal.PUT("/:id/", controller.UpdateJadwalController)
	jadwal.DELETE("/:id/", controller.DeleteJadwalController)
	jadwal.GET("/filter/", controller.GetListJadwalsByDateController)

	//Upload File
	upload := v1.Group("/upload")
	upload.POST("/", controller.UploadFile)
}

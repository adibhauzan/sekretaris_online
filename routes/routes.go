package routes

import (
	"github.com/adibhauzan/sekretaris_online_backend/controllers"
	"github.com/adibhauzan/sekretaris_online_backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(jadwalController *controllers.JadwalController, statusController *controllers.StatusController, userController *controllers.UserController) *gin.Engine {
	r := gin.Default()

	jadwalRoutes := r.Group("/jadwal")
	{
		jadwalRoutes.Use(middleware.AuthMiddleware())

		jadwalRoutes.POST("/", jadwalController.CreateJadwal)
		jadwalRoutes.GET("/", jadwalController.GetAllJadwal)
		jadwalRoutes.GET("/:id", jadwalController.GetJadwalByID)
		jadwalRoutes.PUT("/:id", jadwalController.UpdateJadwal)
		jadwalRoutes.DELETE("/:id", jadwalController.DeleteJadwal)
		jadwalRoutes.GET("/bydatetime", jadwalController.GetJadwalByDatetime)
	}

	statusRoutes := r.Group("/status")
	{
		statusRoutes.POST("/", statusController.CreateStatus)
		statusRoutes.GET("/", statusController.GetAllStatus)
		statusRoutes.GET("/:id", statusController.GetStatusByID)
		statusRoutes.PUT("/:id", statusController.UpdateStatus)
		statusRoutes.DELETE("/:id", statusController.DeleteStatus)
	}

	userRoutes := r.Group("/user")
	{
		userRoutes.POST("/register", userController.CreateUser)
		userRoutes.POST("/login", userController.UserLogin)
		userRoutes.POST("/logout", userController.UserLogout)
	}

	return r
}

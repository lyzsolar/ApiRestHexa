package routes

import (
	"APIDOS/src/mascotas/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutesMascotas(
	r *gin.Engine,
	CreateMascotasController *controllers.CreateMascotasController,
	GetAllMascotasController *controllers.GetAllMascotasController,
	GetByIdMascotasController *controllers.GetByIdMascotasController,
	UpdateMascotasController *controllers.UpdateMascotasController,
	deleteMascotasController *controllers.DeleteMascotasController,
) {
	r.GET("/mascotas", GetAllMascotasController.Execute)
	r.GET("/mascotas/:id", GetByIdMascotasController.Execute)
	r.POST("/mascotas", CreateMascotasController.Execute)
	r.PUT("/mascotas/:id", UpdateMascotasController.Execute)
	r.DELETE("/mascotas/:id", deleteMascotasController.Execute)
}

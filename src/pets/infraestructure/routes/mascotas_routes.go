package routes

import (
	"APIDOS/src/pets/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutesMascotas(
	r *gin.Engine,
	CreateMascotasController *controllers.CreateMascotaController,
	GetAllMascotasController *controllers.GetAllMascotaController,
	UpdateMascotasController *controllers.UpdateMascotaController,
	deleteMascotasController *controllers.DeleteMascotaController,
) {
	r.GET("/Pets", GetAllMascotasController.Execute)
	r.POST("/Pets", CreateMascotasController.Execute)
	r.PUT("/Pets/:id", UpdateMascotasController.Execute)
	r.DELETE("/Pets/:id", deleteMascotasController.Execute)
}

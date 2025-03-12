package routes

import (
	"APIDOS/src/veterinarian/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutesMascotas(r *gin.Engine,
	createVeterinarianController *controllers.CreateVeterinarianController,
	getAllVeterinarianController *controllers.GetAllVeterinarianController,
	updateVeterinarianController *controllers.UpdateVeterinarianController,
	deleteVeterinarianController *controllers.DeleteVeterinarianController) {

	r.GET("/Veterinarians", getAllVeterinarianController.Execute)
	r.POST("/Veterinarians", createVeterinarianController.Execute)
	r.PUT("/Veterinarians/:id", updateVeterinarianController.Execute)
	r.DELETE("/Veterinarians/:id", deleteVeterinarianController.Execute)
}

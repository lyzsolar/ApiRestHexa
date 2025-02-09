package routes

import (
	"APIDOS/src/veterinario/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutesVeterinario(
	r *gin.Engine,
	CreateVeterinarioController *controllers.CreateVeterinarioController,
	GetAllVeterinarioController *controllers.GetAllVeterinarioController,
	GetByIdVeterinarioController *controllers.GetByIdVeterinarioController,
	UpdateVeterinarioController *controllers.UpdateVeterinarioController,
	deleteVeterinarioController *controllers.DeleteVeterinarioController,
) {
	r.GET("/Veterinario", GetAllVeterinarioController.Execute)
	r.GET("/Veterinario/:id", GetByIdVeterinarioController.Execute)
	r.POST("/Veterinario", CreateVeterinarioController.Execute)
	r.PUT("/Veterinario/:id", UpdateVeterinarioController.Execute)
	r.DELETE("/Veterinario/:id", deleteVeterinarioController.Execute)
}

package controllers

import (
	"APIDOS/src/pets/application"
	"github.com/gin-gonic/gin"
)

type GetAllMascotaController struct {
	useCase *application.GetAllMascotas
}

func NewGetAllMascotaController(useCase *application.GetAllMascotas) *GetAllMascotaController {
	return &GetAllMascotaController{useCase: useCase}
}

func (la_c *GetAllMascotaController) Execute(c *gin.Context) {
	mascota, err := la_c.useCase.Execute()
	if err != nil {
		c.JSON(500, gin.H{"error": "Error al obtener las pets"})
		return
	}
	c.JSON(200, mascota)
}

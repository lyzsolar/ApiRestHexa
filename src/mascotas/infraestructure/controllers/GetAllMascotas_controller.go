package controllers

import (
	"APIDOS/src/mascotas/application"
	"github.com/gin-gonic/gin"
)

type GetAllMascotasController struct {
	useCase *application.GetAllMascotas
}

func NewGetAllMascotasController(useCase *application.GetAllMascotas) *GetAllMascotasController {
	return &GetAllMascotasController{useCase: useCase}
}

func (la_c *GetAllMascotasController) Execute(c *gin.Context) {
	mascotas, err := la_c.useCase.Execute()
	if err != nil {
		c.JSON(500, gin.H{"error": "Error al obtener las mascotas"})
		return
	}
	c.JSON(200, mascotas)
}

package controllers

import (
	"APIDOS/src/veterinario/application"
	"github.com/gin-gonic/gin"
)

type GetAllVeterinarioController struct {
	useCase *application.GetAllVeterinario
}

func NewGetAllVeterinarioController(useCase *application.GetAllVeterinario) *GetAllVeterinarioController {
	return &GetAllVeterinarioController{useCase: useCase}
}

func (la_c *GetAllVeterinarioController) Execute(c *gin.Context) {
	veterinario, err := la_c.useCase.Execute()
	if err != nil {
		c.JSON(500, gin.H{"error": "Error al obtener los datos del veterinario"})
		return
	}
	c.JSON(200, veterinario)
}

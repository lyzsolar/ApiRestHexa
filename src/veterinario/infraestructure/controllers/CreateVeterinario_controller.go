package controllers

import (
	"APIDOS/src/veterinario/application"
	"APIDOS/src/veterinario/domain/entities"
	"github.com/gin-gonic/gin"
)

type CreateVeterinarioController struct {
	useCase *application.CreateVeterinario
}

func NewCreateVeterinarioController(useCase *application.CreateVeterinario) *CreateVeterinarioController {
	return &CreateVeterinarioController{useCase: useCase}
}

func (cs_a *CreateVeterinarioController) Execute(c *gin.Context) {
	var veterinario entities.Veterinario
	if err := c.ShouldBindJSON(&veterinario); err != nil {
		c.JSON(400, gin.H{"error": "Datos inválidos"})
		return
	}
	if err := cs_a.useCase.Execute(veterinario); err != nil {
		c.JSON(500, gin.H{"error": "No fue posible registrar los datos del medico"})
		return
	}
	c.JSON(201, gin.H{"message": "Medico guardado correctamente"})
}

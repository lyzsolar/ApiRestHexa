package controllers

import (
	"APIDOS/src/veterinarian/application"
	"APIDOS/src/veterinarian/domain/entities"
	"github.com/gin-gonic/gin"
)

type CreateVeterinarianController struct {
	useCase *application.CreateVeterinaria
}

func NewCreateVeterinariaController(useCase *application.CreateVeterinaria) *CreateVeterinarianController {
	return &CreateVeterinarianController{useCase: useCase}
}

func (cv_c *CreateVeterinarianController) Execute(c *gin.Context) {
	var veterinarian entities.Veterinarian
	if err := c.ShouldBindJSON(&veterinarian); err != nil {
		c.JSON(400, gin.H{"error": "Datos inválidos"})
		return
	}
	if err := cv_c.useCase.Execute(veterinarian); err != nil {
		c.JSON(500, gin.H{"error": "No se pudo crear el veterinario"})
		return
	}
	c.JSON(201, gin.H{"message": "Veterinario creado correctamente"})
}

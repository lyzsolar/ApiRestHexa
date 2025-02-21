package controllers

import (
	"APIDOS/src/veterinarian/application"
	"APIDOS/src/veterinarian/domain/entities"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UpdateVeterinarianController struct {
	useCase *application.UpdateVeterinarian
}

func NewUpdateVeterinarianController(useCase *application.UpdateVeterinarian) *UpdateVeterinarianController {
	return &UpdateVeterinarianController{useCase: useCase}
}

func (uv_c *UpdateVeterinarianController) Execute(c *gin.Context) {
	var veterinarian entities.Veterinarian
	if err := c.ShouldBindJSON(&veterinarian); err != nil {
		c.JSON(400, gin.H{"error": "Datos inválidos"})
		return
	}
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID no válido"})
		return
	}
	veterinarian.ID = idInt
	if err := uv_c.useCase.Execute(veterinarian); err != nil {
		c.JSON(500, gin.H{"error": "No se pudo actualizar el veterinario"})
		return
	}
	c.JSON(200, gin.H{"message": "Veterinario actualizado"})
}

package controllers

import (
	"APIDOS/src/veterinario/application"
	"APIDOS/src/veterinario/domain/entities"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UpdateVeterinarioController struct {
	useCase *application.UpdateVeterinario
}

func NewUpdateVeterinarioController(useCase *application.UpdateVeterinario) *UpdateVeterinarioController {
	return &UpdateVeterinarioController{useCase: useCase}
}

func (us_c *UpdateVeterinarioController) Execute(c *gin.Context) {
	var veterinario entities.Veterinario
	if err := c.ShouldBindJSON(&veterinario); err != nil {
		c.JSON(400, gin.H{"error": "Datos inválidos"})
		return
	}
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID no válido"})
		return
	}
	veterinario.Id = idInt
	if err := us_c.useCase.Execute(veterinario); err != nil {
		c.JSON(500, gin.H{"error": "No se pudo actualizar los datos del medico"})
		return
	}
	c.JSON(200, gin.H{"message": "Medico actualizado"})
}

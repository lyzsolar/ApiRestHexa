package controllers

import (
	"APIDOS/src/pets/application"
	"APIDOS/src/pets/domain/entities"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UpdateMascotaController struct {
	useCase *application.UpdateMascota
}

func NewUpdateMascotaController(useCase *application.UpdateMascota) *UpdateMascotaController {
	return &UpdateMascotaController{useCase: useCase}
}

func (us_c *UpdateMascotaController) Execute(c *gin.Context) {
	var mascota entities.Mascotas
	if err := c.ShouldBindJSON(&mascota); err != nil {
		c.JSON(400, gin.H{"error": "Datos inválidos"})
		return
	}
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID no válido"})
		return
	}
	mascota.ID = idInt
	if err := us_c.useCase.Execute(mascota); err != nil {
		c.JSON(500, gin.H{"error": "No se pudo actualizar la mascota"})
		return
	}
	c.JSON(200, gin.H{"message": "Mascota actualizado"})
}

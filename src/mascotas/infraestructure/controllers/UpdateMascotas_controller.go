package controllers

import (
	"APIDOS/src/mascotas/application"
	"APIDOS/src/mascotas/domain/entities"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UpdateMascotasController struct {
	useCase *application.UpdateMascotas
}

func NewUpdateMascotasController(useCase *application.UpdateMascotas) *UpdateMascotasController {
	return &UpdateMascotasController{useCase: useCase}
}

func (us_c *UpdateMascotasController) Execute(c *gin.Context) {
	var mascotas entities.Mascotas
	if err := c.ShouldBindJSON(&mascotas); err != nil {
		c.JSON(400, gin.H{"error": "Datos inválidos"})
		return
	}
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID no válido"})
		return
	}
	mascotas.Id = idInt
	if err := us_c.useCase.Execute(mascotas); err != nil {
		c.JSON(500, gin.H{"error": "No se pudo actualizar la mascota"})
		return
	}
	c.JSON(200, gin.H{"message": "Mascota actualizado"})
}

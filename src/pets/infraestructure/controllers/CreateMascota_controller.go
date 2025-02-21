package controllers

import (
	"APIDOS/src/pets/application"
	"APIDOS/src/pets/domain/entities"
	"github.com/gin-gonic/gin"
)

type CreateMascotaController struct {
	useCase *application.CreateMascota
}

func NewCreateMascotaController(useCase *application.CreateMascota) *CreateMascotaController {
	return &CreateMascotaController{useCase: useCase}
}

func (cs_m *CreateMascotaController) Execute(c *gin.Context) {
	var mascota entities.Mascotas
	if err := c.ShouldBindJSON(&mascota); err != nil {
		c.JSON(400, gin.H{"error": "Datos inválidos"})
		return
	}
	if err := cs_m.useCase.Execute(mascota); err != nil {
		c.JSON(500, gin.H{"message": "No se pudo crear la mascota"})
		return
	}
	c.JSON(201, gin.H{"message": "Mascota creada con exito"})
}

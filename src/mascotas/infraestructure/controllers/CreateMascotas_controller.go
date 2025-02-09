package controllers

import (
	"APIDOS/src/mascotas/application"
	"APIDOS/src/mascotas/domain/entities"
	"github.com/gin-gonic/gin"
)

type CreateMascotasController struct {
	useCase *application.CreateMascotas
}

func NewCreateMascotasController(useCase *application.CreateMascotas) *CreateMascotasController {
	return &CreateMascotasController{useCase: useCase}
}

func (cs_a *CreateMascotasController) Execute(c *gin.Context) {
	var mascotas entities.Mascotas
	if err := c.ShouldBindJSON(&mascotas); err != nil {
		c.JSON(400, gin.H{"error": "Datos inválidos"})
		return
	}
	if err := cs_a.useCase.Execute(mascotas); err != nil {
		c.JSON(500, gin.H{"error": "No fue posible registrar los datos de la mascota"})
		return
	}
	c.JSON(201, gin.H{"message": "Mascota guardada correctamente"})
}

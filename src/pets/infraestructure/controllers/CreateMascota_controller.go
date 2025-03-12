package controllers

import (
	"APIDOS/src/pets/application"
	"APIDOS/src/pets/domain"
	"APIDOS/src/pets/domain/entities"
	"github.com/gin-gonic/gin"
)

type CreateMascotaController struct {
	useCase *application.CreateMascota
	pets    domain.IMascotas
}

func NewCreateMascotaController(useCase *application.CreateMascota, pets domain.IMascotas) *CreateMascotaController {
	return &CreateMascotaController{useCase: useCase, pets: pets}
}

func (cs_m *CreateMascotaController) Execute(c *gin.Context) {
	var pets entities.Mascotas
	if err := c.ShouldBindJSON(&pets); err != nil {
		c.JSON(400, gin.H{"error": "Datos inválidos"})
		return
	}

	created, err := cs_m.useCase.Execute(pets)
	if err != nil {
		c.JSON(500, gin.H{"error": "No se pudo crear la mascota"})
		return
	}

	c.JSON(201, gin.H{"message": "Mascota creada correctamente", "pet": created})
}

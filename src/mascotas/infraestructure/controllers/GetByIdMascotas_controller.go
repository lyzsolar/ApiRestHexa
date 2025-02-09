package controllers

import (
	"APIDOS/src/mascotas/application"
	"github.com/gin-gonic/gin"
	"strconv"
)

type GetByIdMascotasController struct {
	useCase *application.GetByIdMascotas
}

func NewGetByIdMascotasController(useCase *application.GetByIdMascotas) *GetByIdMascotasController {
	return &GetByIdMascotasController{useCase: useCase}
}

func (laid_c *GetByIdMascotasController) Execute(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Id no válido"})
		return
	}

	mascotas, err := laid_c.useCase.Execute(idInt)
	if err != nil {
		c.JSON(404, gin.H{"error": "mascota no encontrada"})
		return
	}

	c.JSON(200, mascotas)
}

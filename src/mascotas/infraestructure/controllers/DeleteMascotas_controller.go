package controllers

import (
	"APIDOS/src/mascotas/application"
	"github.com/gin-gonic/gin"
	"strconv"
)

type DeleteMascotasController struct {
	useCase *application.DeleteMascotas
}

func NewDeleteMascotasController(useCase *application.DeleteMascotas) *DeleteMascotasController {
	return &DeleteMascotasController{useCase: useCase}
}

func (da_c *DeleteMascotasController) Execute(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID no válido"})
		return
	}
	if err := da_c.useCase.Execute(idInt); err != nil {
		c.JSON(500, gin.H{"error": "No fue posible eliminar la mascota"})
		return
	}
	c.JSON(200, gin.H{"message": "Mascota eliminada con exito"})
}

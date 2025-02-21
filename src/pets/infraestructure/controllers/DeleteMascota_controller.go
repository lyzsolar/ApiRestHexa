package controllers

import (
	"APIDOS/src/pets/application"
	"github.com/gin-gonic/gin"
	"strconv"
)

type DeleteMascotaController struct {
	useCase *application.DeleteMascota
}

func NewDeleteMascotaController(useCase *application.DeleteMascota) *DeleteMascotaController {
	return &DeleteMascotaController{useCase: useCase}
}

func (dm_c *DeleteMascotaController) Execute(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID no valida"})
		return
	}
	if err := dm_c.useCase.Execute(idInt); err != nil {
		c.JSON(500, gin.H{"message": "No se pudo eliminar la mascota"})
		return
	}
	c.JSON(200, gin.H{"message": "Mascota eliminada con exito"})
}

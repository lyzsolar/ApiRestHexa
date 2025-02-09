package controllers

import (
	"APIDOS/src/veterinario/application"
	"github.com/gin-gonic/gin"
	"strconv"
)

type DeleteVeterinarioController struct {
	useCase *application.DeleteVeterinario
}

func NewDeleteVeterinarioController(useCase *application.DeleteVeterinario) *DeleteVeterinarioController {
	return &DeleteVeterinarioController{useCase: useCase}
}

func (da_c *DeleteVeterinarioController) Execute(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID no válido"})
		return
	}
	if err := da_c.useCase.Execute(idInt); err != nil {
		c.JSON(500, gin.H{"error": "No fue posible eliminar al medico"})
		return
	}
	c.JSON(200, gin.H{"message": "Medico eliminado con exito"})
}

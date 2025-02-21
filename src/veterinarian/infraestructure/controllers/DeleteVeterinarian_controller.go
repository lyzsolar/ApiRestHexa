package controllers

import (
	"APIDOS/src/veterinarian/application"
	"github.com/gin-gonic/gin"
	"strconv"
)

type DeleteVeterinarianController struct {
	useCase *application.DeleteVeterinarian
}

func NewDeleteVeterinarianController(useCase *application.DeleteVeterinarian) *DeleteVeterinarianController {
	return &DeleteVeterinarianController{useCase: useCase}
}

func (dv_c *DeleteVeterinarianController) Execute(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID no válido"})
		return
	}
	if err := dv_c.useCase.Execute(idInt); err != nil {
		c.JSON(500, gin.H{"error": "No se pudo eliminar al veterinario"})
		return
	}
	c.JSON(200, gin.H{"message": "Veterinario eliminado"})
}

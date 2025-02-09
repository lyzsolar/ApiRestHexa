package controllers

import (
	"APIDOS/src/veterinario/application"
	"github.com/gin-gonic/gin"
	"strconv"
)

type GetByIdVeterinarioController struct {
	useCase *application.GetByIdVeterinario
}

func NewGetByIdVeterinarioController(useCase *application.GetByIdVeterinario) *GetByIdVeterinarioController {
	return &GetByIdVeterinarioController{useCase: useCase}
}

func (laid_c *GetByIdVeterinarioController) Execute(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Id no válido"})
		return
	}

	veterinario, err := laid_c.useCase.Execute(idInt)
	if err != nil {
		c.JSON(404, gin.H{"error": "medico no encontrado"})
		return
	}

	c.JSON(200, veterinario)
}

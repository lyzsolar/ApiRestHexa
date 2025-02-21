package controllers

import (
	"APIDOS/src/veterinarian/application"
	"github.com/gin-gonic/gin"
)

type GetAllVeterinarianController struct {
	useCase *application.GetAllVeterinarian
}

func NewGetAllVeterinarianController(useCase *application.GetAllVeterinarian) *GetAllVeterinarianController {
	return &GetAllVeterinarianController{useCase: useCase}
}

func (gv_c *GetAllVeterinarianController) Execute(c *gin.Context) {
	veterinarians, err := gv_c.useCase.Execute()
	if err != nil {
		c.JSON(500, gin.H{"error": "Error al obtener los veterinarios"})
		return
	}
	c.JSON(200, veterinarians)
}

package infraestructure

import (
	"APIDOS/src/veterinarian/application"
	"APIDOS/src/veterinarian/infraestructure/controllers"
	"APIDOS/src/veterinarian/infraestructure/database"
)

type DependenciesVeterinarios struct {
	CreateVeterinarianController *controllers.CreateVeterinarianController
	GetAllVeterinarianController *controllers.GetAllVeterinarianController
	UpdateVeterinarianController *controllers.UpdateVeterinarianController
	DeleteVeterinarianController *controllers.DeleteVeterinarianController
}

func InitVeterinariosDependencies() *DependenciesVeterinarios {
	ps := database.NewMySQL()

	return &DependenciesVeterinarios{
		CreateVeterinarianController: controllers.NewCreateVeterinariaController(application.NewCreateVeterinaria(ps)),
		GetAllVeterinarianController: controllers.NewGetAllVeterinarianController(application.NewGetAllVeterinarian(ps)),
		UpdateVeterinarianController: controllers.NewUpdateVeterinarianController(application.NewUpdateVeterinarian(ps)),
		DeleteVeterinarianController: controllers.NewDeleteVeterinarianController(application.NewDeleteVeterinarian(ps)),
	}
}

package infraestructure

import (
	"APIDOS/src/veterinario/application"
	"APIDOS/src/veterinario/infraestructure/controllers"
	"APIDOS/src/veterinario/infraestructure/database"
)

type DependenciesVeterinario struct {
	CreateVeterinarioController  *controllers.CreateVeterinarioController
	DeleteVeterinarioController  *controllers.DeleteVeterinarioController
	GetAllVeterinarioController  *controllers.GetAllVeterinarioController
	GetByIdVeterinarioController *controllers.GetByIdVeterinarioController
	UpdateVeterinarioController  *controllers.UpdateVeterinarioController
}

func InitVeterinario() *DependenciesVeterinario {
	ps := database.NewMySQL()

	return &DependenciesVeterinario{
		CreateVeterinarioController:  controllers.NewCreateVeterinarioController(application.NewCreateVeterinario(ps)),
		DeleteVeterinarioController:  controllers.NewDeleteVeterinarioController(application.NewDeleteVeterinario(ps)),
		GetAllVeterinarioController:  controllers.NewGetAllVeterinarioController(application.NewGetAllVeterinario(ps)),
		GetByIdVeterinarioController: controllers.NewGetByIdVeterinarioController(application.NewGetByIdVeterinario(ps)),
		UpdateVeterinarioController:  controllers.NewUpdateVeterinarioController(application.NewUpdateVeterinario(ps)),
	}
}

package infraestructure

import (
	"APIDOS/src/pets/application"
	"APIDOS/src/pets/infraestructure/controllers"
	"APIDOS/src/pets/infraestructure/database"
)

type DependenciesMascotas struct {
	CreateMascotaController *controllers.CreateMascotaController
	GetAllMascotaController *controllers.GetAllMascotaController
	UpdateMascotaController *controllers.UpdateMascotaController
	DeleteMascotaController *controllers.DeleteMascotaController
}

func InitMascotasDependencies() *DependenciesMascotas {
	ps := database.NewMySQL()

	return &DependenciesMascotas{
		CreateMascotaController: controllers.NewCreateMascotaController(application.NewCreateMascota(ps)),
		GetAllMascotaController: controllers.NewGetAllMascotaController(application.NewGetAllMascota(ps)),
		UpdateMascotaController: controllers.NewUpdateMascotaController(application.NewUpdateMascota(ps)),
		DeleteMascotaController: controllers.NewDeleteMascotaController(application.NewDeleteMascota(ps)),
	}
}

package infraestructure

import (
	"APIDOS/src/mascotas/application"
	"APIDOS/src/mascotas/infraestructure/controllers"
	"APIDOS/src/mascotas/infraestructure/database"
)

type DependenciesMascotas struct {
	CreateMascotasController  *controllers.CreateMascotasController
	DeleteMascotasController  *controllers.DeleteMascotasController
	GetAllMascotasController  *controllers.GetAllMascotasController
	GetByIdMascotasController *controllers.GetByIdMascotasController
	UpdateMascotasController  *controllers.UpdateMascotasController
}

func InitMascotas() *DependenciesMascotas {
	ps := database.NewMySQL()

	return &DependenciesMascotas{
		CreateMascotasController:  controllers.NewCreateMascotasController(application.NewCreateMascotas(ps)),
		DeleteMascotasController:  controllers.NewDeleteMascotasController(application.NewDeleteMascotas(ps)),
		GetAllMascotasController:  controllers.NewGetAllMascotasController(application.NewGetAllMascotas(ps)),
		GetByIdMascotasController: controllers.NewGetByIdMascotasController(application.NewGetByIdMascotas(ps)),
		UpdateMascotasController:  controllers.NewUpdateMascotasController(application.NewUpdateMascotas(ps)),
	}
}

package infraestructure

import (
	"APIDOS/src/pets/application"
	"APIDOS/src/pets/infraestructure/adapter"
	"APIDOS/src/pets/infraestructure/controllers"
	"APIDOS/src/pets/infraestructure/database"
	"log"
)

type DependenciesMascotas struct {
	CreateMascotaController *controllers.CreateMascotaController
	GetAllMascotaController *controllers.GetAllMascotaController
	UpdateMascotaController *controllers.UpdateMascotaController
	DeleteMascotaController *controllers.DeleteMascotaController
}

func InitMascotasDependencies() *DependenciesMascotas {
	ps, err := database.NewMySQL()

	rmqClient, err := adapter.NewRabbitMQAdapter()
	if err != nil {
		log.Fatalf("Error initializing RabbitMQ: %v", err)
	}

	return &DependenciesMascotas{
		CreateMascotaController: controllers.NewCreateMascotaController(application.NewCreatePets(ps, rmqClient), ps),
		GetAllMascotaController: controllers.NewGetAllMascotaController(application.NewGetAllMascota(ps)),
		UpdateMascotaController: controllers.NewUpdateMascotaController(application.NewUpdateMascota(ps)),
		DeleteMascotaController: controllers.NewDeleteMascotaController(application.NewDeleteMascota(ps)),
	}
}

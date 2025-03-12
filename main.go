package main

import (
	petsInfra "APIDOS/src/pets/infraestructure"
	petsRoutes "APIDOS/src/pets/infraestructure/routes"
	veterinarianInfra "APIDOS/src/veterinarian/infraestructure"
	veterinarianRoutes "APIDOS/src/veterinarian/infraestructure/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//configurar cors

	petsDependencies := petsInfra.InitMascotasDependencies()
	petsRoutes.ConfigureRoutesMascotas(r,
		petsDependencies.CreateMascotaController,
		petsDependencies.GetAllMascotaController,
		petsDependencies.UpdateMascotaController,
		petsDependencies.DeleteMascotaController,
	)

	veterinarianDependencies := veterinarianInfra.InitVeterinariosDependencies()
	veterinarianRoutes.ConfigureRoutesMascotas(r,
		veterinarianDependencies.CreateVeterinarianController,
		veterinarianDependencies.GetAllVeterinarianController,
		veterinarianDependencies.UpdateVeterinarianController,
		veterinarianDependencies.DeleteVeterinarianController,
	)

	r.Run(":8080")
}

package main

import (
	mascotasInfraes "APIDOS/src/mascotas/infraestructure"
	mascotasRoutes "APIDOS/src/mascotas/infraestructure/routes"
	veterinarioInfraes "APIDOS/src/veterinario/infraestructure"
	veterinarioRoutes "APIDOS/src/veterinario/infraestructure/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	mascotasDependencies := mascotasInfraes.InitMascotas()
	mascotasRoutes.ConfigureRoutesMascotas(
		r,
		mascotasDependencies.CreateMascotasController,
		mascotasDependencies.GetAllMascotasController,
		mascotasDependencies.GetByIdMascotasController,
		mascotasDependencies.UpdateMascotasController,
		mascotasDependencies.DeleteMascotasController,
	)

	veterinarioDependencies := veterinarioInfraes.InitVeterinario()
	veterinarioRoutes.ConfigureRoutesVeterinario(
		r,
		veterinarioDependencies.CreateVeterinarioController,
		veterinarioDependencies.GetAllVeterinarioController,
		veterinarioDependencies.GetByIdVeterinarioController,
		veterinarioDependencies.UpdateVeterinarioController,
		veterinarioDependencies.DeleteVeterinarioController,
	)

	r.Run(":8080")
}

package application

import (
	"APIDOS/src/veterinario/domain"
	"APIDOS/src/veterinario/domain/entities"
)

type GetAllVeterinario struct {
	db domain.IVeterinario
}

func NewGetAllVeterinario(db domain.IVeterinario) *GetAllVeterinario {
	return &GetAllVeterinario{db: db}
}

func (la *GetAllVeterinario) Execute() ([]entities.Veterinario, error) {
	return la.db.GetAll()
}

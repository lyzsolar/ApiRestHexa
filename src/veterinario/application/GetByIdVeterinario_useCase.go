package application

import (
	"APIDOS/src/veterinario/domain"
	"APIDOS/src/veterinario/domain/entities"
)

type GetByIdVeterinario struct {
	db domain.IVeterinario
}

func NewGetByIdVeterinario(db domain.IVeterinario) *GetByIdVeterinario {
	return &GetByIdVeterinario{db: db}
}

func (la *GetByIdVeterinario) Execute(id int) (entities.Veterinario, error) {
	return la.db.GetById(id)
}

package application

import (
	"APIDOS/src/veterinario/domain"
	"APIDOS/src/veterinario/domain/entities"
)

type CreateVeterinario struct {
	db domain.IVeterinario
}

func NewCreateVeterinario(db domain.IVeterinario) *CreateVeterinario {
	return &CreateVeterinario{db: db}
}

func (ca *CreateVeterinario) Execute(mascotas entities.Veterinario) error {
	return ca.db.Save(mascotas)
}

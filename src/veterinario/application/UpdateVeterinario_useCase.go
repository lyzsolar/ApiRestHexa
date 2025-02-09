package application

import (
	"APIDOS/src/veterinario/domain"
	"APIDOS/src/veterinario/domain/entities"
)

type UpdateVeterinario struct {
	db domain.IVeterinario
}

func NewUpdateVeterinario(db domain.IVeterinario) *UpdateVeterinario {
	return &UpdateVeterinario{db: db}
}

func (ua *UpdateVeterinario) Execute(veterinario entities.Veterinario) error {
	return ua.db.Update(veterinario)
}

package application

import (
	"APIDOS/src/veterinarian/domain"
	"APIDOS/src/veterinarian/domain/entities"
)

type UpdateVeterinarian struct {
	db domain.IVeterinarian
}

func NewUpdateVeterinarian(db domain.IVeterinarian) *UpdateVeterinarian {
	return &UpdateVeterinarian{db: db}
}

func (uv *UpdateVeterinarian) Execute(veterinarian entities.Veterinarian) error {
	return uv.db.Update(veterinarian)
}

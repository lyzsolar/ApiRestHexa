package application

import (
	"APIDOS/src/veterinarian/domain"
	"APIDOS/src/veterinarian/domain/entities"
)

type CreateVeterinaria struct {
	db domain.IVeterinarian
}

func NewCreateVeterinaria(db domain.IVeterinarian) *CreateVeterinaria {
	return &CreateVeterinaria{db: db}
}

func (cp *CreateVeterinaria) Execute(veterinarian entities.Veterinarian) error {
	return cp.db.Save(veterinarian)
}

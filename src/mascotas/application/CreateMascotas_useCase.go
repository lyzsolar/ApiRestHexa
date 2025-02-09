package application

import (
	"APIDOS/src/mascotas/domain"
	"APIDOS/src/mascotas/domain/entities"
)

type CreateMascotas struct {
	db domain.IMascotas
}

func NewCreateMascotas(db domain.IMascotas) *CreateMascotas {
	return &CreateMascotas{db: db}
}

func (ca *CreateMascotas) Execute(mascotas entities.Mascotas) error {
	return ca.db.Save(mascotas)
}

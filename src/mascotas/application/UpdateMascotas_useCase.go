package application

import (
	"APIDOS/src/mascotas/domain"
	"APIDOS/src/mascotas/domain/entities"
)

type UpdateMascotas struct {
	db domain.IMascotas
}

func NewUpdateMascotas(db domain.IMascotas) *UpdateMascotas {
	return &UpdateMascotas{db: db}
}

func (ua *UpdateMascotas) Execute(mascotas entities.Mascotas) error {
	return ua.db.Update(mascotas)
}

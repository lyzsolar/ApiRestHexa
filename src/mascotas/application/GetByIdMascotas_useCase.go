package application

import (
	"APIDOS/src/mascotas/domain"
	"APIDOS/src/mascotas/domain/entities"
)

type GetByIdMascotas struct {
	db domain.IMascotas
}

func NewGetByIdMascotas(db domain.IMascotas) *GetByIdMascotas {
	return &GetByIdMascotas{db: db}
}

func (la *GetByIdMascotas) Execute(id int) (entities.Mascotas, error) {
	return la.db.GetById(id)
}

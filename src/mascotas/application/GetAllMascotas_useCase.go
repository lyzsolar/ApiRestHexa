package application

import (
	"APIDOS/src/mascotas/domain"
	"APIDOS/src/mascotas/domain/entities"
)

type GetAllMascotas struct {
	db domain.IMascotas
}

func NewGetAllMascotas(db domain.IMascotas) *GetAllMascotas {
	return &GetAllMascotas{db: db}
}

func (la *GetAllMascotas) Execute() ([]entities.Mascotas, error) {
	return la.db.GetAll()
}

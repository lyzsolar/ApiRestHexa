package application

import (
	"APIDOS/src/pets/domain"
	"APIDOS/src/pets/domain/entities"
)

type GetAllMascotas struct {
	db domain.IMascotas
}

func NewGetAllMascota(db domain.IMascotas) *GetAllMascotas {
	return &GetAllMascotas{db: db}
}

func (gm *GetAllMascotas) Execute() ([]entities.Mascotas, error) {
	return gm.db.GetAll()
}

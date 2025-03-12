package application

import (
	"APIDOS/src/pets/domain"
)

type DeleteMascota struct {
	db domain.IMascotas
}

func NewDeleteMascota(db domain.IMascotas) *DeleteMascota {
	return &DeleteMascota{db: db}
}

func (dm *DeleteMascota) Execute(id int) error {
	return dm.db.Delete(id)
}

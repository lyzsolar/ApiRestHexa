package application

import (
	"APIDOS/src/mascotas/domain"
)

type DeleteMascotas struct {
	db domain.IMascotas
}

func NewDeleteMascotas(db domain.IMascotas) *DeleteMascotas {
	return &DeleteMascotas{db: db}
}

func (da *DeleteMascotas) Execute(id int) error {
	return da.db.Delete(id)
}

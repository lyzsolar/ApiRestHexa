package application

import (
	"APIDOS/src/veterinario/domain"
)

type DeleteVeterinario struct {
	db domain.IVeterinario
}

func NewDeleteVeterinario(db domain.IVeterinario) *DeleteVeterinario {
	return &DeleteVeterinario{db: db}
}

func (da *DeleteVeterinario) Execute(id int) error {
	return da.db.Delete(id)
}

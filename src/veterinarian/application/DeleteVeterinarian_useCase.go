package application

import "APIDOS/src/veterinarian/domain"

type DeleteVeterinarian struct {
	db domain.IVeterinarian
}

func NewDeleteVeterinarian(db domain.IVeterinarian) *DeleteVeterinarian {
	return &DeleteVeterinarian{db: db}
}

func (dv *DeleteVeterinarian) Execute(id int) error {
	return dv.db.Delete(id)
}

package application

import (
	"APIDOS/src/veterinarian/domain"
	"APIDOS/src/veterinarian/domain/entities"
)

type GetAllVeterinarian struct {
	db domain.IVeterinarian
}

func NewGetAllVeterinarian(db domain.IVeterinarian) *GetAllVeterinarian {
	return &GetAllVeterinarian{db: db}
}

func (gv *GetAllVeterinarian) Execute() ([]entities.Veterinarian, error) {
	return gv.db.GetAll()
}

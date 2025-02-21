package application

import (
	"APIDOS/src/pets/domain"
	"APIDOS/src/pets/domain/entities"
)

type CreateMascota struct {
	db domain.IMascotas
}

func NewCreateMascota(db domain.IMascotas) *CreateMascota {
	return &CreateMascota{db: db}
}

func (cm *CreateMascota) Execute(mascota entities.Mascotas) error {
	return cm.db.Save(mascota)
}

package application

import (
	"APIDOS/src/pets/domain"
	"APIDOS/src/pets/domain/entities"
)

type UpdateMascota struct {
	db domain.IMascotas
}

func NewUpdateMascota(db domain.IMascotas) *UpdateMascota {
	return &UpdateMascota{db: db}
}

func (um *UpdateMascota) Execute(mascota entities.Mascotas) error {
	return um.db.Update(mascota)
}

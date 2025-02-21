package domain

import "APIDOS/src/pets/domain/entities"

type IMascotas interface {
	Save(mascotas entities.Mascotas) error
	GetAll() ([]entities.Mascotas, error)
	Update(mascotas entities.Mascotas) error
	Delete(id int) error
}

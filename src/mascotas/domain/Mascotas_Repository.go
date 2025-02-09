package domain

import "APIDOS/src/mascotas/domain/entities"

type IMascotas interface {
	Save(mascotas entities.Mascotas) error
	GetAll() ([]entities.Mascotas, error)
	GetById(id int) (entities.Mascotas, error)
	Update(mascotas entities.Mascotas) error
	Delete(id int) error
}

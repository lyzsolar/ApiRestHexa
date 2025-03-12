package domain

import "APIDOS/src/veterinarian/domain/entities"

type IVeterinarian interface {
	Save(veterinario entities.Veterinarian) error
	GetAll() ([]entities.Veterinarian, error)
	Update(veterinario entities.Veterinarian) error
	Delete(id int) error
}

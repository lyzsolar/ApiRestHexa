package domain

import "APIDOS/src/veterinario/domain/entities"

type IVeterinario interface {
	Save(veterinario entities.Veterinario) error
	GetAll() ([]entities.Veterinario, error)
	GetById(id int) (entities.Veterinario, error)
	Update(veterinario entities.Veterinario) error
	Delete(id int) error
}

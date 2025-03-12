package application

import (
	"APIDOS/src/pets/application/repository"
	"APIDOS/src/pets/domain"
	"APIDOS/src/pets/domain/entities"
	"log"
)

type CreateMascota struct {
	petsRepo            domain.IMascotas
	serviceNotification repository.IMessageService
}

func NewCreatePets(petsRepo domain.IMascotas, serviceNotification repository.IMessageService) *CreateMascota {
	return &CreateMascota{
		petsRepo:            petsRepo,
		serviceNotification: serviceNotification,
	}
}

func (cm *CreateMascota) Execute(mascota entities.Mascotas) (entities.Mascotas, error) {
	created, err := cm.petsRepo.Save(mascota)
	if err != nil {
		return entities.Mascotas{}, err
	}

	err = cm.serviceNotification.PublishEvent("PetCreated", created)
	if err != nil {
		log.Printf("Error notifying about created pet: %v", err)
		return entities.Mascotas{}, err
	}

	return created, nil
}

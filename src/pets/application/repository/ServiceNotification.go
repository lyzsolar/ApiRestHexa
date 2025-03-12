package repository

import (
	"APIDOS/src/pets/domain/entities"
	"log"
)

type ServiceNotification struct {
	imageService IMessageService
}

func NewServiceNotification(imageService IMessageService) *ServiceNotification {
	return &ServiceNotification{imageService: imageService}
}

func (sn *ServiceNotification) NotifypetCreated(pets entities.Mascotas) error {
	log.Println("Notificando la creacion de las mascotas ...")

	err := sn.imageService.PublishEvent("AppointmentCreated", pets)
	if err != nil {
		log.Println("Error al publicar el evento %v", err)
		return err
	}
	return nil
}

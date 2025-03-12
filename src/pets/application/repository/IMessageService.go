package repository

import "APIDOS/src/pets/domain/entities"

type IMessageService interface {
	PublishEvent(eventType string, pets entities.Mascotas) error
}

package interfaces

import (
	"fcm_sender/models"
)

type ISenderService interface {
	SendNotification(sender *models.Sender) (map[string]interface{}, error)
}

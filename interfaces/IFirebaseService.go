package interfaces

import (
	"context"
	"fcm_sender/models"
	"firebase.google.com/go/messaging"
)

type IFirebaseService interface {
	GetClient() *messaging.Client
	GetContext() context.Context
	GenerateMessage(sender *models.Sender) *messaging.MulticastMessage
}

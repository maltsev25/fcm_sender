package services

import (
	"context"
	"fcm_sender/environment"
	"fcm_sender/models"
	"fcm_sender/utils"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
	"google.golang.org/api/option"
	"log"
	"time"
)

type FirebaseService struct {
	client *messaging.Client
	ctx    context.Context
}

var env = environment.Get()
var _firebaseService *FirebaseService

func init() {
	ctx := context.Background()
	opt := option.WithCredentialsFile("./serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	_firebaseService = &FirebaseService{
		client: client,
		ctx:    ctx,
	}
}

func (firebaseService FirebaseService) GetClient() *messaging.Client {
	return _firebaseService.client
}

func (firebaseService FirebaseService) GetContext() context.Context {
	return _firebaseService.ctx
}

func (firebaseService FirebaseService) GenerateMessage(sender *models.Sender) *messaging.MulticastMessage {
	var data = make(map[string]string)
	data[env.DataPrefix+".recipientId"] = fmt.Sprintf("%s", sender.RecipientId)
	data[env.DataPrefix+".notificationId"] = fmt.Sprintf("%s", sender.NotificationId)
	for key, val := range sender.Data {
		data[env.DataPrefix+"."+key] = val
	}

	ttl := time.Duration(env.NotificationTTL) * time.Minute

	var androidNotification *messaging.AndroidNotification = nil
	var apnsConfig *messaging.APNSConfig = nil
	addingImage(sender, androidNotification, apnsConfig)

	message := &messaging.MulticastMessage{
		Tokens: sender.Destination,
		Notification: &messaging.Notification{
			Title: "ID №" + utils.FormatToBill(fmt.Sprintf("%s", sender.RecipientId)),
			Body:  sender.Message,
		},
		Android: &messaging.AndroidConfig{
			TTL:          &ttl,
			Notification: androidNotification,
		},
		APNS: apnsConfig,
		Data: data,
	}

	return message
}

// в случае получения ссылки на картинку, добавляем её в пуш
func addingImage(sender *models.Sender,
	androidNotification *messaging.AndroidNotification,
	apnsConfig *messaging.APNSConfig) {
	if sender.ImageUrl != "" {
		androidNotification = &messaging.AndroidNotification{
			ImageURL: sender.ImageUrl,
		}
		apnsConfig = &messaging.APNSConfig{
			Payload: &messaging.APNSPayload{
				Aps: &messaging.Aps{
					MutableContent: true,
				},
			},
			FCMOptions: &messaging.APNSFCMOptions{
				ImageURL: sender.ImageUrl,
			},
		}
	}
}

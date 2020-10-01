package servicecontainer

import (
	"fcm_sender/controllers"
	"fcm_sender/services"
	"sync"
)

type IServiceContainer interface {
	InjectSenderController() controllers.SenderController
}

type kernel struct{}

func (k *kernel) InjectSenderController() controllers.SenderController {

	firebaseService := &services.FirebaseService{}
	senderService := &services.SenderService{Firebase: firebaseService}
	senderController := controllers.SenderController{Service: senderService}

	return senderController
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}

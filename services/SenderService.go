package services

import (
	"fcm_sender/interfaces"
	"fcm_sender/models"
	"github.com/stretchr/stew/slice"
)

type SenderService struct {
	Firebase interfaces.IFirebaseService
}

func (service SenderService) SendNotification(sender *models.Sender) (map[string]interface{}, error) {
	message := service.Firebase.GenerateMessage(sender)
	result, err := service.Firebase.GetClient().SendMulticast(service.Firebase.GetContext(), message)
	if err != nil {
		return nil, err
	}

	var failedTokens = make([]string, 0)
	if result.FailureCount > 0 {
		for idx, resp := range result.Responses {
			if !resp.Success {
				failedTokens = append(failedTokens, sender.Destination[idx])
			}
		}
	}
	response := make(map[string]interface{})
	response["status"] = false
	if checkSendingSuccess(sender.Destination, failedTokens) {
		response["status"] = true
	}
	response["failed"] = failedTokens

	return response, nil
}

// проверка успешности отправки пуша
func checkSendingSuccess(tokens, failedTokens []string) bool {
	var successTokens []string
	for _, token := range tokens {
		if !slice.ContainsString(failedTokens, token) {
			successTokens = append(successTokens, token)
		}
	}
	return len(successTokens) > 0
}

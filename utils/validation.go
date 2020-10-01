package utils

import (
	"fcm_sender/models"
	"fmt"
	"github.com/go-playground/validator/v10"
)

// генерация сообшения об ошибке
func generateValidationMessage(field string, rule string, param string) (message string) {
	switch rule {
	case "required":
		return fmt.Sprintf("Field '%s' is '%s'.", field, rule)
	default:
		return fmt.Sprintf("Field '%s' is not '%s %s'.", field, rule, param)
	}
}

// ответ при ошибке валидации
func GenerateValidationResponse(err error) (response models.ValidationResponse) {
	response.Status = false

	var validations []models.Validation

	validationErrors := err.(validator.ValidationErrors)

	for _, value := range validationErrors {
		field, rule, param := value.Field(), value.Tag(), value.Param()
		validation := models.Validation{Field: field, Message: generateValidationMessage(field, rule, param)}
		validations = append(validations, validation)
	}

	response.Validations = validations

	return response
}

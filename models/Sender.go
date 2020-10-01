package models

import "encoding/json"

type Data map[string]string

// структура получаемого сообщения для отправки
type Sender struct {
	RecipientId    json.Number `json:"recipient_id" binding:"required"`
	Message        string      `json:"message" binding:"required"`
	ImageUrl       string      `json:"image_url"`
	Destination    []string    `json:"destination" binding:"required"`
	NotificationId json.Number `json:"notification_id" binding:"required"`
	Data           Data        `json:"data"`
}

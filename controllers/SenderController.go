package controllers

import (
	"fcm_sender/environment"
	"fcm_sender/interfaces"
	"fcm_sender/models"
	"fcm_sender/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var env = environment.Get()

type SenderController struct {
	Service interfaces.ISenderService
}

// отправка уведомления в fcm
func (controller SenderController) SendNotification(c *gin.Context) {
	sender := &models.Sender{}

	err := c.ShouldBindJSON(sender)
	if err != nil {
		response := utils.GenerateValidationResponse(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if !env.ReleaseMode {
		log.Println(sender)
	}

	response, err := controller.Service.SendNotification(sender)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, err.Error())
		return
	}

	c.JSON(http.StatusOK, response)
}

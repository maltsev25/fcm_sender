package controllers

import (
	"fcm_sender/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HeartBeat(c *gin.Context) {
	data := utils.Message(true, "Alive")
	c.JSON(http.StatusOK, data)
}

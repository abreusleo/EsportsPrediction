package controllers

import (
	"Api/internal/player"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PlayerController struct {
	service *player.Service
}

func NewPlayerController(service *player.Service) *PlayerController {
	return &PlayerController{service: service}
}

func (ctrl *PlayerController) GetPossiblePlayers(c *gin.Context) {
	result := ctrl.service.GetPossiblePlayers()

	c.JSON(http.StatusOK, result)
}

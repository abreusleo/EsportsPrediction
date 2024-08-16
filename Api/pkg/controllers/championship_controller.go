package controllers

import (
	"Api/internal/championship"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChampionshipController struct {
	service *championship.Service
}

func NewChampionshipController(service *championship.Service) *ChampionshipController {
	return &ChampionshipController{service: service}
}

func (ctrl *ChampionshipController) GetPossibleChampionships(c *gin.Context) {
	result := ctrl.service.GetPossibleChampionships()

	c.JSON(http.StatusOK, result)
}

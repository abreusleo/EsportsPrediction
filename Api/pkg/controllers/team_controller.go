package controllers

import (
	"Api/internal/team"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TeamController struct {
	service *team.Service
}

func NewTeamController(service *team.Service) *TeamController {
	return &TeamController{service: service}
}

func (ctrl *TeamController) GetPossibleTeams(c *gin.Context) {
	result := ctrl.service.GetPossibleTeams()

	c.JSON(http.StatusOK, result)
}

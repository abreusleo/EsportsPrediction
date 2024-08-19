package controllers

import (
	"Api/internal/prediction"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PredictionController struct {
	service *prediction.Service
}

func NewPredictionController(service *prediction.Service) *PredictionController {
	return &PredictionController{service: service}
}

func (ctrl *PredictionController) GetTypes(c *gin.Context) {
	types := prediction.GetPredictionTypeNames()

	c.JSON(http.StatusOK, types)
}

func (ctrl *PredictionController) Predict(c *gin.Context) {
	predictionType := c.Param("type")
	predictTypeInt, _ := strconv.Atoi(predictionType)

	c.JSON(http.StatusOK, ctrl.service.Predict(predictTypeInt))
}

package main

import (
	"Api/internal/championship"
	"Api/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	ChampionshipController := initChampionshipController()

	r.GET("/championship", ChampionshipController.GetPossibleChampionships)
	r.Run()
}

func initChampionshipController() controllers.ChampionshipController {
	championshipService := championship.NewService()

	championshipController := controllers.NewChampionshipController(championshipService)

	return *championshipController
}

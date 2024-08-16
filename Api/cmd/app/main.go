package main

import (
	"Api/internal/championship"
	"Api/internal/player"
	"Api/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	ChampionshipController := initChampionshipController()
	PlayerController := initPlayerController()

	r.GET("/championship", ChampionshipController.GetPossibleChampionships)
	r.GET("/player", PlayerController.GetPossiblePlayers)
	r.Run()
}

func initChampionshipController() controllers.ChampionshipController {
	championshipService := championship.NewService()

	championshipController := controllers.NewChampionshipController(championshipService)

	return *championshipController
}

func initPlayerController() controllers.PlayerController {
	playerService := player.NewService()

	playerController := controllers.NewPlayerController(playerService)

	return *playerController
}

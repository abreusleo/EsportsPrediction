package main

import (
	"Api/internal/championship"
	"Api/internal/player"
	"Api/internal/team"
	"Api/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	ChampionshipController := initChampionshipController()
	PlayerController := initPlayerController()
	TeamController := initTeamController()

	r.GET("/championship", ChampionshipController.GetPossibleChampionships)
	r.GET("/player", PlayerController.GetPossiblePlayers)
	r.GET("/team", TeamController.GetPossibleTeams)
	r.Run()
}

// TODO Look for another way to initialize all constructors. Pretty sure this a horrible way to do this
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

func initTeamController() controllers.TeamController {
	teamService := team.NewService()

	teamController := controllers.NewTeamController(teamService)

	return *teamController
}

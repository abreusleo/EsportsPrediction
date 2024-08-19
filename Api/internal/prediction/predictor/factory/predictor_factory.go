package factory

import (
	"Api/internal/prediction/predictor"
	"Api/internal/prediction/predictor/implementations"
)

func GetPredictor(predictorType string) predictor.IPredictor {
	if predictorType == "teams" {
		return implementations.NewTeamPredictor()
	}
	if predictorType == "championships" {
		return implementations.NewChampionshipPredictor()
	}
	if predictorType == "players" {
		return implementations.NewPlayerPredictor()
	}
	return nil
}

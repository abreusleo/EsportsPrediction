package predictor

func GetPredictor(predictorType string) IPredictor {
	if predictorType == "teams" {
		return NewTeamPredictor()
	}
	if predictorType == "championships" {
		return NewChampionshipPredictor()
	}
	if predictorType == "players" {
		return NewPlayerPredictor()
	}
	return nil
}

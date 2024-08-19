package implementations

import "Api/internal/prediction/predictor"

type ChampionshipPredictor struct {
	predictor.Predictor
}

func NewChampionshipPredictor() predictor.IPredictor {
	return &ChampionshipPredictor{
		Predictor: predictor.Predictor{},
	}
}
func (cp *ChampionshipPredictor) Predict() string {
	return "Championship!"
}

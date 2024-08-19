package implementations

import "Api/internal/prediction/predictor"

type TeamPredictor struct {
	predictor.Predictor
}

func NewTeamPredictor() predictor.IPredictor {
	return &TeamPredictor{
		Predictor: predictor.Predictor{},
	}
}

func (cp *TeamPredictor) Predict() string {
	return "Team!"
}

package implementations

import "Api/internal/prediction/predictor"

type PlayerPredictor struct {
	predictor.Predictor
}

func NewPlayerPredictor() predictor.IPredictor {
	return &PlayerPredictor{
		Predictor: predictor.Predictor{},
	}
}

func (cp *PlayerPredictor) Predict() string {
	return "Player!"
}

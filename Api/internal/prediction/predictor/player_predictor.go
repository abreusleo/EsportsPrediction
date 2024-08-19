package predictor

type PlayerPredictor struct {
	Predictor
}

func NewPlayerPredictor() IPredictor {
	return &PlayerPredictor{
		Predictor: Predictor{},
	}
}

func (cp *PlayerPredictor) Predict() string {
	return "Player!"
}

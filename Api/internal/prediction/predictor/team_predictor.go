package predictor

type TeamPredictor struct {
	Predictor
}

func NewTeamPredictor() IPredictor {
	return &TeamPredictor{
		Predictor: Predictor{},
	}
}

func (cp *TeamPredictor) Predict() string {
	return "Team!"
}

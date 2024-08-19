package predictor

type ChampionshipPredictor struct {
	Predictor
}

func NewChampionshipPredictor() IPredictor {
	return &ChampionshipPredictor{
		Predictor: Predictor{},
	}
}
func (cp *ChampionshipPredictor) Predict() string {
	return "Championship!"
}

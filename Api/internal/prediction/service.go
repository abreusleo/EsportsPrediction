package prediction

import (
	"Api/internal/prediction/predictor"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Predict(predictorType int) string {
	predictionTypeString := PredictionTypes.String(PredictionTypes(predictorType))
	predict := predictor.GetPredictor(predictionTypeString)

	return predict.Predict()
}

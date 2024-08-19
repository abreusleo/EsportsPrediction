package prediction

import (
	"Api/internal/prediction/predictor/factory"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Predict(predictorType int) string {
	predictionTypeString := PredictionTypes.String(PredictionTypes(predictorType))
	predict := factory.GetPredictor(predictionTypeString)

	return predict.Predict()
}

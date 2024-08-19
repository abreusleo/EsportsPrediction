package prediction

import (
	"Api/internal/prediction/predictor"
	"fmt"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Predict(predictorType int) string {
	predictionTypeString := PredictionTypes.String(PredictionTypes(predictorType))
	fmt.Printf(predictionTypeString)
	predict := predictor.GetPredictor(predictionTypeString)

	return predict.Predict()
}

package prediction

import "encoding/json"

type PredictionDto struct {
	Type  int             `json:"id"`
	Input json.RawMessage `json:"input"`
}

type PredictionTypes int

const (
	Players PredictionTypes = iota
	Teams
	Championships
)

func (s PredictionTypes) String() string {
	switch s {
	case Players:
		return "players"
	case Teams:
		return "teams"
	case Championships:
		return "championships"
	}
	return "unknown"
}

func GetPredictionTypeNames() []string {
	return []string{
		Players.String(),
		Teams.String(),
		Championships.String(),
	}
}

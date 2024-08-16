package player

type GetPossiblePlayersDto struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	PhotoUrl string `json:"photo_url"`
}

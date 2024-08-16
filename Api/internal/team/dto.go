package team

type GetPossibleTeamsDto struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Acronym string `json:"acronym"`
	LogoUrl string `json:"logo_url"`
}

package championship

type GetPossibleChampionshipsDto struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Region  string `json:"region"`
	LogoUrl string `json:"logo_url"`
}

type ChampionshipRegions uint8

const (
	Americas ChampionshipRegions = iota
	EMEA
	Pacific
	Worlds
)

func (s ChampionshipRegions) String() string {
	switch s {
	case Americas:
		return "americas"
	case EMEA:
		return "emea"
	case Pacific:
		return "pacific"
	case Worlds:
		return "worlds"
	}
	return "unknown"
}

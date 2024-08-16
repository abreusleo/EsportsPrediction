package championship

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetPossibleChampionships() []GetPossibleChampionshipsDto {
	dto := GetPossibleChampionshipsDto{
		Id:      0,
		Name:    "VCT Americas",
		Region:  ChampionshipRegions.String(Americas),
		LogoUrl: "https://cdn.thespike.gg/VCT%25202023%2FVCT23-AMERICAS_1678786271315.png",
	}

	secondDto := GetPossibleChampionshipsDto{
		Id:      1,
		Name:    "VCT EMEA",
		Region:  ChampionshipRegions.String(EMEA),
		LogoUrl: "https://cdn.thespike.gg/VCT%25202023%2FVCT23-AMERICAS_1678786271315.png",
	}

	dtoList := []GetPossibleChampionshipsDto{dto, secondDto}

	return dtoList
}

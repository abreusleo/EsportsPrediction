package team

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetPossibleTeams() []GetPossibleTeamsDto {
	dto := GetPossibleTeamsDto{Id: 0, Name: "Loud", Acronym: "LOUD", LogoUrl: "https://yt3.googleusercontent.com/hBufrTeLwDxrrZOjMsQEooQrne6pRAhdSFhOivYfq5gywsmpYmmRLK24YaDZKN3AdGNvX0Z0=s900-c-k-c0x00ffffff-no-rj"}
	secondDto := GetPossibleTeamsDto{Id: 1, Name: "NRG", Acronym: "NRG", LogoUrl: "https://yt3.googleusercontent.com/6h5O0YcxI5_liolywkyVIgW5ZC9hrcWe42vMDqBkEbZrLrsdif3kvAUj04_GQiswh4zqFKsf=s900-c-k-c0x00ffffff-no-rj"}

	dtoList := []GetPossibleTeamsDto{dto, secondDto}

	return dtoList
}

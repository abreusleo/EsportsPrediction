package player

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetPossiblePlayers() []GetPossiblePlayersDto {
	dto := GetPossiblePlayersDto{
		Id:      0,
		Name:    "Carlos",
		PhotoUrl: "https://images.unsplash.com/photo-1563409236302-8442b5e644df?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8M3x8ZHVja3xlbnwwfHwwfHx8MA%3D%3D&w=1000&q=80"
	}

	secondDto := GetPossiblePlayersDto{
		Id:      1,
		Name:    "Bessa",
		PhotoUrl: "https://s2-techtudo.glbimg.com/eZA3MfymxGYzGBJK0Ig8EY7ETwA=/0x0:1080x995/984x0/smart/filters:strip_icc()/i.s3.glbimg.com/v1/AUTH_08fbf48bc0524877943fe86e43087e7a/internal_photos/bs/2020/B/l/bRJmJAQiWdlUY2BWM39w/1.-the-rock.jpg"
	}

	dtoList := []GetPossiblePlayersDto{dto, secondDto}

	return dtoList
}

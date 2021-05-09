package services

type Service struct {
	first  string
	second string
}

func NewMultiplierService(first, second string) *Service {
	return &Service{
		first:  first,
		second: second,
	}
}

func (s *Service) Calculate() (string, error) {
	return "", nil
}

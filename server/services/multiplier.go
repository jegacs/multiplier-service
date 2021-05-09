package services

import (
	"fmt"
	"math"
	"math/big"
)

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
	first, _, err := big.ParseFloat(s.first, 10, 256, big.ToNearestEven)
	if err != nil {
		return "", err
	}
	second, _, err := big.ParseFloat(s.second, 10, 256, big.ToNearestEven)
	if err != nil {
		return "", err
	}

	product := new(big.Float).Mul(first, second)

	productTimes100, _ := new(big.Float).Mul(product, big.NewFloat(100)).Float64()

	truncatedTimesProduct := fmt.Sprintf("%.2f", math.Round(productTimes100)/float64(100))

	return truncatedTimesProduct, nil
}

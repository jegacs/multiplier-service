package services

import (
	"fmt"
	"math/big"
)

type Service struct {
	first  string
	second string
}

// Create a new service to compute a multiplication. First and Second are decimal numbers encoded as strings.
func NewMultiplierService(first, second string) *Service {
	return &Service{
		first:  first,
		second: second,
	}
}

// Calculate performs the multiplication of the values stored in the service struct.
func (s *Service) Calculate() (string, error) {
	// Parse the float with bigger precision (256 bits)
	first, _, err := big.ParseFloat(s.first, 10, 256, big.ToNearestEven)
	if err != nil {
		return "", err
	}
	second, _, err := big.ParseFloat(s.second, 10, 256, big.ToNearestEven)
	if err != nil {
		return "", err
	}
	// Perform the multiplication
	product := new(big.Float).Mul(first, second)

	// Truncate the values to two decimal numbers
	// productTimes100, _ := new(big.Float).Mul(product, big.NewFloat(100)).Float64()
	truncatedTimesProduct := fmt.Sprintf("%f", product)

	return truncatedTimesProduct, nil
}

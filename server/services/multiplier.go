package services

import (
	"fmt"
	"math/big"
	"regexp"

	"github.com/jegacs/multiplier-service/errors"
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
	if !s.isFormatCorrect(s.first) {
		return "", errors.ErrBadFormatNumber
	}

	if !s.isFormatCorrect(s.second) {
		return "", errors.ErrBadFormatNumber
	}

	// Parse the float with bigger precision (256 bits)
	first, _, err := big.ParseFloat(s.first, 10, 256, big.ToNearestEven)
	if err != nil {
		return "", err
	}

	if err := s.checkNumberLimits(first); err != nil {
		return "", err
	}

	second, _, err := big.ParseFloat(s.second, 10, 256, big.ToNearestEven)
	if err != nil {
		return "", err
	}

	if err := s.checkNumberLimits(second); err != nil {
		return "", err
	}
	// Perform the multiplication
	product := new(big.Float).Mul(first, second)

	// Truncate the values to two decimal numbers
	// productTimes100, _ := new(big.Float).Mul(product, big.NewFloat(100)).Float64()
	truncatedTimesProduct := fmt.Sprintf("%.2f", product)

	return truncatedTimesProduct, nil
}

// Check if the format of the string encoded decimal number is correct.
func (s *Service) isFormatCorrect(number string) bool {
	r, err := regexp.Compile(`^-{0,1}[0-9]+(\.[0-9]{1,2})?$`)
	if err != nil {
		panic(err)
	}
	return r.MatchString(number)
}

func (s *Service) checkNumberLimits(number *big.Float) error {
	if s.isGreaterThanLimit(number) {
		return errors.ErrGreaterThanLimit
	}

	if s.isLowerThanLimit(number) {
		return errors.ErrSmallerThanLimit
	}

	return nil
}

func (s *Service) isGreaterThanLimit(number *big.Float) bool {
	upperLimit, _, _ := big.ParseFloat("1000.00", 10, 256, big.ToNearestEven)

	return number.Cmp(upperLimit) == 1
}

func (s *Service) isLowerThanLimit(number *big.Float) bool {
	lowerLimit, _, _ := big.ParseFloat("-1000.00", 10, 256, big.ToNearestEven)

	return number.Cmp(lowerLimit) == -1
}

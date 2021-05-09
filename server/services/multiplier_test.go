package services

import (
	"testing"
)

func TestMultiplierService(t *testing.T) {
	t.Run("when calculating product of two numbers", func(t *testing.T) {
		t.Run("they should have right format", func(t *testing.T) {
			s := NewMultiplierService("1.10", "100.99")
			result, err := s.Calculate()
			if err != nil {
				t.Errorf("error should be nil")
			}

			if result == "" {
				t.Errorf("result should not be empty")
			}
		})

		t.Run("given numbers with wrong format, it should return error", func(t *testing.T) {
			s := NewMultiplierService("1.1110", "100.999")
			result, err := s.Calculate()
			if err == nil {
				t.Errorf("error should not be nil")
			}

			if result != "" {
				t.Errorf("result should be empty")
			}
		})

		t.Run("given input greater than upper limit, it should return error", func(t *testing.T) {
			s := NewMultiplierService("-1000.01", "-100")
			result, err := s.Calculate()
			if err == nil {
				t.Errorf("error should not be nil")
			}

			if result != "" {
				t.Errorf("result should be empty")
			}
		})

		t.Run("given input smaller than lower limit, it should return error", func(t *testing.T) {
			s := NewMultiplierService("-1000", "-1000.01")
			result, err := s.Calculate()
			if err == nil {
				t.Errorf("error should not be nil")
			}

			if result != "" {
				t.Errorf("result should be empty")
			}
		})

		t.Run("given right input, it should compute right product", func(t *testing.T) {
			s := NewMultiplierService("33.00", "33.1")
			result, err := s.Calculate()
			if err != nil {
				t.Errorf("error should be nil")
			}

			if result != "1092.30" {
				t.Errorf("result was %v, expected 1093.3", result)
			}
		})

		t.Run("given product result with longer than two decimals, it should truncate it", func(t *testing.T) {
			s := NewMultiplierService("785.12", "900.49")
			// with three decimals, result is 706,992.709
			result, err := s.Calculate()
			if err != nil {
				t.Errorf("error should be nil")
			}

			if result != "706992.71" {
				t.Errorf("result was %v, expected 706992.71", result)
			}
		})
	})

}

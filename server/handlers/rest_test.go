package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jegacs/multiplier-service/dto"
)

func TestHTTPMultiplyHandler(t *testing.T) {
	t.Run("when requesting POST on /multiplier endpoint", func(t *testing.T) {
		t.Run("given right numbers format", func(t *testing.T) {
			payload := &dto.MultiplierRequest{
				First:  "10",
				Second: "10",
			}

			serializedPayload, err := json.Marshal(payload)
			if err != nil {
				t.Errorf("error should be nil, error was %v", err)
			}
			req, err := http.NewRequest("POST", "/multiplier", bytes.NewReader(serializedPayload))
			if err != nil {
				t.Errorf("err should be nil, err was %v", err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(httpMultiplyHandler)
			handler.ServeHTTP(rr, req)

			t.Run("status code should be OK", func(t *testing.T) {
				if rr.Code != http.StatusOK {
					t.Errorf("expected status code 200, given %v", rr.Code)
				}
			})

			t.Run("response should contain product result", func(t *testing.T) {
				response := &dto.MultiplierResult{}
				serializedResponse, err := ioutil.ReadAll(rr.Body)
				if err != nil {
					t.Errorf("error should be nil, error was %v", err)
				}

				err = json.Unmarshal(serializedResponse, response)
				if err != nil {
					t.Errorf("error should be nil, error was %v", err)
				}

				if response.Result == "" {
					t.Errorf("error should not be empty, but it was")
				}

				if response.Result != "100.00" {
					t.Errorf("expected 100, but it was %v", response.Result)
				}
			})

		})

		t.Run("given wrong numbers format", func(t *testing.T) {
			t.Run("given first field is empty", func(t *testing.T) {
				payload := &dto.MultiplierRequest{
					First:  "",
					Second: "100",
				}

				serializedPayload, err := json.Marshal(payload)
				if err != nil {
					t.Errorf("error should be nil, error was %v", err)
				}
				req, err := http.NewRequest("POST", "/multiplier", bytes.NewReader(serializedPayload))
				if err != nil {
					t.Errorf("err should be nil, err was %v", err)
				}

				rr := httptest.NewRecorder()
				handler := http.HandlerFunc(httpMultiplyHandler)
				handler.ServeHTTP(rr, req)
				if rr.Code != http.StatusBadRequest {
					t.Errorf("expected status code 400, given %v", rr.Code)
				}
			})

			t.Run("given second field is empty", func(t *testing.T) {
				payload := &dto.MultiplierRequest{
					First:  "100",
					Second: "",
				}

				serializedPayload, err := json.Marshal(payload)
				if err != nil {
					t.Errorf("error should be nil, error was %v", err)
				}
				req, err := http.NewRequest("POST", "/multiplier", bytes.NewReader(serializedPayload))
				if err != nil {
					t.Errorf("err should be nil, err was %v", err)
				}

				rr := httptest.NewRecorder()
				handler := http.HandlerFunc(httpMultiplyHandler)
				handler.ServeHTTP(rr, req)
				if rr.Code != http.StatusBadRequest {
					t.Errorf("expected status code 400, given %v", rr.Code)
				}

			})
		})

	})

	t.Run("given no numbers", func(t *testing.T) {
		payload := &dto.MultiplierRequest{
			First:  "abcd",
			Second: "abcd",
		}

		serializedPayload, err := json.Marshal(payload)
		if err != nil {
			t.Errorf("error should be nil, error was %v", err)
		}
		req, err := http.NewRequest("POST", "/multiplier", bytes.NewReader(serializedPayload))
		if err != nil {
			t.Errorf("err should be nil, err was %v", err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(httpMultiplyHandler)
		handler.ServeHTTP(rr, req)
		t.Run("with empty fields", func(t *testing.T) {
			if rr.Code != http.StatusInternalServerError {
				t.Errorf("expected status code 500, given %v", rr.Code)
			}
		})
	})

	t.Run("given empty payload", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/multiplier", bytes.NewReader(nil))
		if err != nil {
			t.Errorf("err should be nil, err was %v", err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(httpMultiplyHandler)
		handler.ServeHTTP(rr, req)
		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code 400, given %v", rr.Code)
		}
	})

	t.Run("when requesting on /multiplier using HTTP method other than POST endpoint", func(t *testing.T) {
		t.Run("given method is PUT", func(t *testing.T) {
			req, err := http.NewRequest("PUT", "/multiplier", bytes.NewReader(nil))
			if err != nil {
				t.Errorf("err should be nil, err was %v", err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(httpMultiplyHandler)
			handler.ServeHTTP(rr, req)
			if rr.Code != http.StatusMethodNotAllowed {
				t.Errorf("expected status code 405, given %v", rr.Code)
			}
		})
	})

}

package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jegacs/multiplier-service/dto"
	"github.com/jegacs/multiplier-service/server/services"
)

type RESTServer struct{}

func (s *RESTServer) SetMultiplyHandler() {
	http.HandleFunc("/multiplier", httpMultiplyHandler)
}

func httpMultiplyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Request received! ")
	switch r.Method {
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "It seems your request is malformed", http.StatusBadRequest)
			return
		}

		request := &dto.MultiplierRequest{}
		err = json.Unmarshal(body, request)
		if err != nil {
			http.Error(w, "It seems your request is malformed", http.StatusBadRequest)
			return
		}

		if request.First == "" || request.Second == "" {
			http.Error(w, "Fields cannot be empty", http.StatusBadRequest)
			return

		}
		log.Printf("First: %v, Second: %v", request.First, request.Second)
		service := services.NewMultiplierService(request.First, request.Second)
		result, err := service.Calculate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := &dto.MultiplierResult{
			Result: result,
		}
		serializedResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "There was an unexpected error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(serializedResponse)
	default:
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
}

func RunHTTPServer(addr string) {
	server := RESTServer{}
	server.SetMultiplyHandler()
	log.Printf("Starting server in %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

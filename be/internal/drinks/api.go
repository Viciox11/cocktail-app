package drinks

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ioartigiano/ioartigiano-be/internal/entity"
)

func RegisterHandlers(router *mux.Router, service Service) {
	router.HandleFunc("/drinks", getDrinks(service)).Methods(http.MethodGet, http.MethodOptions)
}

func getDrinks(service Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req entity.Drink

		err := json.NewDecoder(r.Body).Decode(&req)

		if err != nil {
			log.Error(err)
			return
		}

		err = service.Validate(&req)

		if err != nil {
			log.Error(err)
			return
		}

		product, err := service.GetDrinks(req)
		if err != nil {
			log.Error(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		err = json.NewEncoder(w).Encode(product)

		if err != nil {
			log.Error(err)
			return
		}
	}
}
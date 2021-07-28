package ingredients

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func RegisterHandlers(router *mux.Router, service Service) {
	router.HandleFunc("/ingredients", getIngredients(service)).Methods(http.MethodGet)
}

func getIngredients(service Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		order, err := service.GetIngredients()
		if err != nil {
			log.Error("getIngredients: ", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(order)
		if err != nil {
			log.Error("getIngredients: ", err)
			return
		}
	}
}
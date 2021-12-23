package app

import (
	"encoding/json"
	"net/http"

	"github.com/SaravananPitchaimuthu/Fruits/Fruits/service"
	"github.com/gorilla/mux"
)

type Fruits struct {
	Name     string
	Price    string
	Quantity string
}

type CustomHandlers struct {
	service service.FruitService
}

func (ch CustomHandlers) getAllFruits(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	fruits, err := ch.service.GetAllFruits(status)
	if err != nil {
		WriteResponse(w, err.Code, err.Message)

	} else {
		WriteResponse(w, http.StatusOK, fruits)
	}
}

func (ch CustomHandlers) getFruitById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["fruit_id"]

	fruit, err := ch.service.GetFruit(id)
	if err != nil {
		WriteResponse(w, err.Code, err.Message)

	} else {
		WriteResponse(w, http.StatusOK, fruit)
	}
}

func WriteResponse(w http.ResponseWriter, Code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(Code)
	json.NewEncoder(w).Encode(data)
}

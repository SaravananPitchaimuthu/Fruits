package app

import (
	"encoding/json"
	"net/http"

	"github.com/SaravananPitchaimuthu/Fruits/Fruits/dto"
	"github.com/SaravananPitchaimuthu/Fruits/Fruits/service"
	"github.com/gorilla/mux"
)

type AccountHandler struct {
	service service.AccountService
}

func (h AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	FruitId := vars["fruit_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.FruitId = FruitId
		account, appError := h.service.NewAccount(request)
		if appError != nil {
			WriteResponse(w, appError.Code, err.Error())
		} else {
			WriteResponse(w, http.StatusCreated, account)
		}
	}

}

func (h AccountHandler) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	// get the account_id and fruit_id
	vars := mux.Vars(r)
	accountId := vars["account_id"]
	fruitId := vars["fruit_id"]

	var request dto.TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.AccountId = accountId
		request.FruitId = fruitId
		account, appError := h.service.MakeTransaction(request)

		if appError != nil {
			WriteResponse(w, appError.Code, appError.Message)
		} else {
			WriteResponse(w, http.StatusOK, account)

		}
	}

}

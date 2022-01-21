package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Selahattinn/go-redis/pkg/api/response"
	"github.com/Selahattinn/go-redis/pkg/model"
	"github.com/gorilla/mux"
)

const (
	emptySuccess = ""
)

//endpoint for get key
func (a *API) GetKey(w http.ResponseWriter, r *http.Request) {
	id := a.getParams(r)
	value, err := a.service.GetKeyService().Get(id)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting key: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	response.Write(w, r, value)
}

//endpoint for get all keys
func (a *API) GetAllKeys(w http.ResponseWriter, r *http.Request) {
	keys, err := a.service.GetKeyService().GetAll()
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error when getting all keys: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	response.Write(w, r, keys)
}

//endpoint for store key
func (a *API) StoreKey(w http.ResponseWriter, r *http.Request) {
	var payload model.Key
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error when storing key: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	if payload.Value == "" || payload.ID == "" {
		response.Errorf(w, r, fmt.Errorf("error when storing key: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	err = a.service.GetKeyService().Store(payload)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error when storing key: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	response.Write(w, r, emptySuccess)

}

//endpoint for update key
func (a *API) UpdateKey(w http.ResponseWriter, r *http.Request) {
	var payload model.Key
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error when updating key: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	if payload.Value == "" || payload.ID == "" {
		response.Errorf(w, r, fmt.Errorf("error when updating key: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	err = a.service.GetKeyService().Update(payload)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error when updating key: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	response.Write(w, r, "")
}

//endpoint for delete key
func (a *API) DeleteKey(w http.ResponseWriter, r *http.Request) {
	id := a.getParams(r)
	err := a.service.GetKeyService().Delete(id)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error when deleting key: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	response.Write(w, r, "")
}

//endpoint for delete all keys
func (a *API) DeleteAllKeys(w http.ResponseWriter, r *http.Request) {
	err := a.service.GetKeyService().DeleteAll()
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error when deleting all keys: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	response.Write(w, r, "")
}

//endpoint for check key
func (a *API) CheckKey(w http.ResponseWriter, r *http.Request) {
	id := a.getParams(r)
	found, err := a.service.GetKeyService().CheckExist(id)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error when checking key: %v", err), http.StatusBadRequest, err.Error())
		return
	}
	if found {
		response.Write(w, r, found)
		return
	}
	response.Errorf(w, r, fmt.Errorf("error when checking key: %v", err), http.StatusBadRequest, "")

}

//Internal function for getting parameters from request
func (a *API) getParams(r *http.Request) string {
	params := mux.Vars(r)
	fmt.Println(params["id"])
	return params["id"]
}

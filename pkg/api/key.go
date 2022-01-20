package api

import (
	"net/http"

	"github.com/Selahattinn/go-redis/pkg/api/response"
)

//endpoint for get key
func (a *API) GetKey(w http.ResponseWriter, r *http.Request) {

	response.Write(w, r, "jobs")
}

//endpoint for get all keys
func (a *API) GetAllKeys(w http.ResponseWriter, r *http.Request) {

	response.Write(w, r, "jobs")
}

//endpoint for store key
func (a *API) StoreKey(w http.ResponseWriter, r *http.Request) {

	response.Write(w, r, "jobs")
}

//endpoint for update key
func (a *API) UpdateKey(w http.ResponseWriter, r *http.Request) {

	response.Write(w, r, "jobs")
}

//endpoint for delete key
func (a *API) DeleteKey(w http.ResponseWriter, r *http.Request) {

	response.Write(w, r, "jobs")
}

//endpoint for delete all keys
func (a *API) DeleteAllKeys(w http.ResponseWriter, r *http.Request) {

	response.Write(w, r, "jobs")
}

//endpoint for check key
func (a *API) CheckKey(w http.ResponseWriter, r *http.Request) {

	response.Write(w, r, "jobs")
}

package api

import (
	"net/http"

	"github.com/Selahattinn/go-redis/pkg/api/response"
	"github.com/Selahattinn/go-redis/pkg/service"
	"github.com/gorilla/mux"
)

// API represents the structure of the API
type API struct {
	Router  *mux.Router
	service service.Service
}

// New returns the api settings
func New(router *mux.Router, svc service.Service) (*API, error) {
	api := &API{
		//db:     db,
		Router:  router,
		service: svc,
	}

	// Endpoint for browser preflight requests
	api.Router.Methods("OPTIONS").HandlerFunc(api.corsMiddleware(api.preflightHandler))

	// Endpoint for healtcheck
	api.Router.HandleFunc("/api/v1/health", api.corsMiddleware(api.logMiddleware(api.healthHandler))).Methods("GET")

	// Endpoints for redis
	api.Router.HandleFunc("/keys", api.corsMiddleware(api.logMiddleware(api.StoreKey))).Methods("POST")
	api.Router.HandleFunc("/keys", api.corsMiddleware(api.logMiddleware(api.GetAllKeys))).Methods("GET")
	api.Router.HandleFunc("/keys", api.corsMiddleware(api.logMiddleware(api.DeleteAllKeys))).Methods("DELETE")
	api.Router.HandleFunc("/keys/{id}", api.corsMiddleware(api.logMiddleware(api.CheckKey))).Methods("HEAD")
	api.Router.HandleFunc("/keys/{id}", api.corsMiddleware(api.logMiddleware(api.GetKey))).Methods("GET")
	api.Router.HandleFunc("/keys/{id}", api.corsMiddleware(api.logMiddleware(api.UpdateKey))).Methods("PUT")
	api.Router.HandleFunc("/keys/{id}", api.corsMiddleware(api.logMiddleware(api.DeleteKey))).Methods("DELETE")

	return api, nil
}

func (a *API) healthHandler(w http.ResponseWriter, r *http.Request) {
	response.Write(w, r, struct {
		Status string `json:"status"`
	}{
		"ok",
	})

	return
}

func (a *API) preflightHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	return
}

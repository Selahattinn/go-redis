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
	api.Router.HandleFunc("/api/v1/key", api.corsMiddleware(api.logMiddleware(api.healthHandler))).Methods("GET")

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

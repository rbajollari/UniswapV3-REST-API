package api

import (
	"github.com/gorilla/mux"

	"github.com/rbajollari/UniswapV3-REST-API/api/handlers"
)

type API struct {
	Router *mux.Router
}

func (a *API) Initialize() {
	a.Router = mux.NewRouter()

	// Create Endpoints
	a.Router.HandleFunc("/", handlers.HealthCheck).Methods("GET")
	a.Router.HandleFunc("/tokenpools", handlers.TokenPools).Methods("GET")
	a.Router.HandleFunc("/tokenvolume", handlers.TokenVolume).Methods("GET")
	a.Router.HandleFunc("/blockswaps", handlers.BlockSwaps).Methods("GET")
}


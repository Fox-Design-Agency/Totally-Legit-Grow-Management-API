// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routes

import (
	"net/http"
	"totally-legit-grow-management/v1/internal/server"

	"github.com/gorilla/mux"
)

func growLocationRoutes(r *mux.Router, svr *server.Server) {
	r.HandleFunc("/api/v1/growing-location", svr.GrowingLocation.CreateGrowingLocation).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/growing-location", svr.GrowingLocation.DeleteGrowingLocation).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/growing-location", svr.GrowingLocation.EditGrowingLocation).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/growing-location", svr.GrowingLocation.GetGrowingLocation).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/growing-locations", svr.GrowingLocation.GetAllGrowingLocations).Methods(http.MethodGet)
}

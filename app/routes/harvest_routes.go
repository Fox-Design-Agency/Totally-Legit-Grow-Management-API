// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routes

import (
	"totally-legit-grow-management/v1/pkg/server"

	"github.com/gorilla/mux"
)

func harvestRoutes(r *mux.Router, svr *server.Server) {
	// r.HandleFunc("/api/v1/harvest", svr.CreateHarvest).Methods(http.MethodPost)
	// r.HandleFunc("/api/v1/harvest", svr.DeleteHarvest).Methods(http.MethodDelete)
	// r.HandleFunc("/api/v1/harvest", svr.EditHarvest).Methods(http.MethodPut)
	// r.HandleFunc("/api/v1/harvest", svr.GetHarvest).Methods(http.MethodGet)
	// r.HandleFunc("/api/v1/harvests", svr.GetAllHarvests).Methods(http.MethodGet)
}

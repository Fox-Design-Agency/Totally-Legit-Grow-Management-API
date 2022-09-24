// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routes

import (
	"net/http"
	"totally-legit-grow-management/v1/pkg/server"

	"github.com/gorilla/mux"
)

func plantLifeCycleRoutes(r *mux.Router, svr *server.Server) {
	r.HandleFunc("/api/v1/plant-life-cycle", svr.CreatePlantLifeCycle).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/plant-life-cycle", svr.DeletePlantLifeCycle).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/plant-life-cycle", svr.EditPlantLifeCycle).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/plant-life-cycle", svr.GetPlantLifeCycleByID).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/plant-life-cycles", svr.GetAllPlantLifeCycles).Methods(http.MethodGet)
}

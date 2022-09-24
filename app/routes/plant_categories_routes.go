// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routes

import (
	"net/http"
	"totally-legit-grow-management/v1/pkg/server"

	"github.com/gorilla/mux"
)

func plantCategoriesRoutes(r *mux.Router, svr *server.Server) {
	r.HandleFunc("/api/v1/plant-category", svr.CreatePlantCategory).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/plant-category", svr.DeletePlantCategory).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/plant-category", svr.EditPlantCategory).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/plant-category", svr.GetPlantCategory).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/plant-categories", svr.GetAllPlantCategories).Methods(http.MethodGet)
}

// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routes

import (
	"net/http"
	"totally-legit-grow-management/v1/pkg/server"

	"github.com/gorilla/mux"
)

func nutrientRoutes(r *mux.Router, svr *server.Server) {
	r.HandleFunc("/api/v1/nutrient", svr.CreateNutrient).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/nutrient", svr.DeleteNutrient).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/nutrient", svr.EditNutrient).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/nutrient", svr.GetNutrient).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/nutrients", svr.GetAllNutrients).Methods(http.MethodGet)
}

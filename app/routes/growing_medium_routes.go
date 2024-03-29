// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routes

import (
	"net/http"
	"totally-legit-grow-management/v1/internal/server"

	"github.com/gorilla/mux"
)

func growingMediumRoutes(r *mux.Router, svr *server.Server) {
	r.HandleFunc("/api/v1/growing-medium", svr.GrowingMedium.CreateGrowingMedium).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/growing-medium", svr.GrowingMedium.DeleteGrowingMedium).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/growing-medium", svr.GrowingMedium.EditGrowingMedium).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/growing-medium", svr.GrowingMedium.GetGrowingMedium).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/growing-mediums", svr.GrowingMedium.GetAllGrowingMediums).Methods(http.MethodGet)
}

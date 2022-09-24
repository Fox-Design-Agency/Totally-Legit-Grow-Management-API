// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routes

import (
	"net/http"
	"totally-legit-grow-management/v1/internal/server"

	"github.com/gorilla/mux"
)

func seedRoutes(r *mux.Router, svr *server.Server) {
	r.HandleFunc("/api/v1/seed", svr.Seed.CreateSeed).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/seed", svr.Seed.DeleteSeed).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/seed", svr.Seed.EditSeed).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/seed", svr.Seed.GetSeed).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/seeds", svr.Seed.GetAllSeeds).Methods(http.MethodGet)
}

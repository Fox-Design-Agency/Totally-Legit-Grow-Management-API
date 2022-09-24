// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routes

import (
	"net/http"
	"totally-legit-grow-management/v1/internal/server"

	"github.com/gorilla/mux"
)

func growingGroupRoutes(r *mux.Router, svr *server.Server) {
	r.HandleFunc("/api/v1/growing-group", svr.CreateGrowingGroup).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/growing-group", svr.DeleteGrowingGroup).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/growing-group", svr.EditGrowingGroup).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/growing-group", svr.GetGrowingGroup).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/growing-groups", svr.GetAllGrowingGroups).Methods(http.MethodGet)
}

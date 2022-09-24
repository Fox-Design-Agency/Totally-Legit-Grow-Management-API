// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routes

import (
	"net/http"
	"totally-legit-grow-management/v1/internal/server"

	"github.com/gorilla/mux"
)

func inventoryRoutes(r *mux.Router, svr *server.Server) {
	r.HandleFunc("/api/v1/inventory", svr.CreateInventory).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/inventory", svr.DeleteInventory).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/inventory", svr.EditInventory).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/inventory", svr.GetInventory).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/inventories", svr.GetAllInventories).Methods(http.MethodGet)
}

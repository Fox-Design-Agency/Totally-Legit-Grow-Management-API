// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routes

import (
	"net/http"
	"totally-legit-grow-management/v1/pkg/server"

	"github.com/gorilla/mux"
)

func organizationRoutes(r *mux.Router, svr *server.Server) {
	r.HandleFunc("/api/v1/organization", svr.CreateOrganization).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/organization", svr.DeleteOrganization).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/organization", svr.EditOrganization).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/organization", svr.GetOrganization).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/organizations", svr.GetAllOrganizations).Methods(http.MethodGet)
}

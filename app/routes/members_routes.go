// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routes

import (
	"net/http"
	"totally-legit-grow-management/v1/pkg/server"

	"github.com/gorilla/mux"
)

func membersRoutes(r *mux.Router, svr *server.Server) {
	r.HandleFunc("/api/v1/member", svr.CreateMember).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/member", svr.DeleteMember).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/member", svr.EditMember).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/member", svr.GetMember).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/members", svr.GetAllMembers).Methods(http.MethodGet)
}

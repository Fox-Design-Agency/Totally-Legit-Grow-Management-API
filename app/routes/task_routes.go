// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routes

import (
	"net/http"
	"totally-legit-grow-management/v1/internal/server"

	"github.com/gorilla/mux"
)

func taskRoutes(r *mux.Router, svr *server.Server) {
	r.HandleFunc("/api/v1/task", svr.Task.CreateTask).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/task", svr.Task.DeleteTask).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/task", svr.Task.EditTask).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/task", svr.Task.GetTask).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/tasks", svr.Task.GetAllTasks).Methods(http.MethodGet)
}

// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package taskhandlers

import (
	"net/http"
	taskcontrol "totally-legit-grow-management/v1/internal/logic/taskControl"
)

var _ ITaskHandler = &Task{}

type ITaskHandler interface {
	CreateTask(w http.ResponseWriter, r *http.Request)
	DeleteTask(w http.ResponseWriter, r *http.Request)
	EditTask(w http.ResponseWriter, r *http.Request)
	GetTask(w http.ResponseWriter, r *http.Request)
	GetAllTasks(w http.ResponseWriter, r *http.Request)
}

type Task struct {
	Logic taskcontrol.ITasksLogic
}

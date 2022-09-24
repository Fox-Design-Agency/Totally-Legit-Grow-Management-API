// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package tasks

import (
	routemodels "totally-legit-grow-management/v1/internal/route-models"

	"github.com/jmoiron/sqlx"
)

var _ ITasksDB = &TaskPersistence{}

//
type ITasksDB interface {
	CreateTask(*routemodels.CreateTaskRequest) (*routemodels.CreateTaskResponse, error)
	DeleteTask(*routemodels.DeleteTaskRequest) error
	EditTask(*routemodels.EditTaskRequest) (*routemodels.EditTaskResponse, error)
	GetTask(*routemodels.GetTaskRequest) (*routemodels.GetTaskResponse, error)
	GetAllTasks(*routemodels.GetAllTasksRequest) (*routemodels.GetAllTasksResponse, error)
}

type TaskPersistence struct {
	Postgres *sqlx.DB
}

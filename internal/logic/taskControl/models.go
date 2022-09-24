// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package taskcontrol

import (
	"totally-legit-grow-management/v1/internal/persistence"
	routemodels "totally-legit-grow-management/v1/internal/route-models"
)

var _ ITasksLogic = &TaskControl{}

//
type ITasksLogic interface {
	CreateTask(*routemodels.CreateTaskRequest) (*routemodels.CreateTaskResponse, error)
	DeleteTask(*routemodels.DeleteTaskRequest) error
	EditTask(*routemodels.EditTaskRequest) (*routemodels.EditTaskResponse, error)
	GetTask(*routemodels.GetTaskRequest) (*routemodels.GetTaskResponse, error)
	GetAllTasks(*routemodels.GetAllTasksRequest) (*routemodels.GetAllTasksResponse, error)
}

type TaskControl struct {
	Persistence persistence.ITasksDB
}

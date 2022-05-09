// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routemodels

type CreateTaskRequest struct {
	DisplayName string `json:"displayName"`
}

type CreateTaskResponse struct {
	ID string `json:"id" db:"id"`
}

type DeleteTaskRequest struct {
	ID string `json:"id"`
}

type EditTaskRequest struct {
	DisplayName string `json:"displayName"`
}

type EditTaskResponse struct {
	Task
}

type GetTaskRequest struct {
	ID string `json:"id"`
}

type GetTaskResponse struct {
	Task
}

type GetAllTasksRequest struct {
	OrganizationID string `json:"organizationID"`
}

type GetAllTasksResponse struct {
	Tasks []Task
}

// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package growinggroupcontrol

import (
	growinggroups "totally-legit-grow-management/v1/internal/persistence/growingGroups"
	routemodels "totally-legit-grow-management/v1/internal/route-models"
)

var _ IGrowingGroupsLogic = &GrowingGroupControl{}

//
type IGrowingGroupsLogic interface {
	CreateGrowingGroup(*routemodels.CreateGrowingGroupRequest) (*routemodels.CreateGrowingGroupResponse, error)
	DeleteGrowingGroup(*routemodels.DeleteGrowingGroupRequest) error
	EditGrowingGroup(*routemodels.EditGrowingGroupRequest) (*routemodels.EditGrowingGroupResponse, error)
	GetGrowingGroupByID(*routemodels.GetGrowingGroupRequest) (*routemodels.GetGrowingGroupResponse, error)
	GetAllGrowingGroupsByOrganizationID(*routemodels.GetAllGrowingGroupsRequest) (*routemodels.GetAllGrowingGroupsResponse, error)
}

type GrowingGroupControl struct {
	Persistence growinggroups.IGrowingGroupsDB
}

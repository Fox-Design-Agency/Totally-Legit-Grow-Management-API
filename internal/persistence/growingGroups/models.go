// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package growinggroups

import (
	routemodels "totally-legit-grow-management/v1/internal/route-models"

	"github.com/jmoiron/sqlx"
)

var _ IGrowingGroupsDB = &GrowingGroupPersistence{}

//
type IGrowingGroupsDB interface {
	CreateGrowingGroup(*routemodels.CreateGrowingGroupRequest) (*routemodels.CreateGrowingGroupResponse, error)
	CreateGrowingGroupWithTransaction(*sqlx.Tx, *routemodels.CreateGrowingGroupRequest) (*routemodels.CreateGrowingGroupResponse, error)
	DeleteGrowingGroup(*routemodels.DeleteGrowingGroupRequest) error
	EditGrowingGroup(*routemodels.EditGrowingGroupRequest) (*routemodels.EditGrowingGroupResponse, error)
	GetGrowingGroupByID(*routemodels.GetGrowingGroupRequest) (*routemodels.GetGrowingGroupResponse, error)
	GetGrowingGroupByIDWithTransaction(*sqlx.Tx, *routemodels.GetGrowingGroupRequest) (*routemodels.GetGrowingGroupResponse, error)
	GetAllGrowingGroupsByOrganizationID(*routemodels.GetAllGrowingGroupsRequest) (*routemodels.GetAllGrowingGroupsResponse, error)
	GetAllGrowingGroupsByOrganizationIDWithTransaction(*sqlx.Tx, *routemodels.GetAllGrowingGroupsRequest) (*routemodels.GetAllGrowingGroupsResponse, error)
}

type GrowingGroupPersistence struct {
	Postgres *sqlx.DB
}

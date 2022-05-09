// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routemodels

type CreateGrowingGroupRequest struct {
	DisplayName    string `json:"displayName"`
	OrganizationID string `json:"organizationID"`
}

type CreateGrowingGroupResponse struct {
	ID string `json:"id" db:"id"`
}

type DeleteGrowingGroupRequest struct {
	ID string `json:"id"`
}

type EditGrowingGroupRequest struct {
	DisplayName string `json:"displayName"`
}

type EditGrowingGroupResponse struct {
	GrowingGroup
}

type GetGrowingGroupRequest struct {
	ID string `json:"id"`
}

type GetGrowingGroupResponse struct {
	GrowingGroup
}

type GetAllGrowingGroupsRequest struct {
	OrganizationID string `json:"organizationID"`
}

type GetAllGrowingGroupsResponse struct {
	Groups []GrowingGroup
}

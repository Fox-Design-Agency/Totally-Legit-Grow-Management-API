// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routemodels

type CreateMemberRequest struct {
	DisplayName string `json:"displayName"`
}

type CreateMemberResponse struct {
	ID string `json:"id" db:"id"`
}

type DeleteMemberRequest struct {
	ID string `json:"id"`
}

type EditMemberRequest struct {
	DisplayName string `json:"displayName"`
}

type EditMemberResponse struct {
	Member
}

type GetMemberRequest struct {
	ID string `json:"id"`
}

type GetMemberResponse struct {
	Member
}

type GetAllMembersRequest struct {
	OrganizationID string `json:"organizationID"`
}

type GetAllMembersResponse struct {
	Members []Member
}

// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routemodels

type SlimGrowingGroup struct {
	ID             string `json:"id" db:"id"`
	OrganizationID string `json:"organizationID" db:"organization_id"`
	DisplayName    string `json:"displayName" db:"display_name"`
}

type GrowingGroup struct {
	SlimGrowingGroup
	CreatedAtMember
	UpdatedAtMember
}

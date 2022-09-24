// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routemodels

import "github.com/lib/pq"

type SlimOrganization struct {
	ID          string `json:"id" db:"id"`
	DisplayName string `json:"displayName" db:"display_name"`
}

type Organization struct {
	SlimPlantCategory
	CreatedAt pq.NullTime `json:"createdAt" db:"created_at"`
	CreatedAtMember
	UpdatedAt pq.NullTime `json:"updatedAt" db:"updated_by"`
	UpdatedAtMember
}

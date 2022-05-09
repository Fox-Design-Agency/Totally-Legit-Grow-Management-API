// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routemodels

type SlimGrowSpot struct {
	ID             string `json:"id" db:"id"`
	DisplayName    string `json:"displayName" db:"display_name"`
	GrowingLevelID string `json:"growingLevelID" db:"growing_level_id"`
}

type GrowSpot struct {
	SlimGrowSpot
	CreatedAtMember
	UpdatedAtMember
}

// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routemodels

type CreateGrowingLevelRequest struct {
	DisplayName       string `json:"displayName"`
	GrowingLocationID string `json:"growingLocationIDs"`
}

type CreateGrowingLevelResponse struct {
	ID string `json:"id" db:"id"`
}

type DeleteGrowingLevelRequest struct {
	ID string `json:"id"`
}

type EditGrowingLevelRequest struct {
	DisplayName string `json:"displayName"`
}

type EditGrowingLevelResponse struct {
	GrowingLevel
}

type GetGrowingLevelRequest struct {
	ID string `json:"id"`
}

type GetGrowingLevelResponse struct {
	GrowingLevel
}

type GetAllGrowingLevelsRequest struct {
	GrowingLevelID string `json:"growingLevelID"`
}

type GetAllGrowingLevelsResponse struct {
	GrowingLevels []GrowingLevel
}

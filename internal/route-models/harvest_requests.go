// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routemodels

type CreateHarvestRequest struct {
	DisplayName string `json:"displayName"`
}

type CreateHarvestResponse struct {
	ID string `json:"id" db:"id"`
}

type DeleteHarvestRequest struct {
	ID string `json:"id"`
}

type EditHarvestRequest struct {
	DisplayName string `json:"displayName"`
}

type EditHarvestResponse struct {
	Harvest
}

type GetHarvestRequest struct {
	ID string `json:"id"`
}

type GetHarvestResponse struct {
	Harvest
}

type GetAllHarvestsRequest struct {
	OrganizationID string `json:"organizationID"`
}

type GetAllHarvestsResponse struct {
	Harvests []Harvest
}

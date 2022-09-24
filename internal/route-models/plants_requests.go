// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routemodels

type CreatePlantRequest struct {
	DisplayName string `json:"displayName"`
}

type CreatePlantResponse struct {
	ID string `json:"id" db:"id"`
}

type DeletePlantRequest struct {
	ID string `json:"id"`
}

type EditPlantRequest struct {
	DisplayName string `json:"displayName"`
}

type EditPlantResponse struct {
	Plant
}

type GetPlantRequest struct {
	ID string `json:"id"`
}

type GetPlantResponse struct {
	Plant
}

type GetAllPlantsRequest struct {
	OrganizationID string `json:"organizationID"`
}

type GetAllPlantsResponse struct {
	Plants []Plant
}

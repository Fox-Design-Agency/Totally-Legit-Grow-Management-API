// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routemodels

type CreateGrowSpotPlantRequest struct {
	DisplayName string `json:"displayName"`
}

type CreateGrowSpotPlantResponse struct {
	ID string `json:"id" db:"id"`
}

type DeleteGrowSpotPlantRequest struct {
	ID string `json:"id"`
}

type EditGrowSpotPlantRequest struct {
	DisplayName string `json:"displayName"`
}

type EditGrowSpotPlantResponse struct {
	GrowSpotPlant
}

type GetGrowSpotPlantRequest struct {
	ID string `json:"id"`
}

type GetGrowSpotPlantResponse struct {
	GrowSpotPlant
}

type GetAllGrowSpotPlantsRequest struct {
	OrganizationID string `json:"organizationID"`
}

type GetAllGrowSpotPlantsResponse struct {
	GrowSpotPlants []GrowSpotPlant
}

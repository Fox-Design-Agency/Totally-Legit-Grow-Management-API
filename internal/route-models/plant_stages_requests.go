// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routemodels

type CreatePlantStageRequest struct {
	DisplayName string `json:"displayName"`
}

type CreatePlantStageResponse struct {
	ID string `json:"id" db:"id"`
}

type CreatePlantStageCareRequest struct {
	DisplayName string `json:"displayName"`
}

type CreatePlantStageCareResponse struct {
	ID string `json:"id" db:"id"`
}

type CreatePlantStageNutrientsRequest struct {
	DisplayName string `json:"displayName"`
}

type CreatePlantStageNutrientsResponse struct {
	ID string `json:"id" db:"id"`
}

type ConnectPlantStageRequest struct {
	DisplayName string `json:"displayName"`
}

type ConnectPlantStageResponse struct {
	ID string `json:"id" db:"id"`
}

type DeletePlantStageRequest struct {
	ID string `json:"id"`
}

type EditPlantStageRequest struct {
	DisplayName string `json:"displayName"`
}

type EditPlantStageResponse struct {
	PlantStage
}

type GetPlantStageRequest struct {
	ID string `json:"id"`
}

type GetPlantStageResponse struct {
	PlantStage
}

type GetPlantStageCareByIDRequest struct {
	ID string `json:"id"`
}

type GetPlantStageCareByIDResponse struct {
	PlantStage
}

type GetPlantStageNutrientRequest struct {
	OrganizationID string `json:"organizationID"`
}

type GetPlantStageNutrientResponse struct {
	PlantStage
}

type GetPlantStageCareByPlantStageIDRequest struct {
	OrganizationID string `json:"organizationID"`
}

type GetPlantStageCareByPlantStageIDResponse struct {
	PlantStage
}

type GetAllPlantStagesRequest struct {
	OrganizationID string `json:"organizationID"`
}

type GetAllPlantStagesResponse struct {
	PlantStages []PlantStage
}

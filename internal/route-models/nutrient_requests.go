// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routemodels

type CreateNutrientRequest struct {
	DisplayName string `json:"displayName"`
}

type CreateNutrientResponse struct {
	ID string `json:"id" db:"id"`
}

type DeleteNutrientRequest struct {
	ID string `json:"id"`
}

type EditNutrientRequest struct {
	DisplayName string `json:"displayName"`
}

type EditNutrientResponse struct {
	Nutrient
}

type GetNutrientRequest struct {
	ID string `json:"id"`
}

type GetNutrientResponse struct {
	Nutrient
}

type GetAllNutrientsRequest struct {
	OrganizationID string `json:"organizationID"`
}

type GetAllNutrientsResponse struct {
	Nutrients []Nutrient
}

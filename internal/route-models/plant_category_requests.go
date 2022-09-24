// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routemodels

type CreatePlantCategoryRequest struct {
	DisplayName string `json:"displayName"`
}

type CreatePlantCategoryResponse struct {
	ID string `json:"id" db:"id"`
}

type DeletePlantCategoryRequest struct {
	ID string `json:"id"`
}

type EditPlantCategoryRequest struct {
	DisplayName string `json:"displayName"`
}

type EditPlantCategoryResponse struct {
	PlantCategory
}

type GetPlantCategoryRequest struct {
	ID string `json:"id"`
}

type GetPlantCategoryResponse struct {
	PlantCategory
}

type GetAllPlantCategoriesRequest struct {
	OrganizationID string `json:"organizationID"`
}

type GetAllPlantCategoriesResponse struct {
	PlantCategories []PlantCategory
}

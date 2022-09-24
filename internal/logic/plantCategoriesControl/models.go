// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package plantcategoriescontrol

import (
	"totally-legit-grow-management/v1/internal/persistence"
	routemodels "totally-legit-grow-management/v1/internal/route-models"
)

var _ IPlantCategoriesLogic = &PlantCategoryControl{}

//
type IPlantCategoriesLogic interface {
	CreatePlantCategory(*routemodels.CreatePlantCategoryRequest) (*routemodels.CreatePlantCategoryResponse, error)
	DeletePlantCategory(*routemodels.DeletePlantCategoryRequest) error
	EditPlantCategory(*routemodels.EditPlantCategoryRequest) (*routemodels.EditPlantCategoryResponse, error)
	GetPlantCategory(*routemodels.GetPlantCategoryRequest) (*routemodels.GetPlantCategoryResponse, error)
	GetAllPlantCategories(*routemodels.GetAllPlantCategoriesRequest) (*routemodels.GetAllPlantCategoriesResponse, error)
}

type PlantCategoryControl struct {
	Persistence persistence.IPlantCategoriesDB
}

// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package plantcategories

import (
	routemodels "totally-legit-grow-management/v1/internal/route-models"

	"github.com/jmoiron/sqlx"
)

var _ IPlantCategoriesDB = &PlantCategoryPersistence{}

//
type IPlantCategoriesDB interface {
	CreatePlantCategory(*routemodels.CreatePlantCategoryRequest) (*routemodels.CreatePlantCategoryResponse, error)
	DeletePlantCategory(*routemodels.DeletePlantCategoryRequest) error
	EditPlantCategory(*routemodels.EditPlantCategoryRequest) (*routemodels.EditPlantCategoryResponse, error)
	GetPlantCategory(*routemodels.GetPlantCategoryRequest) (*routemodels.GetPlantCategoryResponse, error)
	GetAllPlantCategories(*routemodels.GetAllPlantCategoriesRequest) (*routemodels.GetAllPlantCategoriesResponse, error)
}

type PlantCategoryPersistence struct {
	Postgres *sqlx.DB
}

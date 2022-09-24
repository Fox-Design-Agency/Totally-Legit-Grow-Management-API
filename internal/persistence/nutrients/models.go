// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package nutrients

import (
	routemodels "totally-legit-grow-management/v1/internal/route-models"

	"github.com/jmoiron/sqlx"
)

var _ INutrientsDB = &NutrientPersistence{}

//
type INutrientsDB interface {
	CreateNutrient(*routemodels.CreateNutrientRequest) (*routemodels.CreateNutrientResponse, error)
	DeleteNutrient(*routemodels.DeleteNutrientRequest) error
	EditNutrient(*routemodels.EditNutrientRequest) (*routemodels.EditNutrientResponse, error)
	GetNutrient(*routemodels.GetNutrientRequest) (*routemodels.GetNutrientResponse, error)
	GetAllNutrients(*routemodels.GetAllNutrientsRequest) (*routemodels.GetAllNutrientsResponse, error)
}

type NutrientPersistence struct {
	Postgres *sqlx.DB
}

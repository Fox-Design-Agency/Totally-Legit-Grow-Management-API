// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package nutrientcontrol

import (
	"totally-legit-grow-management/v1/internal/persistence/nutrients"
	routemodels "totally-legit-grow-management/v1/internal/route-models"
)

var _ INutrientsLogic = &NutrientControl{}

//
type INutrientsLogic interface {
	CreateNutrient(*routemodels.CreateNutrientRequest) (*routemodels.CreateNutrientResponse, error)
	DeleteNutrient(*routemodels.DeleteNutrientRequest) error
	EditNutrient(*routemodels.EditNutrientRequest) (*routemodels.EditNutrientResponse, error)
	GetNutrient(*routemodels.GetNutrientRequest) (*routemodels.GetNutrientResponse, error)
	GetAllNutrients(*routemodels.GetAllNutrientsRequest) (*routemodels.GetAllNutrientsResponse, error)
}

type NutrientControl struct {
	Persistence nutrients.INutrientsDB
}

// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package nutrienthandlers

import (
	"net/http"
	nutrientcontrol "totally-legit-grow-management/v1/internal/logic/nutrientControl"
)

var _ INutrientHandler = &Nutrient{}

type INutrientHandler interface {
	CreateNutrient(w http.ResponseWriter, r *http.Request)
	DeleteNutrient(w http.ResponseWriter, r *http.Request)
	EditNutrient(w http.ResponseWriter, r *http.Request)
	GetNutrient(w http.ResponseWriter, r *http.Request)
	GetAllNutrients(w http.ResponseWriter, r *http.Request)
}

type Nutrient struct {
	Logic nutrientcontrol.INutrientsLogic
}

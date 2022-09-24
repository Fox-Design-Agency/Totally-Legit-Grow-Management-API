// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package plantcategoryhandlers

import (
	"net/http"
	"totally-legit-grow-management/v1/internal/logic"
)

type IPlantCategoryHandler interface {
	CreatePlantCategory(w http.ResponseWriter, r *http.Request)
	DeletePlantCategory(w http.ResponseWriter, r *http.Request)
	EditPlantCategory(w http.ResponseWriter, r *http.Request)
	GetPlantCategory(w http.ResponseWriter, r *http.Request)
	GetAllPlantCategories(w http.ResponseWriter, r *http.Request)
}

type PlantCategory struct {
	Logic logic.IPlantCategoriesLogic
}

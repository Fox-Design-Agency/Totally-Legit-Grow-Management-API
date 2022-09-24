// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package plantlifecyclehandlers

import (
	"net/http"
	"totally-legit-grow-management/v1/internal/logic"
)

type IPlantLifeCycleHandler interface {
	CreatePlantLifeCycle(w http.ResponseWriter, r *http.Request)
	DeletePlantLifeCycle(w http.ResponseWriter, r *http.Request)
	EditPlantLifeCycle(w http.ResponseWriter, r *http.Request)
	GetPlantLifeCycleByID(w http.ResponseWriter, r *http.Request)
	GetAllPlantLifeCycles(w http.ResponseWriter, r *http.Request)
}

type PlantLifeCycle struct {
	Logic logic.IPlantLifeCyclesLogic
}

// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package planthandlers

import (
	"net/http"
	plantcontrol "totally-legit-grow-management/v1/internal/logic/plantControl"
)

var _ IPlantHandlers = &Plant{}

type IPlantHandlers interface {
	CreatePlant(w http.ResponseWriter, r *http.Request)
	DeletePlant(w http.ResponseWriter, r *http.Request)
	EditPlant(w http.ResponseWriter, r *http.Request)
	GetPlant(w http.ResponseWriter, r *http.Request)
	GetAllPlants(w http.ResponseWriter, r *http.Request)
}

type Plant struct {
	Logic plantcontrol.IPlantsLogic
}

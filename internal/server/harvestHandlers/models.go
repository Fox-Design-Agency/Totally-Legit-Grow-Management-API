// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package harvesthandlers

import (
	"net/http"
	harvestcontrol "totally-legit-grow-management/v1/internal/logic/harvestsControl"
)

var _ IHarvestHandler = &Harvest{}

type IHarvestHandler interface {
	CreateHarvest(w http.ResponseWriter, r *http.Request)
	DeleteHarvest(w http.ResponseWriter, r *http.Request)
	EditHarvest(w http.ResponseWriter, r *http.Request)
	GetHarvest(w http.ResponseWriter, r *http.Request)
	GetAllHarvests(w http.ResponseWriter, r *http.Request)
}

type Harvest struct {
	Logic harvestcontrol.IHarvestsLogic
}

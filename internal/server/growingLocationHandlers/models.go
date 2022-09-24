// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package growinglocationshandlers

import (
	"net/http"
	"totally-legit-grow-management/v1/internal/logic"
)

type IGrowingLocationHandlers interface {
	CreateGrowingLocation(w http.ResponseWriter, r *http.Request)
	DeleteGrowingLocation(w http.ResponseWriter, r *http.Request)
	EditGrowingLocation(w http.ResponseWriter, r *http.Request)
	GetGrowingLocation(w http.ResponseWriter, r *http.Request)
	GetAllGrowingLocations(w http.ResponseWriter, r *http.Request)
}

type GrowingLocation struct {
	Logic logic.IGrowingLocationsLogic
}

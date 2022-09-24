// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package growingmediumhandlers

import (
	"net/http"
	growingmediumcontrol "totally-legit-grow-management/v1/internal/logic/growingMediumControl"
)

var _ IGrowingMediumHandler = &GrowingMedium{}

type IGrowingMediumHandler interface {
	CreateGrowingMedium(w http.ResponseWriter, r *http.Request)
	DeleteGrowingMedium(w http.ResponseWriter, r *http.Request)
	EditGrowingMedium(w http.ResponseWriter, r *http.Request)
	GetGrowingMedium(w http.ResponseWriter, r *http.Request)
	GetAllGrowingMediums(w http.ResponseWriter, r *http.Request)
}

type GrowingMedium struct {
	Logic growingmediumcontrol.IGrowingMediumsLogic
}

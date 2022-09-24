// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package seedhandlers

import (
	"net/http"
	seedcontrol "totally-legit-grow-management/v1/internal/logic/seedsControl"
)

var _ ISeedHandler = &Seed{}

type ISeedHandler interface {
	CreateSeed(w http.ResponseWriter, r *http.Request)
	DeleteSeed(w http.ResponseWriter, r *http.Request)
	EditSeed(w http.ResponseWriter, r *http.Request)
	GetSeed(w http.ResponseWriter, r *http.Request)
	GetAllSeeds(w http.ResponseWriter, r *http.Request)
}

type Seed struct {
	Logic seedcontrol.ISeedsLogic
}

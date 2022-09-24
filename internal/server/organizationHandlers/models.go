// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package organizationhandlers

import (
	"net/http"
	"totally-legit-grow-management/v1/internal/logic"
)

type IOrganizationHandler interface {
	CreateOrganization(w http.ResponseWriter, r *http.Request)
	DeleteOrganization(w http.ResponseWriter, r *http.Request)
	EditOrganization(w http.ResponseWriter, r *http.Request)
	GetOrganization(w http.ResponseWriter, r *http.Request)
	GetAllOrganizations(w http.ResponseWriter, r *http.Request)
}

type Organization struct {
	Logic logic.IOrganizationsLogic
}

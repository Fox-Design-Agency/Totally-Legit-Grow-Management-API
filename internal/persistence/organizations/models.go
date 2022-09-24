// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package organizations

import (
	routemodels "totally-legit-grow-management/v1/internal/route-models"

	"github.com/jmoiron/sqlx"
)

var _ IOrganizationsDB = &OrganizationPersistence{}

//
type IOrganizationsDB interface {
	CreateOrganization(*routemodels.CreateOrganizationRequest) (*routemodels.CreateOrganizationResponse, error)
	DeleteOrganization(*routemodels.DeleteOrganizationRequest) error
	EditOrganization(*routemodels.EditOrganizationRequest) (*routemodels.EditOrganizationResponse, error)
	GetOrganization(*routemodels.GetOrganizationRequest) (*routemodels.GetOrganizationResponse, error)
	GetAllOrganizations(*routemodels.GetAllOrganizationsRequest) (*routemodels.GetAllOrganizationsResponse, error)
}

type OrganizationPersistence struct {
	Postgres *sqlx.DB
}

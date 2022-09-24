// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routemodels

type CreateOrganizationRequest struct {
	DisplayName string `json:"displayName"`
}

type CreateOrganizationResponse struct {
	ID string `json:"id" db:"id"`
}

type DeleteOrganizationRequest struct {
	ID string `json:"id"`
}

type EditOrganizationRequest struct {
	DisplayName string `json:"displayName"`
}

type EditOrganizationResponse struct {
	Organization
}

type GetOrganizationRequest struct {
	ID string `json:"id"`
}

type GetOrganizationResponse struct {
	Organization
}

type GetAllOrganizationsRequest struct {
}

type GetAllOrganizationsResponse struct {
	Organizations []Organization
}

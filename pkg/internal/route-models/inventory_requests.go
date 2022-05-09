// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routemodels

type CreateInventoryRequest struct {
	DisplayName string `json:"displayName"`
}

type CreateInventoryResponse struct {
	ID string `json:"id" db:"id"`
}

type DeleteInventoryRequest struct {
	ID string `json:"id"`
}

type EditInventoryRequest struct {
	DisplayName string `json:"displayName"`
}

type EditInventoryResponse struct {
	Inventory
}

type GetInventoryRequest struct {
	ID string `json:"id"`
}

type GetInventoryResponse struct {
	Inventory
}

type GetAllInventoriesRequest struct {
	OrganizationID string `json:"organizationID"`
}

type GetAllInventoriesResponse struct {
	Inventories []Inventory
}

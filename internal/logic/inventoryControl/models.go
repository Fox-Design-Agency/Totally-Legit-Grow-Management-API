// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inventorycontrol

import (
	"totally-legit-grow-management/v1/internal/persistence/inventories"
	routemodels "totally-legit-grow-management/v1/internal/route-models"
)

var _ IInventoriesLogic = &InventoryControl{}

//
type IInventoriesLogic interface {
	CreateInventory(*routemodels.CreateInventoryRequest) (*routemodels.CreateInventoryResponse, error)
	DeleteInventory(*routemodels.DeleteInventoryRequest) error
	EditInventory(*routemodels.EditInventoryRequest) (*routemodels.EditInventoryResponse, error)
	GetInventory(*routemodels.GetInventoryRequest) (*routemodels.GetInventoryResponse, error)
	GetAllInventories(*routemodels.GetAllInventoriesRequest) (*routemodels.GetAllInventoriesResponse, error)
}

type InventoryControl struct {
	Persistence inventories.IInventoriesDB
}

// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inventoryhandlers

import (
	"net/http"
	"totally-legit-grow-management/v1/internal/logic"
)

type IInventoryHandler interface {
	CreateInventory(w http.ResponseWriter, r *http.Request)
	DeleteInventory(w http.ResponseWriter, r *http.Request)
	EditInventory(w http.ResponseWriter, r *http.Request)
	GetInventory(w http.ResponseWriter, r *http.Request)
	GetAllInventories(w http.ResponseWriter, r *http.Request)
}

type Inventory struct {
	Logic logic.IInventoriesLogic
}

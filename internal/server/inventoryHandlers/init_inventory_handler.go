// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inventoryhandlers

import (
	inventorycontrol "totally-legit-grow-management/v1/internal/logic/inventoryControl"
)

func InitInventoryHandler(inventoryLogic inventorycontrol.IInventoriesLogic) *Inventory {
	return &Inventory{
		Logic: inventoryLogic,
	}
}

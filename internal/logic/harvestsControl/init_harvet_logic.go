// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package harvestcontrol

import "totally-legit-grow-management/v1/internal/persistence"

func InitHarvestLogic(harvestDB persistence.IHarvestsDB) *HarvestControl {
	return &HarvestControl{
		Persistence: harvestDB,
	}
}

// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package plantlifecyclecontrol

import "totally-legit-grow-management/v1/internal/persistence"

func InitPlantLifecycleLogic(plantLifeCycleDB persistence.IPlantLifeCyclesDB) *PlantLifeCycleControl {
	return &PlantLifeCycleControl{
		Persistence: plantLifeCycleDB,
	}
}

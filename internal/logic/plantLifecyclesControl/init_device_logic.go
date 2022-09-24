// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package plantlifecyclecontrol

import (
	plantlifecycle "totally-legit-grow-management/v1/internal/persistence/plantLifeCycles"
)

func InitPlantLifecycleLogic(plantLifeCycleDB plantlifecycle.IPlantLifeCyclesDB) *PlantLifeCycleControl {
	return &PlantLifeCycleControl{
		Persistence: plantLifeCycleDB,
	}
}

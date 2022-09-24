// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package plantlifecyclehandlers

import (
	plantlifecyclecontrol "totally-legit-grow-management/v1/internal/logic/plantLifecyclesControl"
)

func InitPlantLifecycleHandler(plantLifecycleLogic plantlifecyclecontrol.IPlantLifeCyclesLogic) *PlantLifeCycle {
	return &PlantLifeCycle{
		Logic: plantLifecycleLogic,
	}
}

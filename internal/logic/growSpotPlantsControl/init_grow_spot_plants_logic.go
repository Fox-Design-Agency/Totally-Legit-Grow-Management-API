// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package growspotplantscontrol

import (
	growspotplants "totally-legit-grow-management/v1/internal/persistence/growSpotPlants"
)

func InitGrowSpotPlantLogic(growSpotPlantDB growspotplants.IGrowSpotPlantsDB) *GrowSpotPlantControl {
	return &GrowSpotPlantControl{
		Persistence: growSpotPlantDB,
	}
}

// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package growspotplanthandlers

import growspotplantscontrol "totally-legit-grow-management/v1/internal/logic/growSpotPlantsControl"

func InitGrowSpotPlantHandler(growSpotPlantLogic growspotplantscontrol.IGrowSpotPlantsLogic) *GrowSpotPlants {
	return &GrowSpotPlants{
		Logic: growSpotPlantLogic,
	}
}

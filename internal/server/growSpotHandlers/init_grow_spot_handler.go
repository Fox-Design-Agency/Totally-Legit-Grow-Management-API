// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package growspothandlers

import (
	growspotcontrol "totally-legit-grow-management/v1/internal/logic/growSpotsControl"
)

func InitGrowSpotHandler(growSpotLogic growspotcontrol.IGrowSpotsLogic) *GrowSpot {
	return &GrowSpot{
		Logic: growSpotLogic,
	}
}

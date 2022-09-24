// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package growspotcontrol

import (
	growspots "totally-legit-grow-management/v1/internal/persistence/growSpots"
)

func InitGrowSpotLogic(growSpotDB growspots.IGrowSpotsDB) *GrowSpotControl {
	return &GrowSpotControl{
		Persistence: growSpotDB,
	}
}

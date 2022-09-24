// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package growinglocationshandlers

import (
	growinglocationcontrol "totally-legit-grow-management/v1/internal/logic/growingLocationsControl"
)

func InitGrowingLocationHandler(growingLocationLogic growinglocationcontrol.IGrowingLocationsLogic) *GrowingLocation {
	return &GrowingLocation{
		Logic: growingLocationLogic,
	}
}

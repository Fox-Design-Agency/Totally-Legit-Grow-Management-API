// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package growinglevelhandlers

import (
	growinglevelcontrol "totally-legit-grow-management/v1/internal/logic/growingLevelsControl"
)

func InitGrowingLevelHandler(growingLevelLogic growinglevelcontrol.IGrowingLevelsLogic) *GrowingLevel {
	return &GrowingLevel{
		Logic: growingLevelLogic,
	}
}

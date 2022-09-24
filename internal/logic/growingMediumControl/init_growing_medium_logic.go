// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package growingmediumcontrol

import (
	growingmedium "totally-legit-grow-management/v1/internal/persistence/growingMedium"
)

func InitGrowingMediumLogic(growingMediumDB growingmedium.IGrowingMediumsDB) *GrowingMediumControl {
	return &GrowingMediumControl{
		Persistence: growingMediumDB,
	}
}

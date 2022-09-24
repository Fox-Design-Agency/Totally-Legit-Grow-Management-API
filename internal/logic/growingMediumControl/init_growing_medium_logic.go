// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package growingmediumcontrol

import "totally-legit-grow-management/v1/internal/persistence"

func InitGrowingMediumLogic(growingMediumDB persistence.IGrowingMediumsDB) *GrowingMediumControl {
	return &GrowingMediumControl{
		Persistence: growingMediumDB,
	}
}

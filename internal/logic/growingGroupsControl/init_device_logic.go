// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package growinggroupcontrol

import "totally-legit-grow-management/v1/internal/persistence"

func InitGrowingGroupLogic(growingGroupDB persistence.IGrowingGroupsDB) *GrowingGroupControl {
	return &GrowingGroupControl{
		Persistence: growingGroupDB,
	}
}
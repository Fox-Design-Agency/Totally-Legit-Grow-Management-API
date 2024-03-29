// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package growinggroupcontrol

import growinggroups "totally-legit-grow-management/v1/internal/persistence/growingGroups"

func InitGrowingGroupLogic(growingGroupDB growinggroups.IGrowingGroupsDB) *GrowingGroupControl {
	return &GrowingGroupControl{
		Persistence: growingGroupDB,
	}
}

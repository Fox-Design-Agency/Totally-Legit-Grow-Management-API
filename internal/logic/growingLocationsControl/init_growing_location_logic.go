// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package growinglocationcontrol

import "totally-legit-grow-management/v1/internal/persistence"

func InitGrowingLocationLogic(growingLocationDB persistence.IGrowingLocationsDB) *GrowingLocationControl {
	return &GrowingLocationControl{
		Persistence: growingLocationDB,
	}
}

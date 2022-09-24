// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package planthandlers

import (
	plantcontrol "totally-legit-grow-management/v1/internal/logic/plantControl"
)

func InitPlantHandler(plantLogic plantcontrol.IPlantsLogic) *Plant {
	return &Plant{
		Logic: plantLogic,
	}
}

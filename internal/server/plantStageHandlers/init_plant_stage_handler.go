// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package plantstagehandlers

import (
	plantstagescontrol "totally-legit-grow-management/v1/internal/logic/plantStagesControl"
)

func InitPlantStageHandler(plantStageLogic plantstagescontrol.IPLantStagesLogic) *PlantStage {
	return &PlantStage{
		Logic: plantStageLogic,
	}
}

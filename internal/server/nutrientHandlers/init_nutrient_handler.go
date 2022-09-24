// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package nutrienthandlers

import (
	nutrientcontrol "totally-legit-grow-management/v1/internal/logic/nutrientControl"
)

func InitNutrientHandler(nutrientLogic nutrientcontrol.INutrientsLogic) *Nutrient {
	return &Nutrient{
		Logic: nutrientLogic,
	}
}

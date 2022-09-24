// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package seedhandlers

import (
	seedcontrol "totally-legit-grow-management/v1/internal/logic/seedsControl"
)

func InitSeedHandler(seedLogic seedcontrol.ISeedsLogic) *Seed {
	return &Seed{
		Logic: seedLogic,
	}
}

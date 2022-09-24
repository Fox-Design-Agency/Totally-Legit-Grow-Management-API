// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package memberhandlers

import (
	membercontrol "totally-legit-grow-management/v1/internal/logic/membersControl"
)

func InitMemberHandler(memberLogic membercontrol.IMembersLogic) *Member {
	return &Member{
		Logic: memberLogic,
	}
}

// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package membercontrol

import "totally-legit-grow-management/v1/internal/persistence"

func InitMembersLogic(memberDB persistence.IMembersDB) *MembersControl {
	return &MembersControl{
		Persistence: memberDB,
	}
}

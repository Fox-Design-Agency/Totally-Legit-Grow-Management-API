// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package members

import (
	routemodels "totally-legit-grow-management/v1/internal/route-models"

	"github.com/jmoiron/sqlx"
)

var _ IMembersDB = &MemberPersistence{}

//
type IMembersDB interface {
	CreateMember(*routemodels.CreateMemberRequest) (*routemodels.CreateMemberResponse, error)
	DeleteMember(*routemodels.DeleteMemberRequest) error
	EditMember(*routemodels.EditMemberRequest) (*routemodels.EditMemberResponse, error)
	GetMember(*routemodels.GetMemberRequest) (*routemodels.GetMemberResponse, error)
	GetAllMembers(*routemodels.GetAllMembersRequest) (*routemodels.GetAllMembersResponse, error)
}

type MemberPersistence struct {
	Postgres *sqlx.DB
}

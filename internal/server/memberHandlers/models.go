// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package memberhandlers

import (
	"net/http"
	memberControl "totally-legit-grow-management/v1/internal/logic/membersControl"
)

var _ IMemberHandler = &Member{}

type IMemberHandler interface {
	CreateMember(w http.ResponseWriter, r *http.Request)
	DeleteMember(w http.ResponseWriter, r *http.Request)
	EditMember(w http.ResponseWriter, r *http.Request)
	GetMember(w http.ResponseWriter, r *http.Request)
	GetAllMembers(w http.ResponseWriter, r *http.Request)
}

type Member struct {
	Logic memberControl.IMembersLogic
}

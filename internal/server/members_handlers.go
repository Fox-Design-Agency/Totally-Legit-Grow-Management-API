// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package server

import (
	"encoding/json"
	"log"
	"net/http"
	"totally-legit-grow-management/v1/internal/helpers"
	routemodels "totally-legit-grow-management/v1/internal/route-models"
)

// CreateMember will handle the initial request and then send the formed request
// to the logic layer before returning the request
func (s *Server) CreateMember(w http.ResponseWriter, r *http.Request) {
	/**********************************************************************
	/
	/	Start Telemetry
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Parse Form or Set Vars
	/
	/**********************************************************************/
	var form routemodels.CreateMemberRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}
	/**********************************************************************
	/
	/	Call Logic Layer
	/
	/**********************************************************************/
	resp, err := s.Logic.CreateMember(&form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}
	/**********************************************************************
	/
	/	Return Success
	/
	/**********************************************************************/
	helpers.SendSuccessHeader(w, resp)
}

// DeleteMember will handle the initial request and then send the formed request
// to the logic layer before returning the request
func (s *Server) DeleteMember(w http.ResponseWriter, r *http.Request) {
	/**********************************************************************
	/
	/	Start Telemetry
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Parse Form or Set Vars
	/
	/**********************************************************************/
	var form routemodels.DeleteMemberRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}
	/**********************************************************************
	/
	/	Call Logic Layer
	/
	/**********************************************************************/
	err := s.Logic.DeleteMember(&form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}
	/**********************************************************************
	/
	/	Return Success
	/
	/**********************************************************************/
	helpers.SendSuccessHeader(w, nil)
}

// EditMember will handle the initial request and then send the formed request
// to the logic layer before returning the request
func (s *Server) EditMember(w http.ResponseWriter, r *http.Request) {
	/**********************************************************************
	/
	/	Start Telemetry
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Parse Form or Set Vars
	/
	/**********************************************************************/
	var form routemodels.EditMemberRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}
	/**********************************************************************
	/
	/	Call Logic Layer
	/
	/**********************************************************************/
	resp, err := s.Logic.EditMember(&form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}
	/**********************************************************************
	/
	/	Return Success
	/
	/**********************************************************************/
	helpers.SendSuccessHeader(w, resp)
}

// GetMember will handle the initial request and then send the formed request
// to the logic layer before returning the request
func (s *Server) GetMember(w http.ResponseWriter, r *http.Request) {
	/**********************************************************************
	/
	/	Start Telemetry
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Parse Form or Set Vars
	/
	/**********************************************************************/
	var form routemodels.GetMemberRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}
	/**********************************************************************
	/
	/	Call Logic Layer
	/
	/**********************************************************************/
	resp, err := s.Logic.GetMember(&form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}
	/**********************************************************************
	/
	/	Return Success
	/
	/**********************************************************************/
	helpers.SendSuccessHeader(w, resp)
}

// GetAllMembers will handle the initial request and then send the formed request
// to the logic layer before returning the request
func (s *Server) GetAllMembers(w http.ResponseWriter, r *http.Request) {
	/**********************************************************************
	/
	/	Start Telemetry
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Parse Form or Set Vars
	/
	/**********************************************************************/
	var form routemodels.GetAllMembersRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}
	/**********************************************************************
	/
	/	Call Logic Layer
	/
	/**********************************************************************/
	resp, err := s.Logic.GetAllMembers(&form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}
	/**********************************************************************
	/
	/	Return Success
	/
	/**********************************************************************/
	helpers.SendSuccessHeader(w, resp)
}

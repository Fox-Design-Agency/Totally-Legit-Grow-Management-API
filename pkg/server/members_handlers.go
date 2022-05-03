package server

import (
	"encoding/json"
	"log"
	"net/http"
	"smart-grow-management-api/v1/pkg/internal/helpers"
	routemodels "smart-grow-management-api/v1/pkg/internal/route-models"
)

func (s *Server) CreateMember(w http.ResponseWriter, r *http.Request) {
	var form routemodels.CreateMemberRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	resp, err := s.Persistence.CreateMember(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, resp)
}

func (s *Server) DeleteMember(w http.ResponseWriter, r *http.Request) {
	var form routemodels.DeleteMemberRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	err := s.Persistence.DeleteMember(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, nil)
}

func (s *Server) EditMember(w http.ResponseWriter, r *http.Request) {
	var form routemodels.EditMemberRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	resp, err := s.Persistence.EditMember(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, resp)
}

func (s *Server) GetMember(w http.ResponseWriter, r *http.Request) {
	var form routemodels.GetMemberRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	resp, err := s.Persistence.GetMember(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, resp)
}

func (s *Server) GetAllMembers(w http.ResponseWriter, r *http.Request) {
	var form routemodels.GetAllMembersRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	resp, err := s.Persistence.GetAllMembers(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, resp)
}

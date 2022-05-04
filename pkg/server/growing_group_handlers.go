package server

import (
	"encoding/json"
	"log"
	"net/http"
	"totally-legit-grow-management/v1/pkg/internal/helpers"
	routemodels "totally-legit-grow-management/v1/pkg/internal/route-models"
)

func (s *Server) CreateGrowingGroup(w http.ResponseWriter, r *http.Request) {
	var form routemodels.CreateGrowingGroupRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	resp, err := s.Persistence.CreateGrowingGroup(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, resp)
}

func (s *Server) DeleteGrowingGroup(w http.ResponseWriter, r *http.Request) {
	var form routemodels.DeleteGrowingGroupRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	err := s.Persistence.DeleteGrowingGroup(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, nil)
}

func (s *Server) EditGrowingGroup(w http.ResponseWriter, r *http.Request) {
	var form routemodels.EditGrowingGroupRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	resp, err := s.Persistence.EditGrowingGroup(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, resp)
}

func (s *Server) GetGrowingGroup(w http.ResponseWriter, r *http.Request) {
	var form routemodels.GetGrowingGroupRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	resp, err := s.Persistence.GetGrowingGroup(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, resp)
}

func (s *Server) GetAllGrowingGroups(w http.ResponseWriter, r *http.Request) {
	var form routemodels.GetAllGrowingGroupsRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	resp, err := s.Persistence.GetAllGrowingGroups(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, resp)
}

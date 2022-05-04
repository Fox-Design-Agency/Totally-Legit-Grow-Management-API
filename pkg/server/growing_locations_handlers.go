package server

import (
	"encoding/json"
	"log"
	"net/http"
	"totally-legit-grow-management/v1/pkg/internal/helpers"
	routemodels "totally-legit-grow-management/v1/pkg/internal/route-models"
)

func (s *Server) CreateGrowingLocation(w http.ResponseWriter, r *http.Request) {
	var form routemodels.CreateGrowingLocationRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	resp, err := s.Persistence.CreateGrowingLocation(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, resp)
}

func (s *Server) DeleteGrowingLocation(w http.ResponseWriter, r *http.Request) {
	var form routemodels.DeleteGrowingLocationRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	err := s.Persistence.DeleteGrowingLocation(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, nil)
}

func (s *Server) EditGrowingLocation(w http.ResponseWriter, r *http.Request) {
	var form routemodels.EditGrowingLocationRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	resp, err := s.Persistence.EditGrowingLocation(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, resp)
}

func (s *Server) GetGrowingLocation(w http.ResponseWriter, r *http.Request) {
	var form routemodels.GetGrowingLocationRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	resp, err := s.Persistence.GetGrowingLocation(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, resp)
}

func (s *Server) GetAllGrowingLocations(w http.ResponseWriter, r *http.Request) {
	var form routemodels.GetAllGrowingLocationsRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	resp, err := s.Persistence.GetAllGrowingLocations(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, resp)
}

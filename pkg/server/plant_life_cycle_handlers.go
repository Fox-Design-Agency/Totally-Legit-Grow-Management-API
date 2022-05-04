package server

import (
	"encoding/json"
	"log"
	"net/http"
	"totally-legit-grow-management/v1/pkg/internal/helpers"
	routemodels "totally-legit-grow-management/v1/pkg/internal/route-models"
)

func (s *Server) CreatePlantLifeCycle(w http.ResponseWriter, r *http.Request) {
	var form routemodels.CreatePlantLifeCycleRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	resp, err := s.Persistence.CreatePlantLifeCycle(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, resp)
}

func (s *Server) DeletePlantLifeCycle(w http.ResponseWriter, r *http.Request) {
	var form routemodels.DeletePlantLifeCycleRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	err := s.Persistence.DeletePlantLifeCycle(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, nil)
}

func (s *Server) EditPlantLifeCycle(w http.ResponseWriter, r *http.Request) {
	var form routemodels.EditPlantLifeCycleRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	resp, err := s.Persistence.EditPlantLifeCycle(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, resp)
}

func (s *Server) GetPlantLifeCycleByID(w http.ResponseWriter, r *http.Request) {
	var form routemodels.GetPlantLifeCycleRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	resp, err := s.Persistence.GetPlantLifeCycleByID(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, resp)
}

func (s *Server) GetAllPlantLifeCycles(w http.ResponseWriter, r *http.Request) {
	var form routemodels.GetAllPlantLifeCyclesRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	resp, err := s.Persistence.GetAllPlantLifeCycles(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, resp)
}

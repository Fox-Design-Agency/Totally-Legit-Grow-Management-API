package server

import (
	"encoding/json"
	"log"
	"net/http"
	"smart-grow-management-api/v1/pkg/internal/helpers"
	routemodels "smart-grow-management-api/v1/pkg/internal/route-models"
)

func (s *Server) CreateDevice(w http.ResponseWriter, r *http.Request) {
	var form routemodels.CreateDeviceRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	resp, err := s.Persistence.CreateDevice(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, resp)
}

func (s *Server) CreateDeviceAction(w http.ResponseWriter, r *http.Request) {
	var form routemodels.CreateDeviceActionRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	resp, err := s.Persistence.CreateDeviceAction(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, resp)
}

func (s *Server) CreateGrowingGroupDevice(w http.ResponseWriter, r *http.Request) {
	var form routemodels.CreateGrowingGroupDeviceRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	resp, err := s.Persistence.CreateGrowingGroupDevice(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, resp)
}

func (s *Server) CreateGrowingLocationDevice(w http.ResponseWriter, r *http.Request) {
	var form routemodels.CreateGrowingLocationDeviceRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	resp, err := s.Persistence.CreateGrowingLocationDevice(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, resp)
}

func (s *Server) CreateGrowingLevelDevice(w http.ResponseWriter, r *http.Request) {
	var form routemodels.CreateGrowingLevelDeviceRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	resp, err := s.Persistence.CreateGrowingLevelDevice(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, resp)
}

func (s *Server) CreateGrowingSpotDevice(w http.ResponseWriter, r *http.Request) {
	var form routemodels.CreateGrowingSpotDeviceRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	resp, err := s.Persistence.CreateGrowingSpotDevice(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, resp)
}

func (s *Server) DeleteDevice(w http.ResponseWriter, r *http.Request) {
	var form routemodels.DeleteDeviceRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	err := s.Persistence.DeleteDevice(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, nil)
}

func (s *Server) DeleteGrowingGroupDevice(w http.ResponseWriter, r *http.Request) {
	var form routemodels.DeleteGrowingGroupDeviceRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	err := s.Persistence.DeleteGrowingGroupDevice(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, nil)
}

func (s *Server) DeleteGrowingLocationDevice(w http.ResponseWriter, r *http.Request) {
	var form routemodels.DeleteGrowingLocationDeviceRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	err := s.Persistence.DeleteGrowingLocationDevice(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, nil)
}

func (s *Server) DeleteGrowingLevelDevice(w http.ResponseWriter, r *http.Request) {
	var form routemodels.DeleteGrowingLevelDeviceRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	err := s.Persistence.DeleteGrowingLevelDevice(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, nil)
}

func (s *Server) DeleteGrowingSpotDevice(w http.ResponseWriter, r *http.Request) {
	var form routemodels.DeleteGrowingSpotDeviceRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	err := s.Persistence.DeleteGrowingSpotDevice(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, nil)
}

func (s *Server) EditDevice(w http.ResponseWriter, r *http.Request) {
	var form routemodels.EditDeviceRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	res, err := s.Persistence.EditDevice(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, res)
}

func (s *Server) GetDevice(w http.ResponseWriter, r *http.Request) {
	var form routemodels.GetDeviceRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	res, err := s.Persistence.GetDevice(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, res)
}

func (s *Server) GetDeviceActions(w http.ResponseWriter, r *http.Request) {
	var form routemodels.GetDeviceActionsRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	res, err := s.Persistence.GetDeviceActions(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, res)
}

func (s *Server) GetAllDevices(w http.ResponseWriter, r *http.Request) {
	var form routemodels.GetAllDevicesRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	res, err := s.Persistence.GetAllDevices(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, res)
}

func (s *Server) GetGrowingGroupDevice(w http.ResponseWriter, r *http.Request) {
	var form routemodels.GetGrowingGroupDeviceRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	res, err := s.Persistence.GetGrowingGroupDevice(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, res)
}

func (s *Server) GetGrowingLocationDevice(w http.ResponseWriter, r *http.Request) {
	var form routemodels.GetGrowingLocationDeviceRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	res, err := s.Persistence.GetGrowingLocationDevice(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, res)
}

func (s *Server) GetGrowingLevelDevice(w http.ResponseWriter, r *http.Request) {
	var form routemodels.GetGrowingLevelDeviceRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	res, err := s.Persistence.GetGrowingLevelDevice(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, res)
}

func (s *Server) GetGrowingSpotDevice(w http.ResponseWriter, r *http.Request) {
	var form routemodels.GetGrowingSpotDeviceRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	res, err := s.Persistence.GetGrowingSpotDevice(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, res)
}

func (s *Server) GetAllGrowingGroupDevices(w http.ResponseWriter, r *http.Request) {
	var form routemodels.GetAllGrowingGroupDevicesRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	res, err := s.Persistence.GetAllGrowingGroupDevices(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, res)
}

func (s *Server) GetAllGrowingLocationDevices(w http.ResponseWriter, r *http.Request) {
	var form routemodels.GetAllGrowingLocationDevicesRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	res, err := s.Persistence.GetAllGrowingLocationDevices(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, res)
}

func (s *Server) GetAllGrowingLevelDevices(w http.ResponseWriter, r *http.Request) {
	var form routemodels.GetAllGrowingLevelDevicesRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	res, err := s.Persistence.GetAllGrowingLevelDevices(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, res)
}

func (s *Server) GetAllGrowingSpotDevices(w http.ResponseWriter, r *http.Request) {
	var form routemodels.GetAllGrowingSpotDevicesRequest
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		log.Println(err)
		// send error msg
		return
	}

	res, err := s.Persistence.GetAllGrowingSpotDevices(form)
	if err != nil {
		helpers.SendErrorHeader(w, http.StatusBadRequest, nil)
		return
	}

	helpers.SendSuccessHeader(w, res)
}

package server

import (
	"encoding/json"
	"log"
	"net/http"
	"totally-legit-grow-management/v1/pkg/internal/helpers"
	routemodels "totally-legit-grow-management/v1/pkg/internal/route-models"
)

func (s *Server) CreateGrowSpotPlant(w http.ResponseWriter, r *http.Request) {
	/**********************************************************************
	/
	/	Start Telemetry Span
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Parse Form or Set Vars
	/
	/**********************************************************************/
	var form routemodels.CreateGrowSpotPlantRequest
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
	resp, err := s.Logic.CreateGrowSpotPlant(&form)
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

func (s *Server) DeleteGrowSpotPlant(w http.ResponseWriter, r *http.Request) {
	/**********************************************************************
	/
	/	Start Telemetry Span
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Parse Form or Set Vars
	/
	/**********************************************************************/
	var form routemodels.DeleteGrowSpotPlantRequest
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
	err := s.Logic.DeleteGrowSpotPlant(&form)
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

func (s *Server) EditGrowSpotPlant(w http.ResponseWriter, r *http.Request) {
	/**********************************************************************
	/
	/	Start Telemetry Span
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Parse Form or Set Vars
	/
	/**********************************************************************/
	var form routemodels.EditGrowSpotPlantRequest
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
	resp, err := s.Logic.EditGrowSpotPlant(&form)
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

func (s *Server) GetGrowSpotPlant(w http.ResponseWriter, r *http.Request) {
	/**********************************************************************
	/
	/	Start Telemetry Span
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Parse Form or Set Vars
	/
	/**********************************************************************/
	var form routemodels.GetGrowSpotPlantRequest
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
	resp, err := s.Logic.GetGrowSpotPlant(&form)
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

func (s *Server) GetAllGrowSpotPlants(w http.ResponseWriter, r *http.Request) {
	/**********************************************************************
	/
	/	Start Telemetry Span
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Parse Form or Set Vars
	/
	/**********************************************************************/
	var form routemodels.GetAllGrowSpotPlantsRequest
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
	resp, err := s.Logic.GetAllGrowSpotPlants(&form)
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

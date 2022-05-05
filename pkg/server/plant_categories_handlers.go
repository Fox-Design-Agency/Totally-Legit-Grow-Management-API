package server

import (
	"encoding/json"
	"log"
	"net/http"
	"totally-legit-grow-management/v1/pkg/internal/helpers"
	routemodels "totally-legit-grow-management/v1/pkg/internal/route-models"
)

func (s *Server) CreatePlantCategory(w http.ResponseWriter, r *http.Request) {
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
	var form routemodels.CreatePlantCategoryRequest
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
	resp, err := s.Logic.CreatePlantCategory(&form)
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

func (s *Server) DeletePlantCategory(w http.ResponseWriter, r *http.Request) {
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
	var form routemodels.DeletePlantCategoryRequest
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
	err := s.Logic.DeletePlantCategory(&form)
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

func (s *Server) EditPlantCategory(w http.ResponseWriter, r *http.Request) {
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
	var form routemodels.EditPlantCategoryRequest
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
	resp, err := s.Logic.EditPlantCategory(&form)
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

func (s *Server) GetPlantCategory(w http.ResponseWriter, r *http.Request) {
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
	var form routemodels.GetPlantCategoryRequest
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
	resp, err := s.Logic.GetPlantCategory(&form)
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

func (s *Server) GetAllPlantCategories(w http.ResponseWriter, r *http.Request) {
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
	var form routemodels.GetAllPlantCategoriesRequest
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
	resp, err := s.Logic.GetAllPlantCategories(&form)
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

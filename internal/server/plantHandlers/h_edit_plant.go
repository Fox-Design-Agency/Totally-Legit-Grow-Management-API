// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package planthandlers

import (
	"encoding/json"
	"log"
	"net/http"
	"totally-legit-grow-management/v1/internal/helpers"
	routemodels "totally-legit-grow-management/v1/internal/route-models"
)

// EditPlant will handle the initial request and then send the formed request
// to the logic layer before returning the request
func (s *Plant) EditPlant(w http.ResponseWriter, r *http.Request) {
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
	var form routemodels.EditPlantRequest
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
	resp, err := s.Logic.EditPlant(&form)
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

// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package nutrientcontrol

import routemodels "totally-legit-grow-management/v1/internal/route-models"

// CreateNutrient will validate the incoming request and then call out to the peristence layer to create a Nutrient
func (lg *NutrientControl) CreateNutrient(req *routemodels.CreateNutrientRequest) (*routemodels.CreateNutrientResponse, error) {
	/**********************************************************************
	/
	/	Run Validation Checks
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Start Transaction
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Make Persistence Layer Call(s)
	/
	/**********************************************************************/

	resp, err := lg.Persistence.CreateNutrient(req)
	if err != nil {
		return nil, err
	}

	/**********************************************************************
	/
	/	Commit Transaction
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Form Response and Return
	/
	/**********************************************************************/
	return resp, nil
}

// DeleteNutrient will validate the incoming request and then call out to the peristence layer to
// soft delete a Nutrient
func (lg *NutrientControl) DeleteNutrient(req *routemodels.DeleteNutrientRequest) error {
	/**********************************************************************
	/
	/	Run Validation Checks
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Start Transaction
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Make Persistence Layer Call(s)
	/
	/**********************************************************************/

	err := lg.Persistence.DeleteNutrient(req)
	if err != nil {
		return err
	}

	/**********************************************************************
	/
	/	Commit Transaction
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Form Response and Return
	/
	/**********************************************************************/
	return nil
}

// EditNutrient will validate the request and call out to the persistence layer to edit the specified
// Nutrient (unimplemented)
func (lg *NutrientControl) EditNutrient(req *routemodels.EditNutrientRequest) (*routemodels.EditNutrientResponse, error) {
	/**********************************************************************
	/
	/	Run Validation Checks
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Start Transaction
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Make Persistence Layer Call(s)
	/
	/**********************************************************************/

	resp, err := lg.Persistence.EditNutrient(req)
	if err != nil {
		return nil, err
	}

	/**********************************************************************
	/
	/	Commit Transaction
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Form Response and Return
	/
	/**********************************************************************/
	return resp, nil
}

// GetNutrient will validate the request and call out to the persistence layer to retrieve the specified
// Nutrient
func (lg *NutrientControl) GetNutrient(req *routemodels.GetNutrientRequest) (*routemodels.GetNutrientResponse, error) {
	/**********************************************************************
	/
	/	Run Validation Checks
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Start Transaction
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Make Persistence Layer Call(s)
	/
	/**********************************************************************/

	resp, err := lg.Persistence.GetNutrient(req)
	if err != nil {
		return nil, err
	}

	/**********************************************************************
	/
	/	Commit Transaction
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Form Response and Return
	/
	/**********************************************************************/
	return resp, nil
}

// GetAllNutrients will validate the request and call out to the persistence layer to retrieve all of the
// Nutrients
func (lg *NutrientControl) GetAllNutrients(req *routemodels.GetAllNutrientsRequest) (*routemodels.GetAllNutrientsResponse, error) {
	/**********************************************************************
	/
	/	Run Validation Checks
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Start Transaction
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Make Persistence Layer Call(s)
	/
	/**********************************************************************/

	resp, err := lg.Persistence.GetAllNutrients(req)
	if err != nil {
		return nil, err
	}

	/**********************************************************************
	/
	/	Commit Transaction
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Form Response and Return
	/
	/**********************************************************************/
	return resp, nil
}

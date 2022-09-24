// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package logic

import routemodels "totally-legit-grow-management/v1/internal/route-models"

var _ IPlantLifeCyclesLogic = &Logic{}

// CreatePlantLifeCycle will validate the incoming request and then call out to the peristence layer to create a Plant Life Cycle
func (lg *Logic) CreatePlantLifeCycle(req *routemodels.CreatePlantLifeCycleRequest) (*routemodels.CreatePlantLifeCycleResponse, error) {
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

	resp, err := lg.Persistence.CreatePlantLifeCycle(req)
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

// DeletePlantLifeCycle will validate the incoming request and then call out to the peristence layer to
// soft delete a Plant Lifecycle
func (lg *Logic) DeletePlantLifeCycle(req *routemodels.DeletePlantLifeCycleRequest) error {
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

	err := lg.Persistence.DeletePlantLifeCycle(req)
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

// EditPlantLifeCycle will validate the request and call out to the persistence layer to edit the specified
// Plant Life Cycle (unimplemented)
func (lg *Logic) EditPlantLifeCycle(req *routemodels.EditPlantLifeCycleRequest) (*routemodels.EditPlantLifeCycleResponse, error) {
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

	resp, err := lg.Persistence.EditPlantLifeCycle(req)
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

// GetPlantLifeCycleByID will validate the request and call out to the persistence layer to retrieve the specified
// Plant Life Cycle
func (lg *Logic) GetPlantLifeCycleByID(req *routemodels.GetPlantLifeCycleRequest) (*routemodels.GetPlantLifeCycleResponse, error) {
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

	resp, err := lg.Persistence.GetPlantLifeCycleByID(req)
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

// GetAllPlantLifeCycles will validate the request and call out to the persistence layer to retrieve all of the
// Plant Life Cycles
func (lg *Logic) GetAllPlantLifeCycles(req *routemodels.GetAllPlantLifeCyclesRequest) (*routemodels.GetAllPlantLifeCyclesResponse, error) {
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

	resp, err := lg.Persistence.GetAllPlantLifeCycles(req)
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

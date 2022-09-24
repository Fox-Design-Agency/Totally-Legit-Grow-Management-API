// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package logic

import routemodels "totally-legit-grow-management/v1/internal/route-models"

var _ ISeedsLogic = &Logic{}

// CreateSeed will validate the incoming request and then call out to the peristence layer to create a Seed
func (lg *Logic) CreateSeed(req *routemodels.CreateSeedRequest) (*routemodels.CreateSeedResponse, error) {
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

	resp, err := lg.Persistence.CreateSeed(req)
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

// DeleteSeed will validate the incoming request and then call out to the peristence layer to
// soft delete a Seed
func (lg *Logic) DeleteSeed(req *routemodels.DeleteSeedRequest) error {
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

	err := lg.Persistence.DeleteSeed(req)
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

// EditSeed will validate the request and call out to the persistence layer to edit the specified
// Seed (unimplemented)
func (lg *Logic) EditSeed(req *routemodels.EditSeedRequest) (*routemodels.EditSeedResponse, error) {
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

	resp, err := lg.Persistence.EditSeed(req)
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

// GetSeed will validate the request and call out to the persistence layer to retrieve the specified
// Seed
func (lg *Logic) GetSeed(req *routemodels.GetSeedRequest) (*routemodels.GetSeedResponse, error) {
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

	resp, err := lg.Persistence.GetSeed(req)
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

// GetAllSeeds will validate the request and call out to the persistence layer to retrieve all of the
// Seeds
func (lg *Logic) GetAllSeeds(req *routemodels.GetAllSeedsRequest) (*routemodels.GetAllSeedsResponse, error) {
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

	resp, err := lg.Persistence.GetAllSeeds(req)
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

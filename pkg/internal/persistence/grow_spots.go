// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package persistence

import (
	routemodels "totally-legit-grow-management/v1/pkg/internal/route-models"

	"github.com/jmoiron/sqlx"
)

var _ IGrowSpotsDB = &Persistence{}

// CreateGrowSpot will insert a new Grow Spot into the persistence layer and return the created ID
func (db *Persistence) CreateGrowSpot(req *routemodels.CreateGrowSpotRequest) (*routemodels.CreateGrowSpotResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result string

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/

	args := []interface{}{
		req.DisplayName,
		req.GrowingLevelID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := db.Postgres.QueryRow(CREATE_GROW_SPOT_SQL, args...).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.CreateGrowSpotResponse{
		ID: result,
	}, nil
}

// CreateGrowSpotWithTransaction will insert a new Grow Spot into the persistence layer within a transaction and return the created ID
func (db *Persistence) CreateGrowSpotWithTransaction(tx *sqlx.Tx, req *routemodels.CreateGrowSpotRequest) (*routemodels.CreateGrowSpotResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result string

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/

	args := []interface{}{
		req.DisplayName,
		req.GrowingLevelID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := tx.QueryRow(CREATE_GROW_SPOT_SQL, args...).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.CreateGrowSpotResponse{
		ID: result,
	}, nil
}

// DeleteGrowSpot will set the Growing Spot to be archived
func (db *Persistence) DeleteGrowSpot(req *routemodels.DeleteGrowSpotRequest) error {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/
	args := []interface{}{
		req.ID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if _, err := db.Postgres.Exec(DELETE_GROW_SPOT_SQL, args...); err != nil {
		// handle err
		return err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return nil
}

// EditGrowSpot is not currently implemented
func (db *Persistence) EditGrowSpot(req *routemodels.EditGrowSpotRequest) (*routemodels.EditGrowSpotResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/

	//args := []interface{}{}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return nil, nil
}

// GetGrowSpotByID will get a Grow Spot By ID from the persistence layer
func (db *Persistence) GetGrowSpotByID(req *routemodels.GetGrowSpotRequest) (*routemodels.GetGrowSpotResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.GrowSpot

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/
	args := []interface{}{
		req.ID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := db.Postgres.Get(&result, GET_GROW_SPOT_BY_ID_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetGrowSpotResponse{
		GrowSpot: result,
	}, nil
}

// GetGrowSpotByIDWithTransaction will get a Grow Spot By ID from the persistence layer within a transaction
func (db *Persistence) GetGrowSpotByIDWithTransaction(tx *sqlx.Tx, req *routemodels.GetGrowSpotRequest) (*routemodels.GetGrowSpotResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.GrowSpot

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/
	args := []interface{}{
		req.ID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := tx.Get(&result, GET_GROW_SPOT_BY_ID_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetGrowSpotResponse{
		GrowSpot: result,
	}, nil
}

// GetAllGrowSpotsByGrowLevelID will get all Grow Spot By Grow Level ID from the persistence layer
func (db *Persistence) GetAllGrowSpotsByGrowLevelID(req *routemodels.GetAllGrowSpotsRequest) (*routemodels.GetAllGrowSpotsResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result []routemodels.GrowSpot

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/
	args := []interface{}{}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := db.Postgres.Select(&result, GET_ALL_GROW_SPOTS_BY_GROW_LEVEL_ID_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetAllGrowSpotsResponse{
		GrowSpots: result,
	}, nil
}

// GetAllGrowSpotsByGrowLevelIDWithTransaction will get all Grow Spot By Grow Level ID from the persistence layer within a transaction
func (db *Persistence) GetAllGrowSpotsByGrowLevelIDWithTransaction(tx *sqlx.Tx, req *routemodels.GetAllGrowSpotsRequest) (*routemodels.GetAllGrowSpotsResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result []routemodels.GrowSpot

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/
	args := []interface{}{}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := tx.Select(&result, GET_ALL_GROW_SPOTS_BY_GROW_LEVEL_ID_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetAllGrowSpotsResponse{
		GrowSpots: result,
	}, nil
}

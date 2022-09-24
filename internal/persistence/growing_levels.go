// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package persistence

import (
	routemodels "totally-legit-grow-management/v1/internal/route-models"

	"github.com/jmoiron/sqlx"
)

var _ IGrowingLevelsDB = &Persistence{}

// CreateGrowingLevel will insert a new Growing Level into the persistence layer and return the created ID
func (db *Persistence) CreateGrowingLevel(req *routemodels.CreateGrowingLevelRequest) (*routemodels.CreateGrowingLevelResponse, error) {
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
		req.GrowingLocationID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := db.Postgres.QueryRow(CREATE_GROWING_LEVEL_SQL, args...).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.CreateGrowingLevelResponse{
		ID: result,
	}, nil
}

// CreateGrowingLevelWithTransaction will insert a new Growing Level into the persistence layer
// within a transaction and return the created ID
func (db *Persistence) CreateGrowingLevelWithTransaction(tx *sqlx.Tx, req *routemodels.CreateGrowingLevelRequest) (*routemodels.CreateGrowingLevelResponse, error) {
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
		req.GrowingLocationID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := db.Postgres.QueryRow(CREATE_GROWING_LEVEL_SQL, args...).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.CreateGrowingLevelResponse{
		ID: result,
	}, nil
}

// DeleteGrowingLevel will set the Growing Level to be archived
func (db *Persistence) DeleteGrowingLevel(req *routemodels.DeleteGrowingLevelRequest) error {
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

	if _, err := db.Postgres.Exec(DELETE_GROWING_LEVEL_SQL, args...); err != nil {
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

// EditGrowingLevel is not currently implemented
func (db *Persistence) EditGrowingLevel(req *routemodels.EditGrowingLevelRequest) (*routemodels.EditGrowingLevelResponse, error) {
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

// GetGrowingLevelByID will get a Grow Level By ID from the persistence layer
func (db *Persistence) GetGrowingLevelByID(req *routemodels.GetGrowingLevelRequest) (*routemodels.GetGrowingLevelResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.GrowingLevel

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

	if err := db.Postgres.Get(&result, GET_GROWING_LEVEL_BY_ID_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetGrowingLevelResponse{
		GrowingLevel: result,
	}, nil
}

// GetGrowingLevelByIDWithTransaction will get a Grow Level By ID from the persistence layer within a transaction
func (db *Persistence) GetGrowingLevelByIDWithTransaction(tx *sqlx.Tx, req *routemodels.GetGrowingLevelRequest) (*routemodels.GetGrowingLevelResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.GrowingLevel

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

	if err := db.Postgres.Get(&result, GET_GROWING_LEVEL_BY_ID_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetGrowingLevelResponse{
		GrowingLevel: result,
	}, nil
}

// GetAllGrowingLevelsByGrowingLocationID will get a Grow Level By Growing Location ID from the persistence layer within a transaction
func (db *Persistence) GetAllGrowingLevelsByGrowingLocationID(req *routemodels.GetAllGrowingLevelsRequest) (*routemodels.GetAllGrowingLevelsResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result []routemodels.GrowingLevel

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

	if err := db.Postgres.Select(&result, GET_ALL_GROWING_LEVELS_BY_GROWING_LOCATION_ID_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetAllGrowingLevelsResponse{
		GrowingLevels: result,
	}, nil
}

// GetAllGrowingLevelsByGrowingLocationIDWithTransaction will get a Grow Level By Growing Location ID from the persistence layer
// within a transaction
func (db *Persistence) GetAllGrowingLevelsByGrowingLocationIDWithTransaction(tx *sqlx.Tx, req *routemodels.GetAllGrowingLevelsRequest) (*routemodels.GetAllGrowingLevelsResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result []routemodels.GrowingLevel

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/

	args := []interface{}{
		req.GrowingLevelID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := db.Postgres.Select(&result, GET_ALL_GROWING_LEVELS_BY_GROWING_LOCATION_ID_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetAllGrowingLevelsResponse{
		GrowingLevels: result,
	}, nil
}

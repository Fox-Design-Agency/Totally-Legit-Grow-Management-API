// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package growinglocations

import (
	routemodels "totally-legit-grow-management/v1/internal/route-models"

	"github.com/jmoiron/sqlx"
)

// CreateGrowingLocation will insert a new Growing Location into the persistence layer and return the created ID
func (db *GrowingLocationPersistence) CreateGrowingLocation(req *routemodels.CreateGrowingLocationRequest) (*routemodels.CreateGrowingLocationResponse, error) {
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
		req.GrowingGroupID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := db.Postgres.QueryRow(CREATE_GROWING_LOCATION_SQL, args...).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.CreateGrowingLocationResponse{
		ID: result,
	}, nil
}

// CreateGrowingLocationWithTransaction will insert a new Growing Location into the persistence layer
// within a transaction and return the created ID
func (db *GrowingLocationPersistence) CreateGrowingLocationWithTransaction(tx *sqlx.Tx, req *routemodels.CreateGrowingLocationRequest) (*routemodels.CreateGrowingLocationResponse, error) {
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
		req.GrowingGroupID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := tx.QueryRow(CREATE_GROWING_LOCATION_SQL, args...).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.CreateGrowingLocationResponse{
		ID: result,
	}, nil
}

// DeleteGrowingLocation will set the Growing Location to be archived
func (db *GrowingLocationPersistence) DeleteGrowingLocation(req *routemodels.DeleteGrowingLocationRequest) error {
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

	if _, err := db.Postgres.Exec(DELETE_GROWING_LOCATION_SQL, args...); err != nil {
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

// EditGrowingLocation is not currently implemented
func (db *GrowingLocationPersistence) EditGrowingLocation(req *routemodels.EditGrowingLocationRequest) (*routemodels.EditGrowingLocationResponse, error) {
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

// GetGrowingLocationByID will get a Growing Location By ID from the persistence layer
func (db *GrowingLocationPersistence) GetGrowingLocationByID(req *routemodels.GetGrowingLocationRequest) (*routemodels.GetGrowingLocationResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.GrowingLocation

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

	if err := db.Postgres.Get(&result, GET_GROWING_LOCATION_BY_ID_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetGrowingLocationResponse{
		GrowingLocation: result,
	}, nil
}

// GetGrowingLocationByIDWithTransaction will get a Growing Location By ID from the persistence layer within a transaction
func (db *GrowingLocationPersistence) GetGrowingLocationByIDWithTransaction(tx *sqlx.Tx, req *routemodels.GetGrowingLocationRequest) (*routemodels.GetGrowingLocationResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.GrowingLocation

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

	if err := db.Postgres.Get(&result, GET_GROWING_LOCATION_BY_ID_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetGrowingLocationResponse{
		GrowingLocation: result,
	}, nil
}

// GetAllGrowingLocationsByGrowingGroupID will get all Growing Location By Growing Group ID from the persistence layer
func (db *GrowingLocationPersistence) GetAllGrowingLocationsByGrowingGroupID(req *routemodels.GetAllGrowingLocationsRequest) (*routemodels.GetAllGrowingLocationsResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result []routemodels.GrowingLocation

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/

	args := []interface{}{
		req.GrowingGroupID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := db.Postgres.Select(&result, GET_ALL_GROWING_LOCATIONS_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetAllGrowingLocationsResponse{
		GrowingLocations: result,
	}, nil
}

// GetAllGrowingLocationsByGrowingGroupIDWithTransaction will get all Growing Location By Growing Group ID from the persistence layer
// within a transaction
func (db *GrowingLocationPersistence) GetAllGrowingLocationsByGrowingGroupIDWithTransaction(tx *sqlx.Tx, req *routemodels.GetAllGrowingLocationsRequest) (*routemodels.GetAllGrowingLocationsResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result []routemodels.GrowingLocation

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/

	args := []interface{}{
		req.GrowingGroupID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := db.Postgres.Select(&result, GET_ALL_GROWING_LOCATIONS_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetAllGrowingLocationsResponse{
		GrowingLocations: result,
	}, nil
}

// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package persistence

import (
	routemodels "totally-legit-grow-management/v1/internal/route-models"

	"github.com/jmoiron/sqlx"
)

var _ IPlantLifeCyclesDB = &Persistence{}

// CreatePlantLifeCycle will insert a new Plant Life Cycle into the persistence layer and return the created ID
func (db *Persistence) CreatePlantLifeCycle(req *routemodels.CreatePlantLifeCycleRequest) (*routemodels.CreatePlantLifeCycleResponse, error) {
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
		req.TotalTimeMeasureUnits,
		req.EstTotalTimeMeasure,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := db.Postgres.QueryRow(CREATE_PLANT_LIFE_CYCLE_SQL, args...).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.CreatePlantLifeCycleResponse{
		ID: result,
	}, nil
}

// CreatePlantLifeCycleWithTransaction will insert a new Plant Life Cycle into the persistence layer
// within a transaction and return the created ID
func (db *Persistence) CreatePlantLifeCycleWithTransaction(tx *sqlx.Tx, req *routemodels.CreatePlantLifeCycleRequest) (*routemodels.CreatePlantLifeCycleResponse, error) {
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
		req.TotalTimeMeasureUnits,
		req.EstTotalTimeMeasure,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := tx.QueryRow(CREATE_PLANT_LIFE_CYCLE_SQL, args...).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.CreatePlantLifeCycleResponse{
		ID: result,
	}, nil
}

// DeletePlantLifeCycle will set the Plant Life Cycle to be archived
func (db *Persistence) DeletePlantLifeCycle(req *routemodels.DeletePlantLifeCycleRequest) error {
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
		req.ID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := db.Postgres.QueryRow(DELETE_PLANT_LIFE_CYCLE_SQL, args...).Scan(result); err != nil {
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

// EditPlantLifeCycle is not currently implemented
func (db *Persistence) EditPlantLifeCycle(req *routemodels.EditPlantLifeCycleRequest) (*routemodels.EditPlantLifeCycleResponse, error) {
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

// GetPlantLifeCycleByID will get a Plant Life Cycle By ID from the persistence layer
func (db *Persistence) GetPlantLifeCycleByID(req *routemodels.GetPlantLifeCycleRequest) (*routemodels.GetPlantLifeCycleResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.PlantLifeCycle

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

	if err := db.Postgres.Get(&result, GET_PLANT_LIFE_CYCLE_BY_ID_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetPlantLifeCycleResponse{
		PlantLifeCycle: result,
	}, nil
}

// GetPlantLifeCycleByIDWithTransaction will get a Plant Life Cycle By ID from the persistence layer within a transaction
func (db *Persistence) GetPlantLifeCycleByIDWithTransaction(tx *sqlx.Tx, req *routemodels.GetPlantLifeCycleRequest) (*routemodels.GetPlantLifeCycleResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.PlantLifeCycle

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

	if err := tx.Get(&result, GET_PLANT_LIFE_CYCLE_BY_ID_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetPlantLifeCycleResponse{
		PlantLifeCycle: result,
	}, nil
}

// GetAllPlantLifeCycles will get all Plant Life Cycles from the persistence layer
func (db *Persistence) GetAllPlantLifeCycles(req *routemodels.GetAllPlantLifeCyclesRequest) (*routemodels.GetAllPlantLifeCyclesResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result []routemodels.PlantLifeCycle

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

	if err := db.Postgres.Select(&result, GET_ALL_PLANT_LIFE_CYCLES_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetAllPlantLifeCyclesResponse{
		PlantLifeCycles: result,
	}, nil
}

// GetAllPlantLifeCyclesWithTransaction will get all Plant Life Cycles from the persistence layer within a transaction
func (db *Persistence) GetAllPlantLifeCyclesWithTransaction(tx *sqlx.Tx, req *routemodels.GetAllPlantLifeCyclesRequest) (*routemodels.GetAllPlantLifeCyclesResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result []routemodels.PlantLifeCycle

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

	if err := tx.Select(&result, GET_ALL_PLANT_LIFE_CYCLES_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetAllPlantLifeCyclesResponse{
		PlantLifeCycles: result,
	}, nil
}

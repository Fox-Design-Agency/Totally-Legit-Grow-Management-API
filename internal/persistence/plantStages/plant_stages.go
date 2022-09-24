// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package plantstages

import (
	routemodels "totally-legit-grow-management/v1/internal/route-models"

	"github.com/jmoiron/sqlx"
)

// CreatePlantStage will insert a new Plant Stage into the persistence layer and return the created ID
func (db *PlantStagesPersistence) CreatePlantStage(req *routemodels.CreatePlantStageRequest) (*routemodels.CreatePlantStageResponse, error) {
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
	}
	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := db.Postgres.QueryRow(CREATE_PLANT_STAGE_SQL, args...).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.CreatePlantStageResponse{
		ID: result,
	}, nil
}

// CreatePlantStageWithTransaction will insert a new Plant Stage into the persistence layer within
// a transaction and return the created ID
func (db *PlantStagesPersistence) CreatePlantStageWithTransaction(tx *sqlx.Tx, req *routemodels.CreatePlantStageRequest) (*routemodels.CreatePlantStageResponse, error) {
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
	}
	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := tx.QueryRow(CREATE_PLANT_STAGE_SQL, args...).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.CreatePlantStageResponse{
		ID: result,
	}, nil
}

// CreatePlantStageCare will insert a new Plant Stage Care into the persistence layer and return the created ID
func (db *PlantStagesPersistence) CreatePlantStageCare(req *routemodels.CreatePlantStageCareRequest) (*routemodels.CreatePlantStageCareResponse, error) {
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
	}
	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := db.Postgres.QueryRow(CREATE_PLANT_STAGE_CARE_SQL, args...).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.CreatePlantStageCareResponse{
		ID: result,
	}, nil
}

// CreatePlantStageCareWithTransaction will insert a new Plant Stage Care into the persistence layer within
// a transaction and return the created ID
func (db *PlantStagesPersistence) CreatePlantStageCareWithTransaction(tx *sqlx.Tx, req *routemodels.CreatePlantStageCareRequest) (*routemodels.CreatePlantStageCareResponse, error) {
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
	}
	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := tx.QueryRow(CREATE_PLANT_STAGE_CARE_SQL, args...).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.CreatePlantStageCareResponse{
		ID: result,
	}, nil
}

// CreatePlantStageNutrients will insert a new Plant Stage Nutrients into the persistence layer and return the created ID
func (db *PlantStagesPersistence) CreatePlantStageNutrients(req *routemodels.CreatePlantStageNutrientsRequest) (*routemodels.CreatePlantStageNutrientsResponse, error) {
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
	}
	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := db.Postgres.QueryRow(CREATE_PLANT_STAGE_NUTRIENTS_SQL, args...).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.CreatePlantStageNutrientsResponse{
		ID: result,
	}, nil
}

// CreatePlantStageNutrientsWithTransaction will insert a new Plant Stage Nutrients into the persistence layer within
// a transaction and return the created ID
func (db *PlantStagesPersistence) CreatePlantStageNutrientsWithTransaction(tx *sqlx.Tx, req *routemodels.CreatePlantStageNutrientsRequest) (*routemodels.CreatePlantStageNutrientsResponse, error) {
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
	}
	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := tx.QueryRow(CREATE_PLANT_STAGE_NUTRIENTS_SQL, args...).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.CreatePlantStageNutrientsResponse{
		ID: result,
	}, nil
}

// ConnectPlantStage will insert a new Plant Stage into the persistence layer and return the created ID
func (db *PlantStagesPersistence) ConnectPlantStage(req *routemodels.ConnectPlantStageRequest) (*routemodels.ConnectPlantStageResponse, error) {
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
	}
	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := db.Postgres.QueryRow(CONNECT_PLANT_STAGE_SQL, args...).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.ConnectPlantStageResponse{
		ID: result,
	}, nil
}

// ConnectPlantStage will insert a new Plant Stage into the persistence layer within a
// transaction and return the created ID
func (db *PlantStagesPersistence) ConnectPlantStageWithTransaction(tx *sqlx.Tx, req *routemodels.ConnectPlantStageRequest) (*routemodels.ConnectPlantStageResponse, error) {
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
	}
	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := tx.QueryRow(CONNECT_PLANT_STAGE_SQL, args...).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.ConnectPlantStageResponse{
		ID: result,
	}, nil
}

// DeletePlantStage will set the Plant Stage to be archived
func (db *PlantStagesPersistence) DeletePlantStage(req *routemodels.DeletePlantStageRequest) error {
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

	if _, err := db.Postgres.Exec(DELETE_PLANT_STAGE_SQL, args...); err != nil {
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

// EditPlantStage is not currently implemented
func (db *PlantStagesPersistence) EditPlantStage(req *routemodels.EditPlantStageRequest) (*routemodels.EditPlantStageResponse, error) {
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

// GetPlantStageByID will get a Plant Stage By ID from the persistence layer
func (db *PlantStagesPersistence) GetPlantStageByID(req *routemodels.GetPlantStageRequest) (*routemodels.GetPlantStageResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.PlantStage

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

	if err := db.Postgres.Get(&result, GET_PLANT_STAGE_BY_ID_SQL, args); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetPlantStageResponse{
		PlantStage: result,
	}, nil
}

// GetPlantStageByIDWithTransaction will get a Plant Stage By ID from the persistence layer within a transaction
func (db *PlantStagesPersistence) GetPlantStageByIDWithTransaction(tx *sqlx.Tx, req *routemodels.GetPlantStageRequest) (*routemodels.GetPlantStageResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.PlantStage

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

	if err := tx.Get(&result, GET_PLANT_STAGE_BY_ID_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetPlantStageResponse{
		PlantStage: result,
	}, nil
}

// GetPlantStageCareByID will get a Plant Care By ID from the persistence layer
func (db *PlantStagesPersistence) GetPlantStageCareByID(req *routemodels.GetPlantStageCareByIDRequest) (*routemodels.GetPlantStageCareByIDResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.PlantStage

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

	if err := db.Postgres.Get(&result, GET_PLANT_STAGE_CARE_BY_ID_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetPlantStageCareByIDResponse{
		PlantStage: result,
	}, nil
}

// GetPlantStageCareByPlantStageIDWithTransaction will get a Plant Care By ID from the persistence layer within a transaction
func (db *PlantStagesPersistence) GetPlantStageCareByPlantStageIDWithTransaction(tx *sqlx.Tx, req *routemodels.GetPlantStageCareByPlantStageIDRequest) (*routemodels.GetPlantStageCareByPlantStageIDResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.PlantStage

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/
	args := []interface{}{
		req.OrganizationID,
	}
	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := tx.Get(&result, GET_PLANT_STAGE_CARE_BY_ID_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetPlantStageCareByPlantStageIDResponse{
		PlantStage: result,
	}, nil
}

// GetPlantStageNutrient will get a Plant Nutrient By ID from the persistence layer
func (db *PlantStagesPersistence) GetPlantStageNutrient(req *routemodels.GetPlantStageNutrientRequest) (*routemodels.GetPlantStageNutrientResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.PlantStage

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/
	args := []interface{}{
		req.OrganizationID,
	}
	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := db.Postgres.Get(&result, GET_PLANT_STAGE_NUTRIENT_BY_ID_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetPlantStageNutrientResponse{
		PlantStage: result,
	}, nil
}

// GetAllPlantStages will get all Plant Stages from the persistence layer
func (db *PlantStagesPersistence) GetAllPlantStages(req *routemodels.GetAllPlantStagesRequest) (*routemodels.GetAllPlantStagesResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result []routemodels.PlantStage
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

	if err := db.Postgres.Select(&result, GET_ALL_PLANT_STAGE_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetAllPlantStagesResponse{
		PlantStages: result,
	}, nil
}

// GetAllPlantStagesWithTransaction will get all Plant Stages from the persistence layer within a transaction
func (db *PlantStagesPersistence) GetAllPlantStagesWithTransaction(tx *sqlx.Tx, req *routemodels.GetAllPlantStagesRequest) (*routemodels.GetAllPlantStagesResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result []routemodels.PlantStage

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

	if err := tx.Select(&result, GET_ALL_PLANT_STAGE_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetAllPlantStagesResponse{
		PlantStages: result,
	}, nil
}

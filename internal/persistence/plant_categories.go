// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package persistence

import routemodels "totally-legit-grow-management/v1/internal/route-models"

var _ IPlantCategoriesDB = &Persistence{}

// CreatePlantCategory will insert a new Plant Category into the persistence layer and return the created ID
func (db *Persistence) CreatePlantCategory(req *routemodels.CreatePlantCategoryRequest) (*routemodels.CreatePlantCategoryResponse, error) {
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

	if err := db.Postgres.QueryRow(CREATE_PLANT_CATEGORY_SQL, args...).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.CreatePlantCategoryResponse{
		ID: result,
	}, nil
}

// DeletePlantCategory will set the Plant Category to be archived
func (db *Persistence) DeletePlantCategory(req *routemodels.DeletePlantCategoryRequest) error {
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

	if err := db.Postgres.QueryRow(DELETE_PLANT_CATEGORY_SQL, args...).Scan(result); err != nil {
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

// EditPlantCategory is not currently implemented
func (db *Persistence) EditPlantCategory(req *routemodels.EditPlantCategoryRequest) (*routemodels.EditPlantCategoryResponse, error) {
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

// GetPlantCategory will get a Plant Category By ID from the persistence layer
func (db *Persistence) GetPlantCategory(req *routemodels.GetPlantCategoryRequest) (*routemodels.GetPlantCategoryResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.PlantCategory

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

	if err := db.Postgres.Get(&result, GET_PLANT_CATEGORY_BY_ID_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetPlantCategoryResponse{
		PlantCategory: result,
	}, nil
}

// GetAllPlantCategories will get all Plant Categories from the persistence layer
func (db *Persistence) GetAllPlantCategories(req *routemodels.GetAllPlantCategoriesRequest) (*routemodels.GetAllPlantCategoriesResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result []routemodels.PlantCategory

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

	if err := db.Postgres.Select(&result, GET_ALL_PLANT_CATEGORY_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetAllPlantCategoriesResponse{
		PlantCategories: result,
	}, nil
}

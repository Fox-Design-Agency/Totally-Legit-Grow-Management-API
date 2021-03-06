// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package logic

import routemodels "totally-legit-grow-management/v1/pkg/internal/route-models"

var _ IPlantCategoriesLogic = &Logic{}

// CreatePlantCategory will validate the incoming request and then call out to the peristence layer to create a Plant Category
func (lg *Logic) CreatePlantCategory(req *routemodels.CreatePlantCategoryRequest) (*routemodels.CreatePlantCategoryResponse, error) {
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

	resp, err := lg.Persistence.CreatePlantCategory(req)
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

// DeletePlantCategory will validate the incoming request and then call out to the peristence layer to
// soft delete a Plant Category
func (lg *Logic) DeletePlantCategory(req *routemodels.DeletePlantCategoryRequest) error {
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

	err := lg.Persistence.DeletePlantCategory(req)
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

// EditPlantCategory will validate the request and call out to the persistence layer to edit the specified
// Plant Category (unimplemented)
func (lg *Logic) EditPlantCategory(req *routemodels.EditPlantCategoryRequest) (*routemodels.EditPlantCategoryResponse, error) {
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

	resp, err := lg.Persistence.EditPlantCategory(req)
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

// GetPlantCategory will validate the request and call out to the persistence layer to retrieve the specified
// Plant Category
func (lg *Logic) GetPlantCategory(req *routemodels.GetPlantCategoryRequest) (*routemodels.GetPlantCategoryResponse, error) {
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

	resp, err := lg.Persistence.GetPlantCategory(req)
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

// GetAllPlantCategories will validate the request and call out to the persistence layer to retrieve all fo the
// Plant Categories
func (lg *Logic) GetAllPlantCategories(req *routemodels.GetAllPlantCategoriesRequest) (*routemodels.GetAllPlantCategoriesResponse, error) {
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

	resp, err := lg.Persistence.GetAllPlantCategories(req)
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

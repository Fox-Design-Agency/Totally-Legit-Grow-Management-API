// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package growingmedium

import routemodels "totally-legit-grow-management/v1/internal/route-models"

// CreateGrowingLocation will insert a new Growing Medium into the persistence layer and return the created ID
func (db *GrowingMediumPersistence) CreateGrowingMedium(req *routemodels.CreateGrowingMediumRequest) (*routemodels.CreateGrowingMediumResponse, error) {
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

	if err := db.Postgres.QueryRow(CREATE_GROWING_MEDIUM_SQL, args...).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.CreateGrowingMediumResponse{
		ID: result,
	}, nil
}

// DeleteGrowingMedium will set the Growing Medium to be archived
func (db *GrowingMediumPersistence) DeleteGrowingMedium(req *routemodels.DeleteGrowingMediumRequest) error {
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

	if _, err := db.Postgres.Exec(DELETE_GROWING_MEDIUM_SQL, args...); err != nil {
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

// EditGrowingMedium is not currently implemented
func (db *GrowingMediumPersistence) EditGrowingMedium(req *routemodels.EditGrowingMediumRequest) (*routemodels.EditGrowingMediumResponse, error) {
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

// GetGrowingMedium will get a Growing Medium By ID from the persistence layer
func (db *GrowingMediumPersistence) GetGrowingMedium(req *routemodels.GetGrowingMediumRequest) (*routemodels.GetGrowingMediumResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.GrowingMedium

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

	if err := db.Postgres.Get(&result, GET_GROWING_MEDIUM_BY_ID_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetGrowingMediumResponse{
		GrowingMedium: result,
	}, nil
}

// GetAllGrowingMediums will get all Growing Mediums from the persistence layer
func (db *GrowingMediumPersistence) GetAllGrowingMediums(req *routemodels.GetAllGrowingMediumsRequest) (*routemodels.GetAllGrowingMediumsResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result []routemodels.GrowingMedium

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

	if err := db.Postgres.Select(&result, GET_ALL_GROWING_MEDIUM_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetAllGrowingMediumsResponse{
		GrowingMediums: result,
	}, nil
}

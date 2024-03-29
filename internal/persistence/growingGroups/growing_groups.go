// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package growinggroups

import (
	routemodels "totally-legit-grow-management/v1/internal/route-models"

	"github.com/jmoiron/sqlx"
)

// CreateGrowingGroup will insert a new Growing Group into the persistence layer and return the created ID
func (db *GrowingGroupPersistence) CreateGrowingGroup(req *routemodels.CreateGrowingGroupRequest) (*routemodels.CreateGrowingGroupResponse, error) {
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
		req.OrganizationID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := db.Postgres.QueryRow(CREATE_GROWING_GROUPS_SQL, args...).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.CreateGrowingGroupResponse{
		ID: result,
	}, nil
}

// CreateGrowingGroupWithTransaction will insert a new Growing Group into the persistence layer
// within a transaction and return the created ID
func (db *GrowingGroupPersistence) CreateGrowingGroupWithTransaction(tx *sqlx.Tx, req *routemodels.CreateGrowingGroupRequest) (*routemodels.CreateGrowingGroupResponse, error) {
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
		req.OrganizationID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := tx.QueryRow(CREATE_GROWING_GROUPS_SQL, args...).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.CreateGrowingGroupResponse{
		ID: result,
	}, nil
}

// DeleteGrowingGroup will set the Growing Group to be archived
func (db *GrowingGroupPersistence) DeleteGrowingGroup(req *routemodels.DeleteGrowingGroupRequest) error {
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

	if _, err := db.Postgres.Exec(DELETE_GROWING_GROUP_SQL, args...); err != nil {
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

// EditGrowingGroup is not currently implemented
func (db *GrowingGroupPersistence) EditGrowingGroup(req *routemodels.EditGrowingGroupRequest) (*routemodels.EditGrowingGroupResponse, error) {
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

// GetGrowingGroupByID will get a Grow Group By ID from the persistence layer
func (db *GrowingGroupPersistence) GetGrowingGroupByID(req *routemodels.GetGrowingGroupRequest) (*routemodels.GetGrowingGroupResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.GrowingGroup

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

	if err := db.Postgres.Get(&result, GET_GROWING_GROUP_BY_ID_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetGrowingGroupResponse{
		GrowingGroup: result,
	}, nil
}

// GetGrowingGroupByIDWithTransaction will get a Grow Group By ID from the persistence layer within a transaction
func (db *GrowingGroupPersistence) GetGrowingGroupByIDWithTransaction(tx *sqlx.Tx, req *routemodels.GetGrowingGroupRequest) (*routemodels.GetGrowingGroupResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.GrowingGroup

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

	if err := tx.Get(&result, GET_GROWING_GROUP_BY_ID_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetGrowingGroupResponse{
		GrowingGroup: result,
	}, nil
}

// GetAllGrowingGroupsByOrganizationID will get all Grow Groups By Organization ID from the persistence layer
func (db *GrowingGroupPersistence) GetAllGrowingGroupsByOrganizationID(req *routemodels.GetAllGrowingGroupsRequest) (*routemodels.GetAllGrowingGroupsResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result []routemodels.GrowingGroup

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

	if err := db.Postgres.Select(&result, GET_ALL_GROWING_GROUPS_BY_ORGANIZATION_ID, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetAllGrowingGroupsResponse{
		Groups: result,
	}, nil
}

// GetAllGrowingGroupsByOrganizationIDWithTransaction will get all Grow Groups By Organization ID from the persistence layer
// within a transaction
func (db *GrowingGroupPersistence) GetAllGrowingGroupsByOrganizationIDWithTransaction(tx *sqlx.Tx, req *routemodels.GetAllGrowingGroupsRequest) (*routemodels.GetAllGrowingGroupsResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result []routemodels.GrowingGroup

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

	if err := tx.Select(&result, GET_ALL_GROWING_GROUPS_BY_ORGANIZATION_ID, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetAllGrowingGroupsResponse{
		Groups: result,
	}, nil
}

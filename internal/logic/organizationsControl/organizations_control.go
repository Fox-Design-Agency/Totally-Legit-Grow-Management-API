// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package organizationcontrol

import routemodels "totally-legit-grow-management/v1/internal/route-models"

// CreateOrganization will validate the incoming request and then call out to the peristence layer to create an Organization
func (lg *OrganizationControl) CreateOrganization(req *routemodels.CreateOrganizationRequest) (*routemodels.CreateOrganizationResponse, error) {
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

	resp, err := lg.Persistence.CreateOrganization(req)
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

// DeleteOrganization will validate the incoming request and then call out to the peristence layer to
// soft delete a Organization
func (lg *OrganizationControl) DeleteOrganization(req *routemodels.DeleteOrganizationRequest) error {
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

	err := lg.Persistence.DeleteOrganization(req)
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

// EditOrganization will validate the request and call out to the persistence layer to edit the specified
// Organization (unimplemented)
func (lg *OrganizationControl) EditOrganization(req *routemodels.EditOrganizationRequest) (*routemodels.EditOrganizationResponse, error) {
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

	resp, err := lg.Persistence.EditOrganization(req)
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

// GetOrganization will validate the request and call out to the persistence layer to retrieve the specified
// Organization
func (lg *OrganizationControl) GetOrganization(req *routemodels.GetOrganizationRequest) (*routemodels.GetOrganizationResponse, error) {
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

	resp, err := lg.Persistence.GetOrganization(req)
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

// GetAllOrganizations will validate the request and call out to the persistence layer to retrieve all fo the
// Organizations
func (lg *OrganizationControl) GetAllOrganizations(req *routemodels.GetAllOrganizationsRequest) (*routemodels.GetAllOrganizationsResponse, error) {
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

	resp, err := lg.Persistence.GetAllOrganizations(req)
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

package persistence

import routemodels "totally-legit-grow-management/v1/pkg/internal/route-models"

var _ IOrganizationsDB = &Persistence{}

func (db *Persistence) CreateOrganization(req *routemodels.CreateOrganizationRequest) (*routemodels.CreateOrganizationResponse, error) {
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

	if err := db.Postgres.QueryRow(CREATE_ORGANIZATION_SQL, args...).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.CreateOrganizationResponse{
		ID: result,
	}, nil
}

func (db *Persistence) DeleteOrganization(req *routemodels.DeleteOrganizationRequest) error {
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

	if err := db.Postgres.QueryRow(DELETE_ORGANIZATION_SQL, args...).Scan(result); err != nil {
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

func (db *Persistence) EditOrganization(req *routemodels.EditOrganizationRequest) (*routemodels.EditOrganizationResponse, error) {
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

func (db *Persistence) GetOrganization(req *routemodels.GetOrganizationRequest) (*routemodels.GetOrganizationResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.Organization

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

	if err := db.Postgres.Get(&result, GET_ORGANIZATION_BY_ID_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetOrganizationResponse{
		Organization: result,
	}, nil
}

func (db *Persistence) GetAllOrganizations(req *routemodels.GetAllOrganizationsRequest) (*routemodels.GetAllOrganizationsResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result []routemodels.Organization

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

	if err := db.Postgres.Select(&result, GET_ALL_ORGANIZATIONS_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetAllOrganizationsResponse{
		Organizations: result,
	}, nil
}

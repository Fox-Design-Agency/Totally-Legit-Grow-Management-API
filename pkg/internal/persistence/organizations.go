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

	SQL := `
	INSERT INTO organizations
	(display_name)
	VALUES ($1)
	RETURNING id
	`
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

	if err := db.Postgres.QueryRow(SQL, args...).Scan(result); err != nil {
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

	SQL := `
	UPDATE organizations SET archived = NOW()
	WHERE organizations.id = $1
	`
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

	if err := db.Postgres.QueryRow(SQL, args...).Scan(result); err != nil {
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

	SQL := `
	SELECT
		organizations.id AS id,
		organizations.display_name AS display_name,
		organizations.created AS created_at,
		organizations.updated AS updated_at,
		cre_member.id AS created_member_id,
		cre_member.display_name AS created_member_name,
		up_member.id AS updated_member_id,
		up_member.display_name AS updated_member_name
	FROM organizations
	LEFT JOIN members AS cre_member ON members.id = organizations.created_by
	LEFT JOIN members AS up_member ON members.id = organizations.updated_by
	WHERE organizations.id = $1
	`
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

	if err := db.Postgres.Get(&result, SQL, args...); err != nil {
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

	SQL := `
	SELECT
		organizations.id AS id,
		organizations.display_name AS display_name,
		organizations.created AS created_at,
		organizations.updated AS updated_at,
		cre_member.id AS created_member_id,
		cre_member.display_name AS created_member_name,
		up_member.id AS updated_member_id,
		up_member.display_name AS updated_member_name
	FROM organizations
	LEFT JOIN members AS cre_member ON members.id = organizations.created_by
	LEFT JOIN members AS up_member ON members.id = organizations.updated_by
	`
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

	if err := db.Postgres.Select(&result, SQL, args...); err != nil {
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

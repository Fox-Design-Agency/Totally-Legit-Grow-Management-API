package persistence

import routemodels "totally-legit-grow-management/v1/pkg/internal/route-models"

var _ IOrganizationsDB = &Persistence{}

func (db *Persistence) CreateOrganization(req *routemodels.CreateOrganizationRequest) (*routemodels.CreateOrganizationResponse, error) {
	var result string

	SQL := `
	INSERT INTO organizations
	(display_name)
	VALUES ($1)
	RETURNING id
	`

	// Make the appropiate SQL Call
	if err := db.Postgres.QueryRow(SQL, req.DisplayName).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.CreateOrganizationResponse{
		ID: result,
	}, nil
}

func (db *Persistence) DeleteOrganization(req *routemodels.DeleteOrganizationRequest) error {
	var result string

	SQL := `
	UPDATE organizations SET archived = NOW()
	WHERE organizations.id = $1
	`

	// Make the appropiate SQL Call
	if err := db.Postgres.QueryRow(SQL, req.ID).Scan(result); err != nil {
		// handle err
		return err
	}

	return nil
}

func (db *Persistence) EditOrganization(req *routemodels.EditOrganizationRequest) (*routemodels.EditOrganizationResponse, error) {
	return nil, nil
}

func (db *Persistence) GetOrganization(req *routemodels.GetOrganizationRequest) (*routemodels.GetOrganizationResponse, error) {
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

	// Make the appropiate SQL Call
	if err := db.Postgres.Get(&result, SQL, req.ID); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetOrganizationResponse{
		Organization: result,
	}, nil
}

func (db *Persistence) GetAllOrganizations(req *routemodels.GetAllOrganizationsRequest) (*routemodels.GetAllOrganizationsResponse, error) {
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

	// Make the appropiate SQL Call
	if err := db.Postgres.Select(&result, SQL); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetAllOrganizationsResponse{
		Organizations: result,
	}, nil
}

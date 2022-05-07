package persistence

import (
	routemodels "totally-legit-grow-management/v1/pkg/internal/route-models"

	"github.com/jmoiron/sqlx"
)

var _ IGrowingGroupsDB = &Persistence{}

func (db *Persistence) CreateGrowingGroup(req *routemodels.CreateGrowingGroupRequest) (*routemodels.CreateGrowingGroupResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result string

	SQL := `
	INSERT INTO growing_groups
	(display_name, organization)
	VALUES ($1, $2)
	RETURNING id
	`

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

	if err := db.Postgres.QueryRow(SQL, args...).Scan(result); err != nil {
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

func (db *Persistence) CreateGrowingGroupWithTransaction(tx *sqlx.Tx, req *routemodels.CreateGrowingGroupRequest) (*routemodels.CreateGrowingGroupResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result string

	SQL := `
	INSERT INTO growing_groups
	(display_name, organization)
	VALUES ($1, $2)
	RETURNING id
	`

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

	if err := tx.QueryRow(SQL, args...).Scan(result); err != nil {
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

func (db *Persistence) DeleteGrowingGroup(req *routemodels.DeleteGrowingGroupRequest) error {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result string

	SQL := `
	UPDATE growing_groups SET archived = NOW()
	WHERE growing_groups.id = $1
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

func (db *Persistence) EditGrowingGroup(req *routemodels.EditGrowingGroupRequest) (*routemodels.EditGrowingGroupResponse, error) {
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

func (db *Persistence) GetGrowingGroupByID(req *routemodels.GetGrowingGroupRequest) (*routemodels.GetGrowingGroupResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.GrowingGroup

	SQL := `
	SELECT
	growing_groups.id AS id,
	growing_groups.organization AS organization_id,
	growing_groups.display_name AS display_name,
	growing_groups.created AS created_at,
	growing_groups.updated AS updated_at,
		cre_member.id AS created_member_id,
		cre_member.display_name AS created_member_name,
		up_member.id AS updated_member_id,
		up_member.display_name AS updated_member_name
	FROM growing_groups
	LEFT JOIN members AS cre_member ON members.id = growing_groups.created_by
	LEFT JOIN members AS up_member ON members.id = growing_groups.updated_by
	WHERE growing_groups.id = $1
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
	return &routemodels.GetGrowingGroupResponse{
		GrowingGroup: result,
	}, nil
}

func (db *Persistence) GetGrowingGroupByIDWithTransaction(tx *sqlx.Tx, req *routemodels.GetGrowingGroupRequest) (*routemodels.GetGrowingGroupResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.GrowingGroup

	SQL := `
	SELECT
	growing_groups.id AS id,
	growing_groups.organization AS organization_id,
	growing_groups.display_name AS display_name,
	growing_groups.created AS created_at,
	growing_groups.updated AS updated_at,
		cre_member.id AS created_member_id,
		cre_member.display_name AS created_member_name,
		up_member.id AS updated_member_id,
		up_member.display_name AS updated_member_name
	FROM growing_groups
	LEFT JOIN members AS cre_member ON members.id = growing_groups.created_by
	LEFT JOIN members AS up_member ON members.id = growing_groups.updated_by
	WHERE growing_groups.id = $1
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

	if err := tx.Get(&result, SQL, args...); err != nil {
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

func (db *Persistence) GetAllGrowingGroupsByOrganizationID(req *routemodels.GetAllGrowingGroupsRequest) (*routemodels.GetAllGrowingGroupsResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result []routemodels.GrowingGroup

	SQL := `
	SELECT
		growing_groups.id AS id,
		growing_groups.organization AS organization_id,
		growing_groups.display_name AS display_name,
		growing_groups.created AS created_at,
		growing_groups.updated AS updated_at,
			cre_member.id AS created_member_id,
			cre_member.display_name AS created_member_name,
			up_member.id AS updated_member_id,
			up_member.display_name AS updated_member_name
	FROM growing_groups
	LEFT JOIN members AS cre_member ON members.id = growing_groups.created_by
	LEFT JOIN members AS up_member ON members.id = growing_groups.updated_by
	WHERE growing_groups.organization = $1
	`

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

	if err := db.Postgres.Select(&result, SQL, args...); err != nil {
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

func (db *Persistence) GetAllGrowingGroupsByOrganizationIDWithTransaction(tx *sqlx.Tx, req *routemodels.GetAllGrowingGroupsRequest) (*routemodels.GetAllGrowingGroupsResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result []routemodels.GrowingGroup

	SQL := `
	SELECT
		growing_groups.id AS id,
		growing_groups.organization AS organization_id,
		growing_groups.display_name AS display_name,
		growing_groups.created AS created_at,
		growing_groups.updated AS updated_at,
			cre_member.id AS created_member_id,
			cre_member.display_name AS created_member_name,
			up_member.id AS updated_member_id,
			up_member.display_name AS updated_member_name
	FROM growing_groups
	LEFT JOIN members AS cre_member ON members.id = growing_groups.created_by
	LEFT JOIN members AS up_member ON members.id = growing_groups.updated_by
	WHERE growing_groups.organization = $1
	`

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

	if err := tx.Select(&result, SQL, args...); err != nil {
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

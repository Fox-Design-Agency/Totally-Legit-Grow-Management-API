package persistence

import routemodels "totally-legit-grow-management/v1/pkg/internal/route-models"

var _ IMembersDB = &Persistence{}

func (db *Persistence) CreateMember(req *routemodels.CreateMemberRequest) (*routemodels.CreateMemberResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result string

	SQL := `
	INSERT INTO members
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
	return &routemodels.CreateMemberResponse{
		ID: result,
	}, nil
}

func (db *Persistence) DeleteMember(req *routemodels.DeleteMemberRequest) error {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result string

	SQL := `
	UPDATE members SET archived = NOW()
	WHERE members.id = $1
	`

	// Make the appropiate SQL Call
	if err := db.Postgres.QueryRow(SQL, req.ID).Scan(result); err != nil {
		// handle err
		return err
	}

	return nil
}

func (db *Persistence) EditMember(req *routemodels.EditMemberRequest) (*routemodels.EditMemberResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	return nil, nil
}

func (db *Persistence) GetMember(req *routemodels.GetMemberRequest) (*routemodels.GetMemberResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.Member

	SQL := `
	SELECT
		members.id AS id,
		members.display_name AS display_name,
		members.created AS created_at,
		members.updated AS updated_at,
		cre_member.id AS created_member_id,
		cre_member.display_name AS created_member_name,
		up_member.id AS updated_member_id,
		up_member.display_name AS updated_member_name
	FROM members
	LEFT JOIN members AS cre_member ON members.id = members.created_by
	LEFT JOIN members AS up_member ON members.id = members.updated_by
	WHERE members.id = $1
	`

	// Make the appropiate SQL Call
	if err := db.Postgres.Get(&result, SQL, req.ID); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetMemberResponse{
		Member: result,
	}, nil
}

func (db *Persistence) GetAllMembers(req *routemodels.GetAllMembersRequest) (*routemodels.GetAllMembersResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result []routemodels.Member

	SQL := `
	SELECT
		members.id AS id,
		members.display_name AS display_name,
		members.created AS created_at,
		members.updated AS updated_at,
		cre_member.id AS created_member_id,
		cre_member.display_name AS created_member_name,
		up_member.id AS updated_member_id,
		up_member.display_name AS updated_member_name
	FROM members
	LEFT JOIN members AS cre_member ON members.id = members.created_by
	LEFT JOIN members AS up_member ON members.id = members.updated_by
	WHERE members.id = $1
	`

	// Make the appropiate SQL Call
	if err := db.Postgres.Select(&result, SQL); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetAllMembersResponse{
		Members: result,
	}, nil
}

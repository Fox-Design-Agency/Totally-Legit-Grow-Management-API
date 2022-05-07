package persistence

import (
	routemodels "totally-legit-grow-management/v1/pkg/internal/route-models"

	"github.com/jmoiron/sqlx"
)

var _ IGrowSpotsDB = &Persistence{}

func (db *Persistence) CreateGrowSpot(req *routemodels.CreateGrowSpotRequest) (*routemodels.CreateGrowSpotResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result string

	SQL := `
	INSERT INTO grow_spots
	(display_name, growing_level)
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
		req.GrowingLevelID,
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
	return &routemodels.CreateGrowSpotResponse{
		ID: result,
	}, nil
}

func (db *Persistence) CreateGrowSpotWithTransaction(tx *sqlx.Tx, req *routemodels.CreateGrowSpotRequest) (*routemodels.CreateGrowSpotResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result string

	SQL := `
	INSERT INTO grow_spots
	(display_name, growing_level)
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
		req.GrowingLevelID,
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
	return &routemodels.CreateGrowSpotResponse{
		ID: result,
	}, nil
}

func (db *Persistence) DeleteGrowSpot(req *routemodels.DeleteGrowSpotRequest) error {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result string

	SQL := `
	UPDATE grow_spots SET archived = NOW()
	WHERE grow_spots.id = $1
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

func (db *Persistence) EditGrowSpot(req *routemodels.EditGrowSpotRequest) (*routemodels.EditGrowSpotResponse, error) {
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

func (db *Persistence) GetGrowSpotByID(req *routemodels.GetGrowSpotRequest) (*routemodels.GetGrowSpotResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.GrowSpot

	SQL := `
	SELECT
		grow_spots.id AS id,
		grow_spots.growing_level AS growing_level_id,
		grow_spots.display_name AS display_name,
		grow_spots.created AS created_at,
		grow_spots.updated AS updated_at,
		cre_member.id AS created_member_id,
		cre_member.display_name AS created_member_name,
		up_member.id AS updated_member_id,
		up_member.display_name AS updated_member_name
	FROM grow_spots
	LEFT JOIN members AS cre_member ON members.id = grow_spots.created_by
	LEFT JOIN members AS up_member ON members.id = grow_spots.updated_by
	WHERE grow_spots.id = $1
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
	return &routemodels.GetGrowSpotResponse{
		GrowSpot: result,
	}, nil
}

func (db *Persistence) GetGrowSpotByIDWithTransaction(tx *sqlx.Tx, req *routemodels.GetGrowSpotRequest) (*routemodels.GetGrowSpotResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.GrowSpot

	SQL := `
	SELECT
		grow_spots.id AS id,
		grow_spots.growing_level AS growing_level_id,
		grow_spots.display_name AS display_name,
		grow_spots.created AS created_at,
		grow_spots.updated AS updated_at,
		cre_member.id AS created_member_id,
		cre_member.display_name AS created_member_name,
		up_member.id AS updated_member_id,
		up_member.display_name AS updated_member_name
	FROM grow_spots
	LEFT JOIN members AS cre_member ON members.id = grow_spots.created_by
	LEFT JOIN members AS up_member ON members.id = grow_spots.updated_by
	WHERE grow_spots.id = $1
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
	return &routemodels.GetGrowSpotResponse{
		GrowSpot: result,
	}, nil
}

func (db *Persistence) GetAllGrowSpotsByGrowLevelID(req *routemodels.GetAllGrowSpotsRequest) (*routemodels.GetAllGrowSpotsResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result []routemodels.GrowSpot

	SQL := `
	SELECT
		grow_spots.id AS id,
		grow_spots.growing_level AS growing_level_id,
		grow_spots.display_name AS display_name,
		grow_spots.created AS created_at,
		grow_spots.updated AS updated_at,
		cre_member.id AS created_member_id,
		cre_member.display_name AS created_member_name,
		up_member.id AS updated_member_id,
		up_member.display_name AS updated_member_name
	FROM grow_spots
	LEFT JOIN members AS cre_member ON members.id = grow_spots.created_by
	LEFT JOIN members AS up_member ON members.id = grow_spots.updated_by
	WHERE grow_spots.growing_level = $1
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
	return &routemodels.GetAllGrowSpotsResponse{
		GrowSpots: result,
	}, nil
}

func (db *Persistence) GetAllGrowSpotsByGrowLevelIDWithTransaction(tx *sqlx.Tx, req *routemodels.GetAllGrowSpotsRequest) (*routemodels.GetAllGrowSpotsResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result []routemodels.GrowSpot

	SQL := `
	SELECT
		grow_spots.id AS id,
		grow_spots.growing_level AS growing_level_id,
		grow_spots.display_name AS display_name,
		grow_spots.created AS created_at,
		grow_spots.updated AS updated_at,
		cre_member.id AS created_member_id,
		cre_member.display_name AS created_member_name,
		up_member.id AS updated_member_id,
		up_member.display_name AS updated_member_name
	FROM grow_spots
	LEFT JOIN members AS cre_member ON members.id = grow_spots.created_by
	LEFT JOIN members AS up_member ON members.id = grow_spots.updated_by
	WHERE grow_spots.growing_level = $1
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

	if err := tx.Select(&result, SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetAllGrowSpotsResponse{
		GrowSpots: result,
	}, nil
}

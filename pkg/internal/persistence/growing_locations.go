package persistence

import (
	routemodels "totally-legit-grow-management/v1/pkg/internal/route-models"

	"github.com/jmoiron/sqlx"
)

var _ IGrowingLocationsDB = &Persistence{}

func (db *Persistence) CreateGrowingLocation(req routemodels.CreateGrowingLocationRequest) (*routemodels.CreateGrowingLocationResponse, error) {
	var result string

	SQL := `
	INSERT INTO growing_locations
	(display_name, growing_group)
	VALUES ($1, $2)
	RETURNING id
	`

	args := []interface{}{
		req.DisplayName,
		req.GrowingGroupID,
	}

	// Make the appropiate SQL Call
	if err := db.Postgres.QueryRow(SQL, args...).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.CreateGrowingLocationResponse{
		ID: result,
	}, nil
}

func (db *Persistence) CreateGrowingLocationWithTransaction(tx *sqlx.Tx, req routemodels.CreateGrowingLocationRequest) (*routemodels.CreateGrowingLocationResponse, error) {
	var result string

	SQL := `
	INSERT INTO growing_locations
	(display_name, growing_group)
	VALUES ($1, $2)
	RETURNING id
	`

	args := []interface{}{
		req.DisplayName,
		req.GrowingGroupID,
	}

	// Make the appropiate SQL Call
	if err := tx.QueryRow(SQL, args...).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.CreateGrowingLocationResponse{
		ID: result,
	}, nil
}

func (db *Persistence) DeleteGrowingLocation(req routemodels.DeleteGrowingLocationRequest) error {
	var result string

	SQL := `
	UPDATE growing_locations SET archived = NOW()
	WHERE growing_locations.id = $1
	`

	// Make the appropiate SQL Call
	if err := db.Postgres.QueryRow(SQL, req.ID).Scan(result); err != nil {
		// handle err
		return err
	}

	return nil
}

func (db *Persistence) EditGrowingLocation(req routemodels.EditGrowingLocationRequest) (*routemodels.EditGrowingLocationResponse, error) {
	return nil, nil
}

func (db *Persistence) GetGrowingLocation(req routemodels.GetGrowingLocationRequest) (*routemodels.GetGrowingLocationResponse, error) {
	var result routemodels.GrowingLocation

	SQL := `
	SELECT
		growing_locations.id AS id,
		growing_locations.growing_group AS growing_group_id,
		growing_locations.display_name AS display_name,
		growing_locations.created AS created_at,
		growing_locations.updated AS updated_at,
		cre_member.id AS created_member_id,
		cre_member.display_name AS created_member_name,
		up_member.id AS updated_member_id,
		up_member.display_name AS updated_member_name
	FROM growing_locations
	LEFT JOIN members AS cre_member ON members.id = growing_locations.created_by
	LEFT JOIN members AS up_member ON members.id = growing_locations.updated_by
	WHERE growing_locations.id = $1
	`

	// Make the appropiate SQL Call
	if err := db.Postgres.Get(&result, SQL, req.ID); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetGrowingLocationResponse{
		GrowingLocation: result,
	}, nil
}

func (db *Persistence) GetGrowingLocationWithTransaction(tx *sqlx.Tx, req routemodels.GetGrowingLocationRequest) (*routemodels.GetGrowingLocationResponse, error) {
	var result routemodels.GrowingLocation

	SQL := `
	SELECT
		growing_locations.id AS id,
		growing_locations.growing_group AS growing_group_id,
		growing_locations.display_name AS display_name,
		growing_locations.created AS created_at,
		growing_locations.updated AS updated_at,
		cre_member.id AS created_member_id,
		cre_member.display_name AS created_member_name,
		up_member.id AS updated_member_id,
		up_member.display_name AS updated_member_name
	FROM growing_locations
	LEFT JOIN members AS cre_member ON members.id = growing_locations.created_by
	LEFT JOIN members AS up_member ON members.id = growing_locations.updated_by
	WHERE growing_locations.id = $1
	`

	// Make the appropiate SQL Call
	if err := db.Postgres.Get(&result, SQL, req.ID); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetGrowingLocationResponse{
		GrowingLocation: result,
	}, nil
}

func (db *Persistence) GetAllGrowingLocations(req routemodels.GetAllGrowingLocationsRequest) (*routemodels.GetAllGrowingLocationsResponse, error) {
	var result []routemodels.GrowingLocation

	SQL := `
	SELECT
		growing_locations.id AS id,
		growing_locations.growing_group AS growing_group_id,
		growing_locations.display_name AS display_name,
		growing_locations.created AS created_at,
		growing_locations.updated AS updated_at,
		cre_member.id AS created_member_id,
		cre_member.display_name AS created_member_name,
		up_member.id AS updated_member_id,
		up_member.display_name AS updated_member_name
	FROM growing_locations
	LEFT JOIN members AS cre_member ON members.id = growing_locations.created_by
	LEFT JOIN members AS up_member ON members.id = growing_locations.updated_by
	WHERE growing_locations.growing_group = $1
	`

	// Make the appropiate SQL Call
	if err := db.Postgres.Select(&result, SQL, req.GrowingGroupID); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetAllGrowingLocationsResponse{
		GrowingLocations: result,
	}, nil
}

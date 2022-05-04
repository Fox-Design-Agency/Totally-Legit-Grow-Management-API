package persistence

import (
	routemodels "totally-legit-grow-management/v1/pkg/internal/route-models"

	"github.com/jmoiron/sqlx"
)

var _ IGrowingLevelsDB = &Persistence{}

func (db *Persistence) CreateGrowingLevel(req routemodels.CreateGrowingLevelRequest) (*routemodels.CreateGrowingLevelResponse, error) {
	var result string

	SQL := `
	INSERT INTO growing_levels
	(display_name, growing_location)
	VALUES ($1, $2)
	RETURNING id
	`

	args := []interface{}{
		req.DisplayName,
		req.GrowingLocationID,
	}

	// Make the appropiate SQL Call
	if err := db.Postgres.QueryRow(SQL, args...).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.CreateGrowingLevelResponse{
		ID: result,
	}, nil
}

func (db *Persistence) CreateGrowingLevelWithTransaction(tx *sqlx.Tx, req routemodels.CreateGrowingLevelRequest) (*routemodels.CreateGrowingLevelResponse, error) {
	var result string

	SQL := `
	INSERT INTO growing_levels
	(display_name, growing_location)
	VALUES ($1, $2)
	RETURNING id
	`

	args := []interface{}{
		req.DisplayName,
		req.GrowingLocationID,
	}

	// Make the appropiate SQL Call
	if err := db.Postgres.QueryRow(SQL, args...).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.CreateGrowingLevelResponse{
		ID: result,
	}, nil
}

func (db *Persistence) DeleteGrowingLevel(req routemodels.DeleteGrowingLevelRequest) error {
	var result string

	SQL := `
	UPDATE growing_levels SET archived = NOW()
	WHERE growing_levels.id = $1
	`

	// Make the appropiate SQL Call
	if err := db.Postgres.QueryRow(SQL, req.ID).Scan(result); err != nil {
		// handle err
		return err
	}

	return nil
}

func (db *Persistence) EditGrowingLevel(req routemodels.EditGrowingLevelRequest) (*routemodels.EditGrowingLevelResponse, error) {
	return nil, nil
}

func (db *Persistence) GetGrowingLevel(req routemodels.GetGrowingLevelRequest) (*routemodels.GetGrowingLevelResponse, error) {
	var result routemodels.GrowingLevel

	SQL := `
	SELECT
		growing_levels.id AS id,
		growing_levels.growing_location AS growing_location_id,
		growing_levels.display_name AS display_name,
		growing_levels.created AS created_at,
		growing_levels.updated AS updated_at,
		cre_member.id AS created_member_id,
		cre_member.display_name AS created_member_name,
		up_member.id AS updated_member_id,
		up_member.display_name AS updated_member_name
	FROM growing_levels
	LEFT JOIN members AS cre_member ON members.id = growing_levels.created_by
	LEFT JOIN members AS up_member ON members.id = growing_levels.updated_by
	WHERE growing_levels.id = $1
	`

	// Make the appropiate SQL Call
	if err := db.Postgres.Get(&result, SQL, req.ID); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetGrowingLevelResponse{
		GrowingLevel: result,
	}, nil
}

func (db *Persistence) GetGrowingLevelWithTransaction(tx *sqlx.Tx, req routemodels.GetGrowingLevelRequest) (*routemodels.GetGrowingLevelResponse, error) {
	var result routemodels.GrowingLevel

	SQL := `
	SELECT
		growing_levels.id AS id,
		growing_levels.growing_location AS growing_location_id,
		growing_levels.display_name AS display_name,
		growing_levels.created AS created_at,
		growing_levels.updated AS updated_at,
		cre_member.id AS created_member_id,
		cre_member.display_name AS created_member_name,
		up_member.id AS updated_member_id,
		up_member.display_name AS updated_member_name
	FROM growing_levels
	LEFT JOIN members AS cre_member ON members.id = growing_levels.created_by
	LEFT JOIN members AS up_member ON members.id = growing_levels.updated_by
	WHERE growing_levels.id = $1
	`

	// Make the appropiate SQL Call
	if err := db.Postgres.Get(&result, SQL, req.ID); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetGrowingLevelResponse{
		GrowingLevel: result,
	}, nil
}

func (db *Persistence) GetAllGrowingLevels(req routemodels.GetAllGrowingLevelsRequest) (*routemodels.GetAllGrowingLevelsResponse, error) {
	var result []routemodels.GrowingLevel

	SQL := `
	SELECT
		growing_levels.id AS id,
		growing_levels.display_name AS display_name,
		growing_levels.created AS created_at,
		growing_levels.updated AS updated_at,
		cre_member.id AS created_member_id,
		cre_member.display_name AS created_member_name,
		up_member.id AS updated_member_id,
		up_member.display_name AS updated_member_name
	FROM growing_levels
	LEFT JOIN members AS cre_member ON members.id = growing_levels.created_by
	LEFT JOIN members AS up_member ON members.id = growing_levels.updated_by
	`

	// Make the appropiate SQL Call
	if err := db.Postgres.Select(&result, SQL); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetAllGrowingLevelsResponse{
		GrowingLevels: result,
	}, nil
}

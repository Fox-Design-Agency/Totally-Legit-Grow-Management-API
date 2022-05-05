package persistence

import routemodels "totally-legit-grow-management/v1/pkg/internal/route-models"

var _ IHarvestsDB = &Persistence{}

func (db *Persistence) CreateHarvest(req *routemodels.CreateHarvestRequest) (*routemodels.CreateHarvestResponse, error) {
	var result string

	SQL := `
	INSERT INTO growing_mediums
	(display_name)
	VALUES ($1)
	RETURNING id
	`

	args := []interface{}{
		req.DisplayName,
	}

	// Make the appropiate SQL Call
	if err := db.Postgres.QueryRow(SQL, args...).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.CreateHarvestResponse{
		ID: result,
	}, nil
}

func (db *Persistence) DeleteHarvest(req *routemodels.DeleteHarvestRequest) error {
	var result string

	SQL := `
	UPDATE growing_mediums SET archived = NOW()
	WHERE growing_mediums.id = $1
	`

	// Make the appropiate SQL Call
	if err := db.Postgres.QueryRow(SQL, req.ID).Scan(result); err != nil {
		// handle err
		return err
	}

	return nil
}

func (db *Persistence) EditHarvest(req *routemodels.EditHarvestRequest) (*routemodels.EditHarvestResponse, error) {
	return nil, nil
}

func (db *Persistence) GetHarvest(req *routemodels.GetHarvestRequest) (*routemodels.GetHarvestResponse, error) {
	var result routemodels.Harvest

	SQL := `
	SELECT
		growing_mediums.id AS id,
		growing_mediums.display_name AS display_name,
		growing_mediums.created AS created_at,
		growing_mediums.updated AS updated_at,
		cre_member.id AS created_member_id,
		cre_member.display_name AS created_member_name,
		up_member.id AS updated_member_id,
		up_member.display_name AS updated_member_name
	FROM growing_mediums
	LEFT JOIN members AS cre_member ON members.id = growing_mediums.created_by
	LEFT JOIN members AS up_member ON members.id = growing_mediums.updated_by
	WHERE growing_mediums.id = $1
	`

	// Make the appropiate SQL Call
	if err := db.Postgres.Get(&result, SQL, req.ID); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetHarvestResponse{
		Harvest: result,
	}, nil
}

func (db *Persistence) GetAllHarvests(*routemodels.GetAllHarvestsRequest) (*routemodels.GetAllHarvestsResponse, error) {
	var result []routemodels.Harvest

	SQL := `
	SELECT
		growing_mediums.id AS id,
		growing_mediums.display_name AS display_name,
		growing_mediums.created AS created_at,
		growing_mediums.updated AS updated_at,
		cre_member.id AS created_member_id,
		cre_member.display_name AS created_member_name,
		up_member.id AS updated_member_id,
		up_member.display_name AS updated_member_name
	FROM growing_mediums
	LEFT JOIN members AS cre_member ON members.id = growing_mediums.created_by
	LEFT JOIN members AS up_member ON members.id = growing_mediums.updated_by
	`

	// Make the appropiate SQL Call
	if err := db.Postgres.Select(&result, SQL); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetAllHarvestsResponse{
		Harvests: result,
	}, nil
}
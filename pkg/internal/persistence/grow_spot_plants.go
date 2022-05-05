package persistence

import routemodels "totally-legit-grow-management/v1/pkg/internal/route-models"

var _ IGrowSpotPlantsDB = &Persistence{}

func (db *Persistence) CreateGrowSpotPlant(req *routemodels.CreateGrowSpotPlantRequest) (*routemodels.CreateGrowSpotPlantResponse, error) {
	var result string

	SQL := `
	INSERT INTO grow_spot_plants
	(plant, plant_life_cycle, grow_spot, seed, seed_amount_units, seed_amount)
	VALUES ($1, $2, $3, $4, $5, $6)
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

	return &routemodels.CreateGrowSpotPlantResponse{
		ID: result,
	}, nil
}

func (db *Persistence) DeleteGrowSpotPlant(req *routemodels.DeleteGrowSpotPlantRequest) error {
	var result string

	SQL := `
	UPDATE grow_spot_plants SET archived = NOW()
	WHERE grow_spot_plants.id = $1
	`

	// Make the appropiate SQL Call
	if err := db.Postgres.QueryRow(SQL, req.ID).Scan(result); err != nil {
		// handle err
		return err
	}

	return nil
}

func (db *Persistence) EditGrowSpotPlant(req *routemodels.EditGrowSpotPlantRequest) (*routemodels.EditGrowSpotPlantResponse, error) {
	return nil, nil
}

func (db *Persistence) GetGrowSpotPlant(req *routemodels.GetGrowSpotPlantRequest) (*routemodels.GetGrowSpotPlantResponse, error) {
	var result routemodels.GrowSpotPlant

	SQL := `
	SELECT
		grow_spot_plants.id AS id,
		grow_spot_plants.created AS created_at,
		grow_spot_plants.updated AS updated_at,
		cre_member.id AS created_member_id,
		cre_member.display_name AS created_member_name,
		up_member.id AS updated_member_id,
		up_member.display_name AS updated_member_name
	FROM grow_spot_plants
	LEFT JOIN members AS cre_member ON members.id = grow_spot_plants.created_by
	LEFT JOIN members AS up_member ON members.id = grow_spot_plants.updated_by
	WHERE grow_spot_plants.id = $1
	`

	// Make the appropiate SQL Call
	if err := db.Postgres.Get(&result, SQL, req.ID); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetGrowSpotPlantResponse{
		GrowSpotPlant: result,
	}, nil
}

func (db *Persistence) GetAllGrowSpotPlants(req *routemodels.GetAllGrowSpotPlantsRequest) (*routemodels.GetAllGrowSpotPlantsResponse, error) {
	var result []routemodels.GrowSpotPlant

	SQL := `
	SELECT
		grow_spot_plants.id AS id,
		grow_spot_plants.created AS created_at,
		grow_spot_plants.updated AS updated_at,
		cre_member.id AS created_member_id,
		cre_member.display_name AS created_member_name,
		up_member.id AS updated_member_id,
		up_member.display_name AS updated_member_name
	FROM grow_spot_plants
	LEFT JOIN members AS cre_member ON members.id = grow_spot_plants.created_by
	LEFT JOIN members AS up_member ON members.id = grow_spot_plants.updated_by
	`

	// Make the appropiate SQL Call
	if err := db.Postgres.Select(&result, SQL); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetAllGrowSpotPlantsResponse{
		GrowSpotPlants: result,
	}, nil
}

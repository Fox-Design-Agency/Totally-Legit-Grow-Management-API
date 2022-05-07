package persistence

import routemodels "totally-legit-grow-management/v1/pkg/internal/route-models"

var _ IGrowSpotPlantsDB = &Persistence{}

func (db *Persistence) CreateGrowSpotPlant(req *routemodels.CreateGrowSpotPlantRequest) (*routemodels.CreateGrowSpotPlantResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result string

	SQL := `
	INSERT INTO grow_spot_plants
	(plant, plant_life_cycle, grow_spot, seed, seed_amount_units, seed_amount)
	VALUES ($1, $2, $3, $4, $5, $6)
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
	return &routemodels.CreateGrowSpotPlantResponse{
		ID: result,
	}, nil
}

func (db *Persistence) DeleteGrowSpotPlant(req *routemodels.DeleteGrowSpotPlantRequest) error {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result string

	SQL := `
	UPDATE grow_spot_plants SET archived = NOW()
	WHERE grow_spot_plants.id = $1
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

func (db *Persistence) EditGrowSpotPlant(req *routemodels.EditGrowSpotPlantRequest) (*routemodels.EditGrowSpotPlantResponse, error) {
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

func (db *Persistence) GetGrowSpotPlant(req *routemodels.GetGrowSpotPlantRequest) (*routemodels.GetGrowSpotPlantResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
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
	return &routemodels.GetGrowSpotPlantResponse{
		GrowSpotPlant: result,
	}, nil
}

func (db *Persistence) GetAllGrowSpotPlants(req *routemodels.GetAllGrowSpotPlantsRequest) (*routemodels.GetAllGrowSpotPlantsResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
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
	return &routemodels.GetAllGrowSpotPlantsResponse{
		GrowSpotPlants: result,
	}, nil
}

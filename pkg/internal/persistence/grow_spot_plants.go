package persistence

import routemodels "totally-legit-grow-management/v1/pkg/internal/route-models"

var _ IGrowSpotPlantsDB = &Persistence{}

// CreateGrowSpotPlant will insert a new Grow Spot Plant into the persistence layer and return the created ID
func (db *Persistence) CreateGrowSpotPlant(req *routemodels.CreateGrowSpotPlantRequest) (*routemodels.CreateGrowSpotPlantResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result string

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

	if err := db.Postgres.QueryRow(CREATE_GROW_SPOT_PLANT_SQL, args...).Scan(result); err != nil {
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

// DeleteGrowSpotPlant will set the Growing Spot Plant to be archived
func (db *Persistence) DeleteGrowSpotPlant(req *routemodels.DeleteGrowSpotPlantRequest) error {
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

	args := []interface{}{
		req.ID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if _, err := db.Postgres.Exec(DELETE_GROW_SPOT_PLANT_SQL, args...); err != nil {
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

// EditGrowSpotPlant is not currently implemented
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

// GetGrowSpotPlant will get a Grow Spot Plant from the persistence layer
func (db *Persistence) GetGrowSpotPlant(req *routemodels.GetGrowSpotPlantRequest) (*routemodels.GetGrowSpotPlantResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.GrowSpotPlant

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

	if err := db.Postgres.Get(&result, GET_GROW_SPOT_PLANT_BY_ID_SQL, args...); err != nil {
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

// GetAllGrowSpotPlants will get all Grow Spot Plants from the persistence layer
func (db *Persistence) GetAllGrowSpotPlants(req *routemodels.GetAllGrowSpotPlantsRequest) (*routemodels.GetAllGrowSpotPlantsResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result []routemodels.GrowSpotPlant

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

	if err := db.Postgres.Select(&result, GET_ALL_GROW_SPOT_PLANTS_SQL, args...); err != nil {
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

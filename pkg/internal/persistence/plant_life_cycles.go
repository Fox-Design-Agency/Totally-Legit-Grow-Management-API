package persistence

import (
	routemodels "totally-legit-grow-management/v1/pkg/internal/route-models"

	"github.com/jmoiron/sqlx"
)

var _ IPlantLifeCyclesDB = &Persistence{}

func (db *Persistence) CreatePlantLifeCycle(req *routemodels.CreatePlantLifeCycleRequest) (*routemodels.CreatePlantLifeCycleResponse, error) {
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
		req.TotalTimeMeasureUnits,
		req.EstTotalTimeMeasure,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := db.Postgres.QueryRow(CREATE_PLANT_LIFE_CYCLE_SQL, args...).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.CreatePlantLifeCycleResponse{
		ID: result,
	}, nil
}

func (db *Persistence) CreatePlantLifeCycleWithTransaction(tx *sqlx.Tx, req *routemodels.CreatePlantLifeCycleRequest) (*routemodels.CreatePlantLifeCycleResponse, error) {
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
		req.TotalTimeMeasureUnits,
		req.EstTotalTimeMeasure,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := tx.QueryRow(CREATE_PLANT_LIFE_CYCLE_SQL, args...).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.CreatePlantLifeCycleResponse{
		ID: result,
	}, nil
}

func (db *Persistence) DeletePlantLifeCycle(req *routemodels.DeletePlantLifeCycleRequest) error {
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
		req.ID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := db.Postgres.QueryRow(DELETE_PLANT_LIFE_CYCLE_SQL, args...).Scan(result); err != nil {
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

func (db *Persistence) EditPlantLifeCycle(req *routemodels.EditPlantLifeCycleRequest) (*routemodels.EditPlantLifeCycleResponse, error) {
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

func (db *Persistence) GetPlantLifeCycleByID(req *routemodels.GetPlantLifeCycleRequest) (*routemodels.GetPlantLifeCycleResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.PlantLifeCycle

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

	if err := db.Postgres.Get(&result, GET_PLANT_LIFE_CYCLE_BY_ID_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetPlantLifeCycleResponse{
		PlantLifeCycle: result,
	}, nil
}

func (db *Persistence) GetPlantLifeCycleByIDWithTransaction(tx *sqlx.Tx, req *routemodels.GetPlantLifeCycleRequest) (*routemodels.GetPlantLifeCycleResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.PlantLifeCycle

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

	if err := tx.Get(&result, GET_PLANT_LIFE_CYCLE_BY_ID_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetPlantLifeCycleResponse{
		PlantLifeCycle: result,
	}, nil
}

func (db *Persistence) GetAllPlantLifeCycles(req *routemodels.GetAllPlantLifeCyclesRequest) (*routemodels.GetAllPlantLifeCyclesResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result []routemodels.PlantLifeCycle

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

	if err := db.Postgres.Select(&result, GET_ALL_PLANT_LIFE_CYCLES_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetAllPlantLifeCyclesResponse{
		PlantLifeCycles: result,
	}, nil
}

func (db *Persistence) GetAllPlantLifeCyclesWithTransaction(tx *sqlx.Tx, req *routemodels.GetAllPlantLifeCyclesRequest) (*routemodels.GetAllPlantLifeCyclesResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result []routemodels.PlantLifeCycle

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

	if err := tx.Select(&result, GET_ALL_PLANT_LIFE_CYCLES_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetAllPlantLifeCyclesResponse{
		PlantLifeCycles: result,
	}, nil
}

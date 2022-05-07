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

	SQL := `
	INSERT INTO plant_life_cycles
	(display_name, total_time_measure_units, est_total_time_measure)
	VALUES ($1, $2, $3)
	RETURNING id
	`

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

	if err := db.Postgres.QueryRow(SQL, args...).Scan(result); err != nil {
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

	SQL := `
	INSERT INTO plant_life_cycles
	(display_name, total_time_measure_units, est_total_time_measure)
	VALUES ($1, $2, $3)
	RETURNING id
	`

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

	if err := tx.QueryRow(SQL, args...).Scan(result); err != nil {
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

	SQL := `
	UPDATE plant_life_cycles SET archived = NOW()
	WHERE plant_life_cycles.id = $1
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

	SQL := `
	SELECT
		plant_life_cycles.id AS id,
		plant_life_cycles.actual_time_measure AS actual_time_measure,
		plant_life_cycles.display_name AS display_name,
		plant_life_cycles.total_time_measure_units AS total_time_measure_units,
		plant_life_cycles.est_total_time_measure AS est_total_time_measure,
		plant_life_cycles.created AS created_at,
		plant_life_cycles.updated AS updated_at,
		cre_member.id AS created_member_id,
		cre_member.display_name AS created_member_name,
		up_member.id AS updated_member_id,
		up_member.display_name AS updated_member_name
	FROM plant_life_cycles
	LEFT JOIN members AS cre_member ON members.id = plant_life_cycles.created_by
	LEFT JOIN members AS up_member ON members.id = plant_life_cycles.updated_by
	WHERE plant_life_cycles.id = $1
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

	SQL := `
	SELECT
		plant_life_cycles.id AS id,
		plant_life_cycles.actual_time_measure AS actual_time_measure,
		plant_life_cycles.display_name AS display_name,
		plant_life_cycles.total_time_measure_units AS total_time_measure_units,
		plant_life_cycles.est_total_time_measure AS est_total_time_measure,
		plant_life_cycles.created AS created_at,
		plant_life_cycles.updated AS updated_at,
		cre_member.id AS created_member_id,
		cre_member.display_name AS created_member_name,
		up_member.id AS updated_member_id,
		up_member.display_name AS updated_member_name
	FROM plant_life_cycles
	LEFT JOIN members AS cre_member ON members.id = plant_life_cycles.created_by
	LEFT JOIN members AS up_member ON members.id = plant_life_cycles.updated_by
	WHERE plant_life_cycles.id = $1
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

	SQL := `
	SELECT
		plant_life_cycles.id AS id,
		plant_life_cycles.actual_time_measure AS actual_time_measure,
		plant_life_cycles.display_name AS display_name,
		plant_life_cycles.total_time_measure_units AS total_time_measure_units,
		plant_life_cycles.est_total_time_measure AS est_total_time_measure,
		plant_life_cycles.created AS created_at,
		plant_life_cycles.updated AS updated_at,
		cre_member.id AS created_member_id,
		cre_member.display_name AS created_member_name,
		up_member.id AS updated_member_id,
		up_member.display_name AS updated_member_name
	FROM plant_life_cycles
	LEFT JOIN members AS cre_member ON members.id = plant_life_cycles.created_by
	LEFT JOIN members AS up_member ON members.id = plant_life_cycles.updated_by
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

	SQL := `
	SELECT
		plant_life_cycles.id AS id,
		plant_life_cycles.actual_time_measure AS actual_time_measure,
		plant_life_cycles.display_name AS display_name,
		plant_life_cycles.total_time_measure_units AS total_time_measure_units,
		plant_life_cycles.est_total_time_measure AS est_total_time_measure,
		plant_life_cycles.created AS created_at,
		plant_life_cycles.updated AS updated_at,
		cre_member.id AS created_member_id,
		cre_member.display_name AS created_member_name,
		up_member.id AS updated_member_id,
		up_member.display_name AS updated_member_name
	FROM plant_life_cycles
	LEFT JOIN members AS cre_member ON members.id = plant_life_cycles.created_by
	LEFT JOIN members AS up_member ON members.id = plant_life_cycles.updated_by
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
	return &routemodels.GetAllPlantLifeCyclesResponse{
		PlantLifeCycles: result,
	}, nil
}

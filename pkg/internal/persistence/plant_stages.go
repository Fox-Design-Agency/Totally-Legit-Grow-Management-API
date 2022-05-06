package persistence

import (
	routemodels "totally-legit-grow-management/v1/pkg/internal/route-models"

	"github.com/jmoiron/sqlx"
)

var _ IPlantStagesDB = &Persistence{}

func (db *Persistence) CreatePlantStage(req *routemodels.CreatePlantStageRequest) (*routemodels.CreatePlantStageResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result string

	SQL := `
	INSERT INTO nutrients
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
	return &routemodels.CreatePlantStageResponse{
		ID: result,
	}, nil
}

func (db *Persistence) CreatePlantStageWithTransaction(tx *sqlx.Tx, req *routemodels.CreatePlantStageRequest) (*routemodels.CreatePlantStageResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result string

	SQL := `
	INSERT INTO nutrients
	(display_name)
	VALUES ($1)
	RETURNING id
	`

	// Make the appropiate SQL Call
	if err := tx.QueryRow(SQL, req.DisplayName).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.CreatePlantStageResponse{
		ID: result,
	}, nil
}

func (db *Persistence) CreatePlantStageCare(req *routemodels.CreatePlantStageCareRequest) (*routemodels.CreatePlantStageCareResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result string

	SQL := `
	INSERT INTO nutrients
	(display_name)
	VALUES ($1)
	RETURNING id
	`

	// Make the appropiate SQL Call
	if err := db.Postgres.QueryRow(SQL, req.DisplayName).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.CreatePlantStageCareResponse{
		ID: result,
	}, nil
}

func (db *Persistence) CreatePlantStageCareWithTransaction(tx *sqlx.Tx, req *routemodels.CreatePlantStageCareRequest) (*routemodels.CreatePlantStageCareResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result string

	SQL := `
	INSERT INTO nutrients
	(display_name)
	VALUES ($1)
	RETURNING id
	`

	// Make the appropiate SQL Call
	if err := tx.QueryRow(SQL, req.DisplayName).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.CreatePlantStageCareResponse{
		ID: result,
	}, nil
}

func (db *Persistence) CreatePlantStageNutrients(req *routemodels.CreatePlantStageNutrientsRequest) (*routemodels.CreatePlantStageNutrientsResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result string

	SQL := `
	INSERT INTO nutrients
	(display_name)
	VALUES ($1)
	RETURNING id
	`

	// Make the appropiate SQL Call
	if err := db.Postgres.QueryRow(SQL, req.DisplayName).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.CreatePlantStageNutrientsResponse{
		ID: result,
	}, nil
}

func (db *Persistence) CreatePlantStageNutrientsWithTransaction(tx *sqlx.Tx, req *routemodels.CreatePlantStageNutrientsRequest) (*routemodels.CreatePlantStageNutrientsResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result string

	SQL := `
	INSERT INTO nutrients
	(display_name)
	VALUES ($1)
	RETURNING id
	`

	// Make the appropiate SQL Call
	if err := tx.QueryRow(SQL, req.DisplayName).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.CreatePlantStageNutrientsResponse{
		ID: result,
	}, nil
}

func (db *Persistence) ConnectPlantStage(req *routemodels.ConnectPlantStageRequest) (*routemodels.ConnectPlantStageResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result string

	SQL := `
	INSERT INTO nutrients
	(display_name)
	VALUES ($1)
	RETURNING id
	`

	// Make the appropiate SQL Call
	if err := db.Postgres.QueryRow(SQL, req.DisplayName).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.ConnectPlantStageResponse{
		ID: result,
	}, nil
}

func (db *Persistence) ConnectPlantStageWithTransaction(tx *sqlx.Tx, req *routemodels.ConnectPlantStageRequest) (*routemodels.ConnectPlantStageResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result string

	SQL := `
	INSERT INTO nutrients
	(display_name)
	VALUES ($1)
	RETURNING id
	`

	// Make the appropiate SQL Call
	if err := tx.QueryRow(SQL, req.DisplayName).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.ConnectPlantStageResponse{
		ID: result,
	}, nil
}

func (db *Persistence) DeletePlantStage(req *routemodels.DeletePlantStageRequest) error {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result string

	SQL := `
	UPDATE nutrients SET archived = NOW()
	WHERE nutrients.id = $1
	`

	// Make the appropiate SQL Call
	if err := db.Postgres.QueryRow(SQL, req.ID).Scan(result); err != nil {
		// handle err
		return err
	}

	return nil
}

func (db *Persistence) EditPlantStage(req *routemodels.EditPlantStageRequest) (*routemodels.EditPlantStageResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	return nil, nil
}

func (db *Persistence) GetPlantStageByID(req *routemodels.GetPlantStageRequest) (*routemodels.GetPlantStageResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.PlantStage

	SQL := `
	SELECT
		nutrients.id AS id,
		nutrients.display_name AS display_name,
		nutrients.created AS created_at,
		nutrients.updated AS updated_at,
		cre_member.id AS created_member_id,
		cre_member.display_name AS created_member_name,
		up_member.id AS updated_member_id,
		up_member.display_name AS updated_member_name
	FROM nutrients
	LEFT JOIN members AS cre_member ON members.id = nutrients.created_by
	LEFT JOIN members AS up_member ON members.id = nutrients.updated_by
	WHERE nutrients.id = $1
	`

	// Make the appropiate SQL Call
	if err := db.Postgres.Get(&result, SQL, req.ID); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetPlantStageResponse{
		PlantStage: result,
	}, nil
}

func (db *Persistence) GetPlantStageByIDWithTransaction(tx *sqlx.Tx, req *routemodels.GetPlantStageRequest) (*routemodels.GetPlantStageResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.PlantStage

	SQL := `
	SELECT
		nutrients.id AS id,
		nutrients.display_name AS display_name,
		nutrients.created AS created_at,
		nutrients.updated AS updated_at,
		cre_member.id AS created_member_id,
		cre_member.display_name AS created_member_name,
		up_member.id AS updated_member_id,
		up_member.display_name AS updated_member_name
	FROM nutrients
	LEFT JOIN members AS cre_member ON members.id = nutrients.created_by
	LEFT JOIN members AS up_member ON members.id = nutrients.updated_by
	WHERE nutrients.id = $1
	`

	// Make the appropiate SQL Call
	if err := tx.Get(&result, SQL, req.ID); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetPlantStageResponse{
		PlantStage: result,
	}, nil
}

func (db *Persistence) GetPlantStageCareByID(req *routemodels.GetPlantStageCareByIDRequest) (*routemodels.GetPlantStageCareByIDResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.PlantStage

	SQL := `
	SELECT
		nutrients.id AS id,
		nutrients.display_name AS display_name,
		nutrients.created AS created_at,
		nutrients.updated AS updated_at,
		cre_member.id AS created_member_id,
		cre_member.display_name AS created_member_name,
		up_member.id AS updated_member_id,
		up_member.display_name AS updated_member_name
	FROM nutrients
	LEFT JOIN members AS cre_member ON members.id = nutrients.created_by
	LEFT JOIN members AS up_member ON members.id = nutrients.updated_by
	WHERE nutrients.id = $1
	`

	// Make the appropiate SQL Call
	if err := db.Postgres.Get(&result, SQL, req.ID); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetPlantStageCareByIDResponse{
		PlantStage: result,
	}, nil
}

func (db *Persistence) GetPlantStageCareByPlantStageIDWithTransaction(tx *sqlx.Tx, req *routemodels.GetPlantStageCareByPlantStageIDRequest) (*routemodels.GetPlantStageCareByPlantStageIDResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.PlantStage

	SQL := `
	SELECT
		nutrients.id AS id,
		nutrients.display_name AS display_name,
		nutrients.created AS created_at,
		nutrients.updated AS updated_at,
		cre_member.id AS created_member_id,
		cre_member.display_name AS created_member_name,
		up_member.id AS updated_member_id,
		up_member.display_name AS updated_member_name
	FROM nutrients
	LEFT JOIN members AS cre_member ON members.id = nutrients.created_by
	LEFT JOIN members AS up_member ON members.id = nutrients.updated_by
	WHERE nutrients.id = $1
	`

	// Make the appropiate SQL Call
	if err := tx.Get(&result, SQL, req.OrganizationID); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetPlantStageCareByPlantStageIDResponse{
		PlantStage: result,
	}, nil
}

func (db *Persistence) GetPlantStageNutrient(req *routemodels.GetPlantStageNutrientRequest) (*routemodels.GetPlantStageNutrientResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.PlantStage

	SQL := `
	SELECT
		nutrients.id AS id,
		nutrients.display_name AS display_name,
		nutrients.created AS created_at,
		nutrients.updated AS updated_at,
		cre_member.id AS created_member_id,
		cre_member.display_name AS created_member_name,
		up_member.id AS updated_member_id,
		up_member.display_name AS updated_member_name
	FROM nutrients
	LEFT JOIN members AS cre_member ON members.id = nutrients.created_by
	LEFT JOIN members AS up_member ON members.id = nutrients.updated_by
	WHERE nutrients.id = $1
	`

	// Make the appropiate SQL Call
	if err := db.Postgres.Get(&result, SQL, req.OrganizationID); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetPlantStageNutrientResponse{
		PlantStage: result,
	}, nil
}

func (db *Persistence) GetAllPlantStages(req *routemodels.GetAllPlantStagesRequest) (*routemodels.GetAllPlantStagesResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result []routemodels.PlantStage

	SQL := `
	SELECT
		nutrients.id AS id,
		nutrients.display_name AS display_name,
		nutrients.created AS created_at,
		nutrients.updated AS updated_at,
		cre_member.id AS created_member_id,
		cre_member.display_name AS created_member_name,
		up_member.id AS updated_member_id,
		up_member.display_name AS updated_member_name
	FROM nutrients
	LEFT JOIN members AS cre_member ON members.id = nutrients.created_by
	LEFT JOIN members AS up_member ON members.id = nutrients.updated_by
	`

	// Make the appropiate SQL Call
	if err := db.Postgres.Select(&result, SQL); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetAllPlantStagesResponse{
		PlantStages: result,
	}, nil
}

func (db *Persistence) GetAllPlantStagesWithTransaction(tx *sqlx.Tx, req *routemodels.GetAllPlantStagesRequest) (*routemodels.GetAllPlantStagesResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result []routemodels.PlantStage

	SQL := `
	SELECT
		nutrients.id AS id,
		nutrients.display_name AS display_name,
		nutrients.created AS created_at,
		nutrients.updated AS updated_at,
		cre_member.id AS created_member_id,
		cre_member.display_name AS created_member_name,
		up_member.id AS updated_member_id,
		up_member.display_name AS updated_member_name
	FROM nutrients
	LEFT JOIN members AS cre_member ON members.id = nutrients.created_by
	LEFT JOIN members AS up_member ON members.id = nutrients.updated_by
	`

	// Make the appropiate SQL Call
	if err := tx.Select(&result, SQL); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetAllPlantStagesResponse{
		PlantStages: result,
	}, nil
}

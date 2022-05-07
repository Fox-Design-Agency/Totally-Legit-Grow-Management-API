package persistence

import routemodels "totally-legit-grow-management/v1/pkg/internal/route-models"

var _ INutrientsDB = &Persistence{}

func (db *Persistence) CreateNutrient(req *routemodels.CreateNutrientRequest) (*routemodels.CreateNutrientResponse, error) {
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
	return &routemodels.CreateNutrientResponse{
		ID: result,
	}, nil
}

func (db *Persistence) DeleteNutrient(req *routemodels.DeleteNutrientRequest) error {
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

func (db *Persistence) EditNutrient(req *routemodels.EditNutrientRequest) (*routemodels.EditNutrientResponse, error) {
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

func (db *Persistence) GetNutrient(req *routemodels.GetNutrientRequest) (*routemodels.GetNutrientResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.Nutrient

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
	return &routemodels.GetNutrientResponse{
		Nutrient: result,
	}, nil
}

func (db *Persistence) GetAllNutrients(req *routemodels.GetAllNutrientsRequest) (*routemodels.GetAllNutrientsResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result []routemodels.Nutrient

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
	return &routemodels.GetAllNutrientsResponse{
		Nutrients: result,
	}, nil
}

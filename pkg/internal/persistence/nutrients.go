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

	if err := db.Postgres.QueryRow(CREATE_NUTRIENT_SQL, args...).Scan(result); err != nil {
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

	if _, err := db.Postgres.Exec(DELETE_NUTRIENT_SQL, args...); err != nil {
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

	if err := db.Postgres.Get(&result, GET_NUTRIENT_BY_ID_SQL, args...); err != nil {
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

	if err := db.Postgres.Select(&result, GET_ALL_NUTRIENTS_SQL, args...); err != nil {
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

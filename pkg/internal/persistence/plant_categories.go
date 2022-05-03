package persistence

import routemodels "smart-grow-management-api/v1/pkg/internal/route-models"

var _plantcats IPlantCategoriesDB = &Persistence{}

func (db *Persistence) CreatePlantCategory(req routemodels.CreatePlantCategoryRequest) (*routemodels.CreatePlantCategoryResponse, error) {
	var result string

	SQL := `
	INSERT INTO plant_categories
	(display_name)
	VALUES ($1)
	RETURNING id
	`

	// Make the appropiate SQL Call
	if err := db.Postgres.QueryRow(SQL, req.DisplayName).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.CreatePlantCategoryResponse{
		ID: result,
	}, nil
}

func (db *Persistence) DeletePlantCategory(req routemodels.DeletePlantCategoryRequest) error {
	var result string

	SQL := `
	UPDATE plant_categories SET archived = NOW()
	WHERE plant_categories.id = $1
	`

	// Make the appropiate SQL Call
	if err := db.Postgres.QueryRow(SQL, req.ID).Scan(result); err != nil {
		// handle err
		return err
	}

	return nil
}

func (db *Persistence) EditPlantCategory(req routemodels.EditPlantCategoryRequest) (*routemodels.EditPlantCategoryResponse, error) {
	return nil, nil
}

func (db *Persistence) GetPlantCategory(req routemodels.GetPlantCategoryRequest) (*routemodels.GetPlantCategoryResponse, error) {
	var result routemodels.PlantCategory

	SQL := `
	SELECT
		plant_categories.id AS id,
		plant_categories.display_name AS display_name,
		plant_categories.created AS created_at,
		plant_categories.updated AS updated_at,
		cre_member.id AS created_member_id,
		cre_member.display_name AS created_member_name,
		up_member.id AS updated_member_id,
		up_member.display_name AS updated_member_name
	FROM plant_categories
	LEFT JOIN members AS cre_member ON members.id = plant_categories.created_by
	LEFT JOIN members AS up_member ON members.id = plant_categories.updated_by
	WHERE plant_categories.id = $1
	`

	// Make the appropiate SQL Call
	if err := db.Postgres.Get(&result, SQL, req.ID); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetPlantCategoryResponse{
		PlantCategory: result,
	}, nil
}

func (db *Persistence) GetAllPlantCategories(req routemodels.GetAllPlantCategoriesRequest) (*routemodels.GetAllPlantCategoriesResponse, error) {
	var result []routemodels.PlantCategory

	SQL := `
	SELECT
		plant_categories.id AS id,
		plant_categories.display_name AS display_name,
		plant_categories.created AS created_at,
		plant_categories.updated AS updated_at,
		cre_member.id AS created_member_id,
		cre_member.display_name AS created_member_name,
		up_member.id AS updated_member_id,
		up_member.display_name AS updated_member_name
	FROM plant_categories
	LEFT JOIN members AS cre_member ON members.id = plant_categories.created_by
	LEFT JOIN members AS up_member ON members.id = plant_categories.updated_by
	`

	// Make the appropiate SQL Call
	if err := db.Postgres.Select(&result, SQL); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetAllPlantCategoriesResponse{
		PlantCategories: result,
	}, nil
}

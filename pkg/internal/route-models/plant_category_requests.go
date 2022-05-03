package routemodels

type CreatePlantCategoryRequest struct {
	DisplayName string `json:"displayName"`
}

type CreatePlantCategoryResponse struct {
	ID string `json:"id" db:"id"`
}

type DeletePlantCategoryRequest struct {
	ID string `json:"id"`
}

type EditPlantCategoryRequest struct {
	DisplayName string `json:"displayName"`
}

type EditPlantCategoryResponse struct {
	PlantCategory
}

type GetPlantCategoryRequest struct {
	ID string `json:"id"`
}

type GetPlantCategoryResponse struct {
	PlantCategory
}

type GetAllPlantCategoriesRequest struct {
	OrganizationID string `json:"organizationID"`
}

type GetAllPlantCategoriesResponse struct {
	PlantCategories []PlantCategory
}

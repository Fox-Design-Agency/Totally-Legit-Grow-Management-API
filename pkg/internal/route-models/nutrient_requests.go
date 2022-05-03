package routemodels

type CreateNutrientRequest struct {
	DisplayName string `json:"displayName"`
}

type CreateNutrientResponse struct {
	ID string `json:"id" db:"id"`
}

type DeleteNutrientRequest struct {
	ID string `json:"id"`
}

type EditNutrientRequest struct {
	DisplayName string `json:"displayName"`
}

type EditNutrientResponse struct {
	Nutrient
}

type GetNutrientRequest struct {
	ID string `json:"id"`
}

type GetNutrientResponse struct {
	Nutrient
}

type GetAllNutrientsRequest struct {
	OrganizationID []Nutrient
}

type GetAllNutrientsResponse struct {
	Nutrients []Nutrient
}

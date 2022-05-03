package routemodels

type SlimNutrient struct {
	ID          string `json:"id" db:"id"`
	DisplayName string `json:"displayName" db:"display_name"`
}

type Nutrient struct {
	SlimNutrient
	CreatedAtMember
	UpdatedAtMember
}

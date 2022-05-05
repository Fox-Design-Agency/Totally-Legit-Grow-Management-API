package routemodels

type SlimGrowingMedium struct {
	ID          string `json:"id" db:"id"`
	DisplayName string `json:"displayName" db:"display_name"`
}

type GrowingMedium struct {
	SlimGrowingMedium
	CreatedAtMember
	UpdatedAtMember
}

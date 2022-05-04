package routemodels

type SlimGrowingLevel struct {
	ID                string `json:"id" db:"id"`
	DisplayName       string `json:"displayName" db:"display_name"`
	GrowingLocationID string `json:"growingLocationID" db:"growing_location_id"`
}

type GrowingLevel struct {
	SlimGrowingLevel
	CreatedAtMember
	UpdatedAtMember
}

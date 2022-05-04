package routemodels

type SlimGrowSpot struct {
	ID             string `json:"id" db:"id"`
	DisplayName    string `json:"displayName" db:"display_name"`
	GrowingLevelID string `json:"growingLevelID" db:"growing_level_id"`
}

type GrowSpot struct {
	SlimGrowSpot
	CreatedAtMember
	UpdatedAtMember
}

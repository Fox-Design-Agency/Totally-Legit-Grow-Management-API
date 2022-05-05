package routemodels

type SlimGrowSpotPlant struct {
	ID          string `json:"id" db:"id"`
	DisplayName string `json:"displayName" db:"display_name"`
}

type GrowSpotPlant struct {
	SlimGrowSpotPlant
	CreatedAtMember
	UpdatedAtMember
}

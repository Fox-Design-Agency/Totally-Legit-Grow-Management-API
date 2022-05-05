package routemodels

type SlimHarvest struct {
	ID          string `json:"id" db:"id"`
	DisplayName string `json:"displayName" db:"display_name"`
}

type Harvest struct {
	SlimHarvest
	CreatedAtMember
	UpdatedAtMember
}

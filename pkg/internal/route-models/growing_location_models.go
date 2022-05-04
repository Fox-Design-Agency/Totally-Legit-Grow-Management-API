package routemodels

type SlimGrowingLocation struct {
	ID             string `json:"id" db:"id"`
	DisplayName    string `json:"displayName" db:"display_name"`
	GrowingGroupID string `json:"growingGroupID" db:"growing_group_id"`
}

type GrowingLocation struct {
	SlimGrowingLocation
	CreatedAtMember
	UpdatedAtMember
}

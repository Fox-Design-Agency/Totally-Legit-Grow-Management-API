package routemodels

type SlimGrowingGroup struct {
	ID             string `json:"id" db:"id"`
	OrganizationID string `json:"organizationID" db:"organization_id"`
	DisplayName    string `json:"displayName" db:"display_name"`
}

type GrowingGroup struct {
	SlimGrowingGroup
	CreatedAtMember
	UpdatedAtMember
}

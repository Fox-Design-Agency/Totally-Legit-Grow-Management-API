package routemodels

type SlimSeed struct {
	ID          string `json:"id" db:"id"`
	DisplayName string `json:"displayName" db:"display_name"`
}

type Seed struct {
	SlimSeed
	CreatedAtMember
	UpdatedAtMember
}

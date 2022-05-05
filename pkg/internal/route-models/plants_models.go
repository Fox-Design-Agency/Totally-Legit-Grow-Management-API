package routemodels

type SlimPlant struct {
	ID          string `json:"id" db:"id"`
	DisplayName string `json:"displayName" db:"display_name"`
}

type Plant struct {
	SlimPlant
	CreatedAtMember
	UpdatedAtMember
}

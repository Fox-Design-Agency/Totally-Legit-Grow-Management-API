package routemodels

type SlimInventory struct {
	ID          string `json:"id" db:"id"`
	DisplayName string `json:"displayName" db:"display_name"`
}

type Inventory struct {
	SlimInventory
	CreatedAtMember
	UpdatedAtMember
}

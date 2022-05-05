package routemodels

type SlimProduct struct {
	ID          string `json:"id" db:"id"`
	DisplayName string `json:"displayName" db:"display_name"`
}

type Product struct {
	SlimProduct
	CreatedAtMember
	UpdatedAtMember
}

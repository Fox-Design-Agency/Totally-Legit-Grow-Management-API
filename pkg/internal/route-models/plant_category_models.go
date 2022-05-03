package routemodels

type SlimPlantCategory struct {
	ID          string `json:"id" db:"id"`
	DisplayName string `json:"displayName" db:"display_name"`
}

type PlantCategory struct {
	SlimPlantCategory
	CreatedAtMember
	UpdatedAtMember
}

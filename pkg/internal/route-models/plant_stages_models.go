package routemodels

type SlimPlantStage struct {
	ID          string `json:"id" db:"id"`
	DisplayName string `json:"displayName" db:"display_name"`
}

type PlantStage struct {
	SlimPlantStage
	CreatedAtMember
	UpdatedAtMember
}

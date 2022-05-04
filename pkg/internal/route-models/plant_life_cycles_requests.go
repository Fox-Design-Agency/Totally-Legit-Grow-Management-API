package routemodels

type CreatePlantLifeCycleRequest struct {
	DisplayName           string `json:"displayName" db:"display_name"`
	TotalTimeMeasureUnits string `json:"totalTimeMeasureUnits" db:"total_time_measure_units"`
	EstTotalTimeMeasure   int    `json:"estTotalTimeMeasure" db:"est_total_time_measure"`
}

type CreatePlantLifeCycleResponse struct {
	ID string `json:"id" db:"id"`
}

type DeletePlantLifeCycleRequest struct {
	ID string `json:"id"`
}

type EditPlantLifeCycleRequest struct {
	DisplayName string `json:"displayName"`
}

type EditPlantLifeCycleResponse struct {
	PlantLifeCycle
}

type GetPlantLifeCycleRequest struct {
	ID string `json:"id"`
}

type GetPlantLifeCycleResponse struct {
	PlantLifeCycle
}

type GetAllPlantLifeCyclesRequest struct {
	OrganizationID string `json:"organizationID"`
}

type GetAllPlantLifeCyclesResponse struct {
	PlantLifeCycles []PlantLifeCycle
}

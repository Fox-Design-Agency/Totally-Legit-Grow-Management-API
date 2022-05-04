package routemodels

type SlimPlantLifeCycle struct {
	ID                    string  `json:"id" db:"id"`
	ActualTimeMeasure     float64 `json:"actualTimeMeasure" db:"actual_time_measure"`
	DisplayName           string  `json:"displayName" db:"display_name"`
	TotalTimeMeasureUnits string  `json:"totalTimeMeasureUnits" db:"total_time_measure_units"`
	EstTotalTimeMeasure   int     `json:"estTotalTimeMeasure" db:"est_total_time_measure"`
}

type PlantLifeCycle struct {
	SlimPlantLifeCycle
	CreatedAtMember
	UpdatedAtMember
}

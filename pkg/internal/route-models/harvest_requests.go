package routemodels

type CreateHarvestRequest struct {
	DisplayName string `json:"displayName"`
}

type CreateHarvestResponse struct {
	ID string `json:"id" db:"id"`
}

type DeleteHarvestRequest struct {
	ID string `json:"id"`
}

type EditHarvestRequest struct {
	DisplayName string `json:"displayName"`
}

type EditHarvestResponse struct {
	Harvest
}

type GetHarvestRequest struct {
	ID string `json:"id"`
}

type GetHarvestResponse struct {
	Harvest
}

type GetAllHarvestsRequest struct {
	OrganizationID string `json:"organizationID"`
}

type GetAllHarvestsResponse struct {
	Harvests []Harvest
}

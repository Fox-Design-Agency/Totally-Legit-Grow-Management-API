package routemodels

type CreateGrowSpotRequest struct {
	DisplayName    string `json:"displayName"`
	GrowingLevelID string `json:"growingLevelID"`
}

type CreateGrowSpotResponse struct {
	ID string `json:"id" db:"id"`
}

type DeleteGrowSpotRequest struct {
	ID string `json:"id"`
}

type EditGrowSpotRequest struct {
	DisplayName string `json:"displayName"`
}

type EditGrowSpotResponse struct {
	GrowSpot
}

type GetGrowSpotRequest struct {
	ID string `json:"id"`
}

type GetGrowSpotResponse struct {
	GrowSpot
}

type GetAllGrowSpotsRequest struct {
	GrowingLevelID string `json:"growingLevelID"`
}

type GetAllGrowSpotsResponse struct {
	GrowSpots []GrowSpot
}

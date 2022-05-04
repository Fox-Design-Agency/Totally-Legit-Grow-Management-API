package routemodels

type CreateGrowingLocationRequest struct {
	DisplayName    string `json:"displayName"`
	GrowingGroupID string `json:"growingGroupID"`
}

type CreateGrowingLocationResponse struct {
	ID string `json:"id" db:"id"`
}

type DeleteGrowingLocationRequest struct {
	ID string `json:"id"`
}

type EditGrowingLocationRequest struct {
	DisplayName string `json:"displayName"`
}

type EditGrowingLocationResponse struct {
	GrowingLocation
}

type GetGrowingLocationRequest struct {
	ID string `json:"id"`
}

type GetGrowingLocationResponse struct {
	GrowingLocation
}

type GetAllGrowingLocationsRequest struct {
	GrowingGroupID string `json:"growingGroupID"`
}

type GetAllGrowingLocationsResponse struct {
	GrowingLocations []GrowingLocation
}

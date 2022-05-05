package routemodels

type CreateGrowingMediumRequest struct {
	DisplayName string `json:"displayName"`
}

type CreateGrowingMediumResponse struct {
	ID string `json:"id" db:"id"`
}

type DeleteGrowingMediumRequest struct {
	ID string `json:"id"`
}

type EditGrowingMediumRequest struct {
	DisplayName string `json:"displayName"`
}

type EditGrowingMediumResponse struct {
	GrowingMedium
}

type GetGrowingMediumRequest struct {
	ID string `json:"id"`
}

type GetGrowingMediumResponse struct {
	GrowingMedium
}

type GetAllGrowingMediumsRequest struct {
	OrganizationID string `json:"organizationID"`
}

type GetAllGrowingMediumsResponse struct {
	GrowingMediums []GrowingMedium
}

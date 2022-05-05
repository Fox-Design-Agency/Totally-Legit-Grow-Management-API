package routemodels

type CreateSeedRequest struct {
	DisplayName string `json:"displayName"`
}

type CreateSeedResponse struct {
	ID string `json:"id" db:"id"`
}

type DeleteSeedRequest struct {
	ID string `json:"id"`
}

type EditSeedRequest struct {
	DisplayName string `json:"displayName"`
}

type EditSeedResponse struct {
	Seed
}

type GetSeedRequest struct {
	ID string `json:"id"`
}

type GetSeedResponse struct {
	Seed
}

type GetAllSeedsRequest struct {
	OrganizationID string `json:"organizationID"`
}

type GetAllSeedsResponse struct {
	Seeds []Seed
}

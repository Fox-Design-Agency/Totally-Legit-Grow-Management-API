package routemodels

type CreateOrganizationRequest struct {
	DisplayName string `json:"displayName"`
}

type CreateOrganizationResponse struct {
	ID string `json:"id" db:"id"`
}

type DeleteOrganizationRequest struct {
	ID string `json:"id"`
}

type EditOrganizationRequest struct {
	DisplayName string `json:"displayName"`
}

type EditOrganizationResponse struct {
	Organization
}

type GetOrganizationRequest struct {
	ID string `json:"id"`
}

type GetOrganizationResponse struct {
	Organization
}

type GetAllOrganizationsRequest struct {
}

type GetAllOrganizationsResponse struct {
	Organizations []Organization
}

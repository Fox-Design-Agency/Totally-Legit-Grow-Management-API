package routemodels

type CreateGrowingGroupRequest struct {
	DisplayName    string `json:"displayName"`
	OrganizationID string `json:"organizationID"`
}

type CreateGrowingGroupResponse struct {
	ID string `json:"id" db:"id"`
}

type DeleteGrowingGroupRequest struct {
	ID string `json:"id"`
}

type EditGrowingGroupRequest struct {
	DisplayName string `json:"displayName"`
}

type EditGrowingGroupResponse struct {
	GrowingGroup
}

type GetGrowingGroupRequest struct {
	ID string `json:"id"`
}

type GetGrowingGroupResponse struct {
	GrowingGroup
}

type GetAllGrowingGroupsRequest struct {
	OrganizationID string `json:"organizationID"`
}

type GetAllGrowingGroupsResponse struct {
	Groups []GrowingGroup
}

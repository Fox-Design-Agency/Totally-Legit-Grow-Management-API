package routemodels

import "github.com/lib/pq"

type SlimOrganization struct {
	ID          string `json:"id" db:"id"`
	DisplayName string `json:"displayName" db:"display_name"`
}

type Organization struct {
	SlimPlantCategory
	CreatedAt pq.NullTime `json:"createdAt" db:"created_at"`
	CreatedAtMember
	UpdatedAt pq.NullTime `json:"updatedAt" db:"updated_by"`
	UpdatedAtMember
}

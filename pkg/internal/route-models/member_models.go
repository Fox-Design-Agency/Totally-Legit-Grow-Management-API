package routemodels

import "github.com/lib/pq"

type SlimMember struct {
	ID          string `json:"id" db:"member_id"`
	DisplayName string `json:"displayName" db:"display_name"`
}

type CreatedAtMember struct {
	CreatedAt     pq.NullTime `json:"createdAt" db:"created_at"`
	CreatedByID   string      `json:"createdByID" db:"created_member_id"`
	CreatedByName string      `json:"createdByName" db:"created_member_name"`
}

type UpdatedAtMember struct {
	UpdatedAt     pq.NullTime `json:"updatedAt" db:"updated_at"`
	UpdatedByID   string      `json:"updatedByID" db:"updated_member_id"`
	UpdatedByName string      `json:"updatedByName" db:"updated_member_name"`
}

type Member struct {
	SlimNutrient
	CreatedAtMember
	UpdatedAtMember
}

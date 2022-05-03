package routemodels

import "github.com/lib/pq"

type SlimNutrient struct {
	ID          string `json:"id" db:"id"`
	DisplayName string `json:"displayName" db:"display_name"`
}

type Nutrient struct {
	SlimNutrient
	CreatedAt     pq.NullTime `json:"createdAt" db:"created_at"`
	CreatedByID   string      `json:"createdByID" db:"created_member_id"`
	CreatedByName string      `json:"createdByName" db:"created_member_name"`
	UpdatedAt     pq.NullTime `json:"updatedAt" db:"updated_by"`
	UpdatedByID   string      `json:"updatedByID" db:"updated_member_id"`
	UpdatedByName string      `json:"updatedByName" db:"updated_member_name"`
}

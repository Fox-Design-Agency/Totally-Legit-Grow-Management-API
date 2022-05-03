package routemodels

type CreatedAtMember struct {
	CreatedByID   string `json:"createdByID" db:"created_member_id"`
	CreatedByName string `json:"createdByName" db:"created_member_name"`
}

type UpdatedAtMember struct {
	UpdatedByID   string `json:"updatedByID" db:"updated_member_id"`
	UpdatedByName string `json:"updatedByName" db:"updated_member_name"`
}

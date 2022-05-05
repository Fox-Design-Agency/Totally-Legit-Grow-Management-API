package routemodels

type SlimTask struct {
	ID          string `json:"id" db:"id"`
	DisplayName string `json:"displayName" db:"display_name"`
}

type Task struct {
	SlimTask
	CreatedAtMember
	UpdatedAtMember
}

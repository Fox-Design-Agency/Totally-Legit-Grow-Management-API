package persistence

const (
	CREATE_DEVICE_SQL = `
		INSERT INTO devices
		(display_name)
		VALUES ($1)
		RETURNING id
		`
)

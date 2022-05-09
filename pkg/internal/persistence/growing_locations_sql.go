package persistence

const (
	CREATE_GROWING_LOCATION_SQL = `
		INSERT INTO growing_locations
		(display_name, growing_group)
		VALUES ($1, $2)
		RETURNING id
		`

	DELETE_GROWING_LOCATION_SQL = `
		UPDATE growing_locations SET archived = NOW()
		WHERE growing_locations.id = $1
		`

	GET_GROWING_LOCATION_BASE_SQL = `
		SELECT
			growing_locations.id AS id,
			growing_locations.growing_group AS growing_group_id,
			growing_locations.display_name AS display_name,
			growing_locations.created AS created_at,
			growing_locations.updated AS updated_at,
			cre_member.id AS created_member_id,
			cre_member.display_name AS created_member_name,
			up_member.id AS updated_member_id,
			up_member.display_name AS updated_member_name
		FROM growing_locations
		LEFT JOIN members AS cre_member ON members.id = growing_locations.created_by
		LEFT JOIN members AS up_member ON members.id = growing_locations.updated_by
		
		`

	GET_GROWING_LOCATION_BY_ID_SQL = GET_GROWING_LOCATION_BASE_SQL + `
			WHERE growing_locations.id = $1
		`

	GET_ALL_GROWING_LOCATIONS_SQL = GET_GROWING_LOCATION_BASE_SQL + `
		WHERE growing_locations.growing_group = $1
	`
)

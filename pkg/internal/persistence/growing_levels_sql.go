package persistence

const (
	CREATE_GROWING_LEVEL_SQL = `
		INSERT INTO growing_levels
		(display_name, growing_location)
		VALUES ($1, $2)
		RETURNING id
		`

	DELETE_GROWING_LEVEL_SQL = `
		UPDATE growing_levels SET archived = NOW()
		WHERE growing_levels.id = $1
		`

	GET_GROWING_LEVEL_BASE_SQL = `
		SELECT
			growing_levels.id AS id,
			growing_levels.growing_location AS growing_location_id,
			growing_levels.display_name AS display_name,
			growing_levels.created AS created_at,
			growing_levels.updated AS updated_at,
			cre_member.id AS created_member_id,
			cre_member.display_name AS created_member_name,
			up_member.id AS updated_member_id,
			up_member.display_name AS updated_member_name
		FROM growing_levels
		LEFT JOIN members AS cre_member ON members.id = growing_levels.created_by
		LEFT JOIN members AS up_member ON members.id = growing_levels.updated_by
		`

	GET_GROWING_LEVEL_BY_ID_SQL = GET_GROWING_LEVEL_BASE_SQL + `
	WHERE growing_levels.id = $1
	`

	GET_ALL_GROWING_LEVELS_BY_GROWING_LOCATION_ID_SQL = GET_GROWING_LEVEL_BASE_SQL + `
	WHERE growing_levels.growing_location = $1
	`
)

package persistence

const (
	CREATE_SEED_SQL = `
		INSERT INTO growing_mediums
		(display_name)
		VALUES ($1)
		RETURNING id
		`

	DELETE_SEED_SQL = `
		UPDATE growing_mediums SET archived = NOW()
		WHERE growing_mediums.id = $1
		`

	GET_SEED_BASE_SQL = `
		SELECT
			growing_mediums.id AS id,
			growing_mediums.display_name AS display_name,
			growing_mediums.created AS created_at,
			growing_mediums.updated AS updated_at,
			cre_member.id AS created_member_id,
			cre_member.display_name AS created_member_name,
			up_member.id AS updated_member_id,
			up_member.display_name AS updated_member_name
		FROM growing_mediums
		LEFT JOIN members AS cre_member ON members.id = growing_mediums.created_by
		LEFT JOIN members AS up_member ON members.id = growing_mediums.updated_by
		`

	GET_SEED_BASE_BY_ID_SQL = GET_SEED_BASE_SQL + `
		WHERE growing_mediums.id = $1
		`

	GET_ALL_SEEDS_SQL = GET_SEED_BASE_SQL
)

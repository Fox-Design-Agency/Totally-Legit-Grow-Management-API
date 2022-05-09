package persistence

const (
	CREATE_MEMBER_SQL = `
		INSERT INTO members
		(display_name)
		VALUES ($1)
		RETURNING id
		`

	DELETE_MEMBER_SQL = `
		UPDATE members SET archived = NOW()
		WHERE members.id = $1
		`

	GET_MEMBER_BASE_SQL = `
		SELECT
			members.id AS id,
			members.display_name AS display_name,
			members.created AS created_at,
			members.updated AS updated_at,
			cre_member.id AS created_member_id,
			cre_member.display_name AS created_member_name,
			up_member.id AS updated_member_id,
			up_member.display_name AS updated_member_name
		FROM members
		LEFT JOIN members AS cre_member ON members.id = members.created_by
		LEFT JOIN members AS up_member ON members.id = members.updated_by
		`

	GET_MEMBER_BY_ID_SQL = GET_MEMBER_BASE_SQL + `
		WHERE members.id = $1
		`

	GET_ALL_MEMBERS_SQL = GET_MEMBER_BASE_SQL
)

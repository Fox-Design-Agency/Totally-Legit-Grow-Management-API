// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package nutrients

const (
	CREATE_NUTRIENT_SQL = `
		INSERT INTO nutrients
		(display_name)
		VALUES ($1)
		RETURNING id
		`

	DELETE_NUTRIENT_SQL = `
		UPDATE nutrients SET archived = NOW()
		WHERE nutrients.id = $1
		`

	GET_NUTRIENT_BASE_SQL = `
		SELECT
			nutrients.id AS id,
			nutrients.display_name AS display_name,
			nutrients.created AS created_at,
			nutrients.updated AS updated_at,
			cre_member.id AS created_member_id,
			cre_member.display_name AS created_member_name,
			up_member.id AS updated_member_id,
			up_member.display_name AS updated_member_name
		FROM nutrients
		LEFT JOIN members AS cre_member ON members.id = nutrients.created_by
		LEFT JOIN members AS up_member ON members.id = nutrients.updated_by
		`

	GET_NUTRIENT_BY_ID_SQL = GET_NUTRIENT_BASE_SQL + `
		WHERE nutrients.id = $1
	`

	GET_ALL_NUTRIENTS_SQL = GET_NUTRIENT_BASE_SQL
)

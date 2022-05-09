// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package persistence

const (
	CREATE_GROW_SPOT_SQL = `
		INSERT INTO grow_spots
		(display_name, growing_level)
		VALUES ($1, $2)
		RETURNING id
		`

	DELETE_GROW_SPOT_SQL = `
		UPDATE grow_spots SET archived = NOW()
		WHERE grow_spots.id = $1
		`

	GET_GROW_SPOT_BASE_SQL = `
		SELECT
			grow_spots.id AS id,
			grow_spots.growing_level AS growing_level_id,
			grow_spots.display_name AS display_name,
			grow_spots.created AS created_at,
			grow_spots.updated AS updated_at,
			cre_member.id AS created_member_id,
			cre_member.display_name AS created_member_name,
			up_member.id AS updated_member_id,
			up_member.display_name AS updated_member_name
		FROM grow_spots
		LEFT JOIN members AS cre_member ON members.id = grow_spots.created_by
		LEFT JOIN members AS up_member ON members.id = grow_spots.updated_by
		`

	GET_GROW_SPOT_BY_ID_SQL = GET_GROW_SPOT_BASE_SQL + `
		WHERE grow_spots.id = $1
		`

	GET_ALL_GROW_SPOTS_BY_GROW_LEVEL_ID_SQL = GET_GROW_SPOT_BASE_SQL + `
		WHERE grow_spots.growing_level = $1
		`
)

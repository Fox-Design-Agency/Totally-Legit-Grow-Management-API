// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package persistence

const (
	CREATE_GROWING_GROUPS_SQL = `
		INSERT INTO growing_groups
		(display_name, organization)
		VALUES ($1, $2)
		RETURNING id
		`

	DELETE_GROWING_GROUP_SQL = `
		UPDATE growing_groups SET archived = NOW()
		WHERE growing_groups.id = $1
		`

	GET_GROWING_GROUP_BASE_SQL = `
		SELECT
		growing_groups.id AS id,
		growing_groups.organization AS organization_id,
		growing_groups.display_name AS display_name,
		growing_groups.created AS created_at,
		growing_groups.updated AS updated_at,
			cre_member.id AS created_member_id,
			cre_member.display_name AS created_member_name,
			up_member.id AS updated_member_id,
			up_member.display_name AS updated_member_name
		FROM growing_groups
		LEFT JOIN members AS cre_member ON members.id = growing_groups.created_by
		LEFT JOIN members AS up_member ON members.id = growing_groups.updated_by
		`

	GET_GROWING_GROUP_BY_ID_SQL = GET_GROWING_GROUP_BASE_SQL + `
	WHERE growing_groups.id = $1
	`

	GET_ALL_GROWING_GROUPS_BY_ORGANIZATION_ID = GET_GROWING_GROUP_BASE_SQL + `
	WHERE growing_groups.organization = $1
	`
)

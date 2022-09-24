// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package organizations

const (
	CREATE_ORGANIZATION_SQL = `
		INSERT INTO organizations
		(display_name)
		VALUES ($1)
		RETURNING id
		`

	DELETE_ORGANIZATION_SQL = `
		UPDATE organizations SET archived = NOW()
		WHERE organizations.id = $1
		`

	GET_ORGANIZATION_BASE_SQL = `
		SELECT
			organizations.id AS id,
			organizations.display_name AS display_name,
			organizations.created AS created_at,
			organizations.updated AS updated_at,
			cre_member.id AS created_member_id,
			cre_member.display_name AS created_member_name,
			up_member.id AS updated_member_id,
			up_member.display_name AS updated_member_name
		FROM organizations
		LEFT JOIN members AS cre_member ON members.id = organizations.created_by
		LEFT JOIN members AS up_member ON members.id = organizations.updated_by
		`

	GET_ORGANIZATION_BY_ID_SQL = GET_ORGANIZATION_BASE_SQL + `
		WHERE organizations.id = $1
	`

	GET_ALL_ORGANIZATIONS_SQL = GET_ORGANIZATION_BASE_SQL
)

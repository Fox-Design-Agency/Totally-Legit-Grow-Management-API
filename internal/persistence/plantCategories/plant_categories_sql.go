// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package plantcategories

const (
	CREATE_PLANT_CATEGORY_SQL = `
		INSERT INTO plant_categories
		(display_name)
		VALUES ($1)
		RETURNING id
		`

	DELETE_PLANT_CATEGORY_SQL = `
		UPDATE plant_categories SET archived = NOW()
		WHERE plant_categories.id = $1
		`

	GET_PLANT_CATEGORY_BASE_SQL = `
		SELECT
			plant_categories.id AS id,
			plant_categories.display_name AS display_name,
			plant_categories.created AS created_at,
			plant_categories.updated AS updated_at,
			cre_member.id AS created_member_id,
			cre_member.display_name AS created_member_name,
			up_member.id AS updated_member_id,
			up_member.display_name AS updated_member_name
		FROM plant_categories
		LEFT JOIN members AS cre_member ON members.id = plant_categories.created_by
		LEFT JOIN members AS up_member ON members.id = plant_categories.updated_by
		`

	GET_PLANT_CATEGORY_BY_ID_SQL = GET_PLANT_CATEGORY_BASE_SQL + `
	WHERE plant_categories.id = $1
	`

	GET_ALL_PLANT_CATEGORY_SQL = GET_PLANT_CATEGORY_BASE_SQL
)

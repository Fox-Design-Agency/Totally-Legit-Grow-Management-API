// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package persistence

const (
	CREATE_GROW_SPOT_PLANT_SQL = `
		INSERT INTO grow_spot_plants
		(plant, plant_life_cycle, grow_spot, seed, seed_amount_units, seed_amount)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
		`

	DELETE_GROW_SPOT_PLANT_SQL = `
		UPDATE grow_spot_plants SET archived = NOW()
		WHERE grow_spot_plants.id = $1
		`

	GET_GROW_SPOT_PLANT_BASE_SQL = `
		SELECT
			grow_spot_plants.id AS id,
			grow_spot_plants.created AS created_at,
			grow_spot_plants.updated AS updated_at,
			cre_member.id AS created_member_id,
			cre_member.display_name AS created_member_name,
			up_member.id AS updated_member_id,
			up_member.display_name AS updated_member_name
		FROM grow_spot_plants
		LEFT JOIN members AS cre_member ON members.id = grow_spot_plants.created_by
		LEFT JOIN members AS up_member ON members.id = grow_spot_plants.updated_by
		`

	GET_GROW_SPOT_PLANT_BY_ID_SQL = GET_GROW_SPOT_PLANT_BASE_SQL + `
		WHERE grow_spot_plants.id = $1`

	GET_ALL_GROW_SPOT_PLANTS_SQL = GET_GROW_SPOT_PLANT_BASE_SQL
)

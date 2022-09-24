// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package plantlifecycle

const (
	CREATE_PLANT_LIFE_CYCLE_SQL = `
		INSERT INTO plant_life_cycles
		(display_name, total_time_measure_units, est_total_time_measure)
		VALUES ($1, $2, $3)
		RETURNING id
		`

	DELETE_PLANT_LIFE_CYCLE_SQL = `
		UPDATE plant_life_cycles SET archived = NOW()
		WHERE plant_life_cycles.id = $1
		`

	GET_PLANT_LIFE_CYCLE_BASE_SQL = `
		SELECT
			plant_life_cycles.id AS id,
			plant_life_cycles.actual_time_measure AS actual_time_measure,
			plant_life_cycles.display_name AS display_name,
			plant_life_cycles.total_time_measure_units AS total_time_measure_units,
			plant_life_cycles.est_total_time_measure AS est_total_time_measure,
			plant_life_cycles.created AS created_at,
			plant_life_cycles.updated AS updated_at,
			cre_member.id AS created_member_id,
			cre_member.display_name AS created_member_name,
			up_member.id AS updated_member_id,
			up_member.display_name AS updated_member_name
		FROM plant_life_cycles
		LEFT JOIN members AS cre_member ON members.id = plant_life_cycles.created_by
		LEFT JOIN members AS up_member ON members.id = plant_life_cycles.updated_by
		WHERE plant_life_cycles.id = $1
		`

	GET_PLANT_LIFE_CYCLE_BY_ID_SQL = GET_PLANT_LIFE_CYCLE_BASE_SQL + `
		WHERE plant_life_cycles.id = $1
		`

	GET_ALL_PLANT_LIFE_CYCLES_SQL = GET_PLANT_LIFE_CYCLE_BASE_SQL
)

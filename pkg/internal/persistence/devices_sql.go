// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package persistence

const (
	CREATE_DEVICE_SQL = `
		INSERT INTO devices
		(display_name)
		VALUES ($1)
		RETURNING id
		`

	CREATE_DEVICE_ACTIONS_SQL = `
		INSERT INTO device_actions
		(device, display_name, invoke_route, info)
		VALUES ($1, $2, $3, $4)
		`

	CREATE_GROWING_GROUP_DEVICES_SQL = `
		INSERT INTO growing_group_devices
		(device, growing_group)
		VALUES ($1, $2)
		`

	CREATE_GROWING_LOCATION_DEVICES_SQL = `
		INSERT INTO growing_location_devices
		(device, growing_location)
		VALUES ($1, $2)
		`

	CREATE_GROWING_LEVEL_DEVICES_SQL = `
		INSERT INTO growing_level_devices
		(device, growing_level)
		VALUES ($1, $2)
		`

	CREATE_GROW_SPOT_DEVICES_SQL = `
		INSERT INTO grow_spots_devices
		(device, grow_spot)
		VALUES ($1, $2)
		`

	DELETE_DEVICE_SQL = `
		UPDATE devices SET archived = NOW()
		WHERE devices.id = $1
		`

	DELETE_GROWING_GROUP_DEVICE_SQL = `
		UPDATE growing_group_devices SET archived = NOW()
		WHERE growing_group_devices.device = $1
			AND growing_group_devices.growing_group = $2
		`

	DELETE_GROWING_LOCATION_DEVICE_SQL = `
		UPDATE growing_location_devices SET archived = NOW()
		WHERE growing_location_devices.device = $1
			AND growing_location_devices.growing_location = $2
		`

	DELETE_GROWING_LEVEL_DEVICE_SQL = `
		UPDATE growing_level_devices SET archived = NOW()
		WHERE growing_level_devices.device = $1
			AND growing_level_devices.growing_location = $2
		`

	DELETE_GROW_SPOT_DEVICE_SQL = `
		UPDATE grow_spots_devices SET archived = NOW()
		WHERE grow_spots_devices.device = $1
			AND grow_spots_devices.growing_location = $2
		`

	GET_DEVICE_BASE_SQL = `
		SELECT
			devices.id AS id,
			devices.display_name AS display_name,
			devices.created AS created_at,
			devices.updated AS updated_at,
			cre_member.id AS created_member_id,
			cre_member.display_name AS created_member_name,
			up_member.id AS updated_member_id,
			up_member.display_name AS updated_member_name
		FROM devices
		LEFT JOIN members AS cre_member ON members.id = devices.created_by
		LEFT JOIN members AS up_member ON members.id = devices.updated_by
		`

	GET_DEVICE_BY_ID_SQL = GET_DEVICE_BASE_SQL + `
			WHERE devices.id = $1
			`

	GET_DEVICE_ACTIONS_SQL = `
		SELECT
			device_actions.id AS id,
			device_actions.display_name AS display_name,
			device_actions.invoke_route AS invoke_route,
			device_actions.info AS action_info,
			device_actions.created AS created_at,
			device_actions.updated AS updated_at,
			cre_member.id AS created_member_id,
			cre_member.display_name AS created_member_name,
			up_member.id AS updated_member_id,
			up_member.display_name AS updated_member_name
		FROM device_actions
		JOIN devices ON devices.id = device_actions.device
		LEFT JOIN members AS cre_member ON members.id = device_actions.created_by
		LEFT JOIN members AS up_member ON members.id = device_actions.updated_by
		WHERE device_actions.device = $1
		`

	GET_ALL_DEVICES_SQL = GET_DEVICE_BASE_SQL

	GET_GROWING_GROUP_DEVICE_BASE_SQL = `
	SELECT
		devices.id AS id,
		devices.display_name AS display_name,
		devices.created AS created_at,
		devices.updated AS updated_at,
		cre_member.id AS created_member_id,
		cre_member.display_name AS created_member_name,
		up_member.id AS updated_member_id,
		up_member.display_name AS updated_member_name
	FROM growing_group_devices
	JOIN devices ON devices.device = growing_group_devices.device
	LEFT JOIN members AS cre_member ON members.id = growing_group_devices.created_by
	LEFT JOIN members AS up_member ON members.id = growing_group_devices.updated_by
	`

	GET_GROWING_GROUP_DEVICE_SQL = GET_GROWING_GROUP_DEVICE_BASE_SQL + ` 
	WHERE growing_group_devices.growing_group = $1
	AND growing_group_devices.device = $2
	`

	GET_ALL_GROWING_GROUP_DEVICES_SQL = GET_GROWING_GROUP_DEVICE_BASE_SQL +
		`
		WHERE growing_group_devices.growing_group = $1
		`
	GET_GROWING_LOCATION_DEVICE_BASE_SQL = `
	SELECT
		devices.id AS id,
		devices.display_name AS display_name,
		devices.created AS created_at,
		devices.updated AS updated_at,
		cre_member.id AS created_member_id,
		cre_member.display_name AS created_member_name,
		up_member.id AS updated_member_id,
		up_member.display_name AS updated_member_name
	FROM growing_location_devices
	JOIN devices ON devices.device = growing_location_devices.device
	LEFT JOIN members AS cre_member ON members.id = growing_location_devices.created_by
	LEFT JOIN members AS up_member ON members.id = growing_location_devices.updated_by
	WHERE growing_location_devices.growing_group = $1
		AND growing_location_devices.device = $2
	`

	GET_GROWING_LOCATION_DEVICE_SQL = GET_GROWING_LOCATION_DEVICE_BASE_SQL + `
	WHERE growing_location_devices.growing_group = $1
		AND growing_location_devices.device = $2
	`

	GET_ALL_GROWING_LOCATION_DEVICES_SQL = GET_GROWING_LOCATION_DEVICE_BASE_SQL + `
		WHERE growing_location_devices.growing_group = $1
		`

	GET_GROWING_LEVEL_DEVICE_BASE_SQL = `
	SELECT
		devices.id AS id,
		devices.display_name AS display_name,
		devices.created AS created_at,
		devices.updated AS updated_at,
		cre_member.id AS created_member_id,
		cre_member.display_name AS created_member_name,
		up_member.id AS updated_member_id,
		up_member.display_name AS updated_member_name
	FROM growing_level_devices
	JOIN devices ON devices.device = growing_level_devices.device
	LEFT JOIN members AS cre_member ON members.id = growing_level_devices.created_by
	LEFT JOIN members AS up_member ON members.id = growing_level_devices.updated_by
	`

	GET_GROWING_LEVEL_DEVICE_SQL = GET_GROWING_LEVEL_DEVICE_BASE_SQL + `
	WHERE growing_level_devices.growing_group = $1
		AND growing_level_devices.device = $2
	`

	GET_ALL_GROWING_LEVEL_DEVICES_SQL = GET_GROWING_LEVEL_DEVICE_BASE_SQL + `
		WHERE growing_level_devices.growing_group = $1
		`

	GET_GROWING_SPOT_BASE_SQL = `
	SELECT
		devices.id AS id,
		devices.display_name AS display_name,
		devices.created AS created_at,
		devices.updated AS updated_at,
		cre_member.id AS created_member_id,
		cre_member.display_name AS created_member_name,
		up_member.id AS updated_member_id,
		up_member.display_name AS updated_member_name
	FROM grow_spots_devices
	JOIN devices ON devices.device = grow_spots_devices.device
	LEFT JOIN members AS cre_member ON members.id = grow_spots_devices.created_by
	LEFT JOIN members AS up_member ON members.id = grow_spots_devices.updated_by
	`

	GET_GROWING_SPOT_SQL = GET_GROWING_SPOT_BASE_SQL +
		`
	WHERE grow_spots_devices.growing_group = $1
		AND grow_spots_devices.device = $2
	`

	GET_ALL_GROWING_SPOT_DEVICES_SQL = GET_GROWING_SPOT_BASE_SQL +
		`
		WHERE grow_spots_devices.growing_group = $1
		`
)

package persistence

import (
	routemodels "smart-grow-management-api/v1/pkg/internal/route-models"

	"github.com/jmoiron/sqlx"
)

var _devices IDevicesDB = &Persistence{}

func (db *Persistence) CreateDevice(req routemodels.CreateDeviceRequest) (*routemodels.CreateDeviceResponse, error) {
	var result string

	SQL := `
	INSERT INTO devices
	(display_name)
	VALUES ($1)
	RETURNING id
	`

	// Make the appropiate SQL Call
	if err := db.Postgres.QueryRow(SQL, req.DisplayName).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.CreateDeviceResponse{
		ID: result,
	}, nil
}

func (db *Persistence) CreateDeviceWithTransaction(tx *sqlx.Tx, req routemodels.CreateDeviceRequest) (*routemodels.CreateDeviceResponse, error) {
	var result string

	SQL := `
	INSERT INTO devices
	(display_name)
	VALUES ($1)
	RETURNING id
	`

	// Make the appropiate SQL Call
	if err := tx.QueryRow(SQL, req.DisplayName).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.CreateDeviceResponse{
		ID: result,
	}, nil
}

func (db *Persistence) CreateDeviceAction(req routemodels.CreateDeviceActionRequest) (*routemodels.CreateDeviceActionResponse, error) {
	var result string

	SQL := `
	INSERT INTO device_actions
	(device, display_name, invoke_route, info)
	VALUES ($1, $2, $3, $4)
	`

	args := []interface{}{
		req.ID,
		req.DisplayName,
		req.InvokeRoute,
		req.Information,
	}

	// Make the appropiate SQL Call
	if _, err := db.Postgres.Exec(SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.CreateDeviceActionResponse{
		ID: result,
	}, nil
}

func (db *Persistence) CreateDeviceActionWithTransaction(tx *sqlx.Tx, req routemodels.CreateDeviceActionRequest) (*routemodels.CreateDeviceActionResponse, error) {
	var result string

	SQL := `
	INSERT INTO device_actions
	(device, display_name, invoke_route, info)
	VALUES ($1, $2, $3, $4)
	`

	args := []interface{}{
		req.ID,
		req.DisplayName,
		req.InvokeRoute,
		req.Information,
	}

	// Make the appropiate SQL Call
	if _, err := tx.Exec(SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.CreateDeviceActionResponse{
		ID: result,
	}, nil
}

func (db *Persistence) CreateGrowingGroupDevice(req routemodels.CreateGrowingGroupDeviceRequest) (*routemodels.CreateGrowingGroupDeviceResponse, error) {

	SQL := `
	INSERT INTO growing_group_devices
	(device, growing_group)
	VALUES ($1, $2)
	`

	args := []interface{}{
		req.DeviceID,
		req.GrowingGroupID,
	}

	// Make the appropiate SQL Call
	if _, err := db.Postgres.Exec(SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.CreateGrowingGroupDeviceResponse{
		ID: "",
	}, nil
}

func (db *Persistence) CreateGrowingGroupDeviceWithTransaction(tx *sqlx.Tx, req routemodels.CreateGrowingGroupDeviceRequest) (*routemodels.CreateGrowingGroupDeviceResponse, error) {

	SQL := `
	INSERT INTO growing_group_devices
	(device, growing_group)
	VALUES ($1, $2)
	`

	args := []interface{}{
		req.DeviceID,
		req.GrowingGroupID,
	}

	// Make the appropiate SQL Call
	if _, err := tx.Exec(SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.CreateGrowingGroupDeviceResponse{
		ID: "",
	}, nil
}

func (db *Persistence) CreateGrowingLocationDevice(req routemodels.CreateGrowingLocationDeviceRequest) (*routemodels.CreateGrowingLocationDeviceResponse, error) {

	SQL := `
	INSERT INTO growing_location_devices
	(device, growing_location)
	VALUES ($1, $2)
	`

	args := []interface{}{
		req.DeviceID,
		req.GrowingLocationID,
	}

	// Make the appropiate SQL Call
	if _, err := db.Postgres.Exec(SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.CreateGrowingLocationDeviceResponse{
		ID: "",
	}, nil
}

func (db *Persistence) CreateGrowingLocationDeviceWithTransaction(tx *sqlx.Tx, req routemodels.CreateGrowingLocationDeviceRequest) (*routemodels.CreateGrowingLocationDeviceResponse, error) {

	SQL := `
	INSERT INTO growing_location_devices
	(device, growing_location)
	VALUES ($1, $2)
	`

	args := []interface{}{
		req.DeviceID,
		req.GrowingLocationID,
	}

	// Make the appropiate SQL Call
	if _, err := db.Postgres.Exec(SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.CreateGrowingLocationDeviceResponse{
		ID: "",
	}, nil
}

func (db *Persistence) CreateGrowingLevelDevice(req routemodels.CreateGrowingLevelDeviceRequest) (*routemodels.CreateGrowingLevelDeviceResponse, error) {

	SQL := `
	INSERT INTO growing_level_devices
	(device, growing_level)
	VALUES ($1, $2)
	`

	args := []interface{}{
		req.DeviceID,
		req.GrowingLevelID,
	}

	// Make the appropiate SQL Call
	if _, err := db.Postgres.Exec(SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.CreateGrowingLevelDeviceResponse{
		ID: "",
	}, nil
}

func (db *Persistence) CreateGrowingLevelDeviceWithTransaction(sql *sqlx.Tx, req routemodels.CreateGrowingLevelDeviceRequest) (*routemodels.CreateGrowingLevelDeviceResponse, error) {

	SQL := `
	INSERT INTO growing_level_devices
	(device, growing_level)
	VALUES ($1, $2)
	`

	args := []interface{}{
		req.DeviceID,
		req.GrowingLevelID,
	}

	// Make the appropiate SQL Call
	if _, err := db.Postgres.Exec(SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.CreateGrowingLevelDeviceResponse{
		ID: "",
	}, nil
}

func (db *Persistence) CreateGrowingSpotDevice(req routemodels.CreateGrowingSpotDeviceRequest) (*routemodels.CreateGrowingSpotDeviceResponse, error) {

	SQL := `
	INSERT INTO grow_spots_devices
	(device, grow_spot)
	VALUES ($1, $2)
	`

	args := []interface{}{
		req.DeviceID,
		req.GrowingLevelID,
	}

	// Make the appropiate SQL Call
	if _, err := db.Postgres.Exec(SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.CreateGrowingSpotDeviceResponse{
		ID: "",
	}, nil
}

func (db *Persistence) CreateGrowingSpotDeviceWithTransaction(tx *sqlx.Tx, req routemodels.CreateGrowingSpotDeviceRequest) (*routemodels.CreateGrowingSpotDeviceResponse, error) {

	SQL := `
	INSERT INTO grow_spots_devices
	(device, grow_spot)
	VALUES ($1, $2)
	`

	args := []interface{}{
		req.DeviceID,
		req.GrowingLevelID,
	}

	// Make the appropiate SQL Call
	if _, err := db.Postgres.Exec(SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.CreateGrowingSpotDeviceResponse{
		ID: "",
	}, nil
}

func (db *Persistence) DeleteDevice(req routemodels.DeleteDeviceRequest) error {
	var result string

	SQL := `
	UPDATE devices SET archived = NOW()
	WHERE devices.id = $1
	`

	// Make the appropiate SQL Call
	if err := db.Postgres.QueryRow(SQL, req.ID).Scan(result); err != nil {
		// handle err
		return err
	}

	return nil
}

func (db *Persistence) DeleteGrowingGroupDevice(req routemodels.DeleteGrowingGroupDeviceRequest) error {
	var result string

	SQL := `
	UPDATE growing_group_devices SET archived = NOW()
	WHERE growing_group_devices.device = $1
		AND growing_group_devices.growing_group = $2
	`
	args := []interface{}{
		req.DeviceID,
		req.GrowingGroupID,
	}

	// Make the appropiate SQL Call
	if err := db.Postgres.QueryRow(SQL, args...).Scan(result); err != nil {
		// handle err
		return err
	}

	return nil
}

func (db *Persistence) DeleteGrowingLocationDevice(req routemodels.DeleteGrowingLocationDeviceRequest) error {
	var result string

	SQL := `
	UPDATE growing_location_devices SET archived = NOW()
	WHERE growing_location_devices.device = $1
		AND growing_location_devices.growing_location = $2
	`
	args := []interface{}{
		req.DeviceID,
		req.GrowingLocationID,
	}

	// Make the appropiate SQL Call
	if err := db.Postgres.QueryRow(SQL, args...).Scan(result); err != nil {
		// handle err
		return err
	}

	return nil
}

func (db *Persistence) DeleteGrowingLevelDevice(req routemodels.DeleteGrowingLevelDeviceRequest) error {
	var result string

	SQL := `
	UPDATE growing_level_devices SET archived = NOW()
	WHERE growing_level_devices.device = $1
		AND growing_level_devices.growing_location = $2
	`
	args := []interface{}{
		req.DeviceID,
		req.GrowingLevelID,
	}

	// Make the appropiate SQL Call
	if err := db.Postgres.QueryRow(SQL, args...).Scan(result); err != nil {
		// handle err
		return err
	}

	return nil
}

func (db *Persistence) DeleteGrowingSpotDevice(req routemodels.DeleteGrowingSpotDeviceRequest) error {
	var result string

	SQL := `
	UPDATE grow_spots_devices SET archived = NOW()
	WHERE grow_spots_devices.device = $1
		AND grow_spots_devices.growing_location = $2
	`
	args := []interface{}{
		req.DeviceID,
		req.GrowingSpotID,
	}

	// Make the appropiate SQL Call
	if err := db.Postgres.QueryRow(SQL, args...).Scan(result); err != nil {
		// handle err
		return err
	}

	return nil
}

func (db *Persistence) EditDevice(req routemodels.EditDeviceRequest) (*routemodels.EditDeviceResponse, error) {
	return nil, nil
}

func (db *Persistence) GetDevice(req routemodels.GetDeviceRequest) (*routemodels.GetDeviceResponse, error) {
	var result routemodels.Device

	SQL := `
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
	WHERE devices.id = $1
	`

	// Make the appropiate SQL Call
	if err := db.Postgres.Get(&result, SQL, req.ID); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetDeviceResponse{
		Device: result,
	}, nil
}

func (db *Persistence) GetDeviceActions(req routemodels.GetDeviceActionsRequest) (*routemodels.GetDeviceActionsResponse, error) {
	var result []routemodels.DeviceAction

	SQL := `
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

	// Make the appropiate SQL Call
	if err := db.Postgres.Select(&result, SQL, req.DeviceID); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetDeviceActionsResponse{
		Actions: result,
	}, nil
}

func (db *Persistence) GetDeviceActionsWithTransaction(tx *sqlx.Tx, req routemodels.GetDeviceActionsRequest) (*routemodels.GetDeviceActionsResponse, error) {
	var result []routemodels.DeviceAction

	SQL := `
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

	// Make the appropiate SQL Call
	if err := tx.Select(&result, SQL, req.DeviceID); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetDeviceActionsResponse{
		Actions: result,
	}, nil
}

func (db *Persistence) GetAllDevices(req routemodels.GetAllDevicesRequest) (*routemodels.GetAllDevicesResponse, error) {
	var result []routemodels.Device

	SQL := `
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

	// Make the appropiate SQL Call
	if err := db.Postgres.Select(&result, SQL); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetAllDevicesResponse{
		Devices: result,
	}, nil
}

func (db *Persistence) GetAllGrowingGroupDevices(req routemodels.GetAllGrowingGroupDevicesRequest) (*routemodels.GetAllGrowingGroupDevicesResponse, error) {
	var result []routemodels.Device

	SQL := `
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
	WHERE growing_group_devices.growing_group = $1
	`
	args := []interface{}{
		req.GrowingGroupID,
	}

	// Make the appropiate SQL Call
	if err := db.Postgres.Select(&result, SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetAllGrowingGroupDevicesResponse{
		Devices: result,
	}, nil
}
func (db *Persistence) GetAllGrowingGroupDeviceWithTransaction(tx *sqlx.Tx, req routemodels.GetAllGrowingGroupDevicesRequest) (*routemodels.GetAllGrowingGroupDevicesResponse, error) {
	var result []routemodels.Device

	SQL := `
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
	WHERE growing_group_devices.growing_group = $1
	`
	args := []interface{}{
		req.GrowingGroupID,
	}

	// Make the appropiate SQL Call
	if err := tx.Select(&result, SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetAllGrowingGroupDevicesResponse{
		Devices: result,
	}, nil
}

func (db *Persistence) GetAllGrowingLocationDevices(req routemodels.GetAllGrowingLocationDevicesRequest) (*routemodels.GetAllGrowingLocationDevicesResponse, error) {
	var result []routemodels.Device

	SQL := `
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
	`
	args := []interface{}{
		req.GrowingLocationID,
	}

	// Make the appropiate SQL Call
	if err := db.Postgres.Select(&result, SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetAllGrowingLocationDevicesResponse{
		Devices: result,
	}, nil
}

func (db *Persistence) GetAllGrowingLocationDeviceWithTransaction(tx *sqlx.Tx, req routemodels.GetAllGrowingLocationDevicesRequest) (*routemodels.GetAllGrowingLocationDevicesResponse, error) {
	var result []routemodels.Device

	SQL := `
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
	`
	args := []interface{}{
		req.GrowingLocationID,
	}

	// Make the appropiate SQL Call
	if err := tx.Select(&result, SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetAllGrowingLocationDevicesResponse{
		Devices: result,
	}, nil
}

func (db *Persistence) GetAllGrowingLevelDevices(req routemodels.GetAllGrowingLevelDevicesRequest) (*routemodels.GetAllGrowingLevelDevicesResponse, error) {
	var result []routemodels.Device

	SQL := `
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
	WHERE growing_level_devices.growing_group = $1
	`
	args := []interface{}{
		req.GrowingLevelID,
	}

	// Make the appropiate SQL Call
	if err := db.Postgres.Get(&result, SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetAllGrowingLevelDevicesResponse{
		Devices: result,
	}, nil
}

func (db *Persistence) GetAllGrowingLevelDeviceWithTransaction(tx *sqlx.Tx, req routemodels.GetAllGrowingLevelDevicesRequest) (*routemodels.GetAllGrowingLevelDevicesResponse, error) {
	var result []routemodels.Device

	SQL := `
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
	WHERE growing_level_devices.growing_group = $1
		AND growing_level_devices.device = $2
	`
	args := []interface{}{
		req.GrowingLevelID,
	}

	// Make the appropiate SQL Call
	if err := tx.Select(&result, SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetAllGrowingLevelDevicesResponse{
		Devices: result,
	}, nil
}

func (db *Persistence) GetAllGrowingSpotDevices(req routemodels.GetAllGrowingSpotDevicesRequest) (*routemodels.GetAllGrowingSpotDevicesResponse, error) {
	var result []routemodels.Device

	SQL := `
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
	WHERE grow_spots_devices.growing_group = $1
		AND grow_spots_devices.device = $2
	`
	args := []interface{}{
		req.GrowingSpotID,
	}

	// Make the appropiate SQL Call
	if err := db.Postgres.Select(&result, SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetAllGrowingSpotDevicesResponse{
		Devices: result,
	}, nil
}

func (db *Persistence) GetAllGrowingSpotDeviceWithTransaction(tx *sqlx.Tx, req routemodels.GetAllGrowingSpotDevicesRequest) (*routemodels.GetAllGrowingSpotDevicesResponse, error) {
	var result []routemodels.Device

	SQL := `
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
	WHERE grow_spots_devices.growing_group = $1
		AND grow_spots_devices.device = $2
	`
	args := []interface{}{
		req.GrowingSpotID,
	}

	// Make the appropiate SQL Call
	if err := tx.Select(&result, SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetAllGrowingSpotDevicesResponse{
		Devices: result,
	}, nil
}

func (db *Persistence) GetGrowingGroupDevice(req routemodels.GetGrowingGroupDeviceRequest) (*routemodels.GetGrowingGroupDeviceResponse, error) {
	var result routemodels.Device

	SQL := `
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
	WHERE growing_group_devices.growing_group = $1
		AND growing_group_devices.device = $2
	`
	args := []interface{}{
		req.GrowingGroupID,
		req.DeviceID,
	}

	// Make the appropiate SQL Call
	if err := db.Postgres.Get(&result, SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetGrowingGroupDeviceResponse{
		Device: result,
	}, nil
}

func (db *Persistence) GetGrowingGroupDeviceWithTransaction(tx *sqlx.Tx, req routemodels.GetGrowingGroupDeviceRequest) (*routemodels.GetGrowingGroupDeviceResponse, error) {
	var result routemodels.Device

	SQL := `
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
	WHERE growing_group_devices.growing_group = $1
		AND growing_group_devices.device = $2
	`
	args := []interface{}{
		req.GrowingGroupID,
		req.DeviceID,
	}

	// Make the appropiate SQL Call
	if err := tx.Get(&result, SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetGrowingGroupDeviceResponse{
		Device: result,
	}, nil
}

func (db *Persistence) GetGrowingLocationDevice(req routemodels.GetGrowingLocationDeviceRequest) (*routemodels.GetGrowingLocationDeviceResponse, error) {
	var result routemodels.Device

	SQL := `
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
	args := []interface{}{
		req.GrowingLocationID,
		req.DeviceID,
	}

	// Make the appropiate SQL Call
	if err := db.Postgres.Get(&result, SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetGrowingLocationDeviceResponse{
		Device: result,
	}, nil
}

func (db *Persistence) GetGrowingLocationDeviceWithTransaction(tx *sqlx.Tx, req routemodels.GetGrowingLocationDeviceRequest) (*routemodels.GetGrowingLocationDeviceResponse, error) {
	var result routemodels.Device

	SQL := `
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
	args := []interface{}{
		req.GrowingLocationID,
		req.DeviceID,
	}

	// Make the appropiate SQL Call
	if err := db.Postgres.Get(&result, SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetGrowingLocationDeviceResponse{
		Device: result,
	}, nil
}

func (db *Persistence) GetGrowingLevelDevice(req routemodels.GetGrowingLevelDeviceRequest) (*routemodels.GetGrowingLevelDeviceResponse, error) {
	var result routemodels.Device

	SQL := `
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
	WHERE growing_level_devices.growing_group = $1
		AND growing_level_devices.device = $2
	`
	args := []interface{}{
		req.GrowingLevelID,
		req.DeviceID,
	}

	// Make the appropiate SQL Call
	if err := db.Postgres.Get(&result, SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetGrowingLevelDeviceResponse{
		Device: result,
	}, nil
}

func (db *Persistence) GetGrowingLevelDeviceWithTransaction(tx *sqlx.Tx, req routemodels.GetGrowingLevelDeviceRequest) (*routemodels.GetGrowingLevelDeviceResponse, error) {
	var result routemodels.Device

	SQL := `
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
	WHERE growing_level_devices.growing_group = $1
		AND growing_level_devices.device = $2
	`
	args := []interface{}{
		req.GrowingLevelID,
		req.DeviceID,
	}

	// Make the appropiate SQL Call
	if err := tx.Get(&result, SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetGrowingLevelDeviceResponse{
		Device: result,
	}, nil
}

func (db *Persistence) GetGrowingSpotDevice(req routemodels.GetGrowingSpotDeviceRequest) (*routemodels.GetGrowingSpotDeviceResponse, error) {
	var result routemodels.Device

	SQL := `
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
	WHERE grow_spots_devices.growing_group = $1
		AND grow_spots_devices.device = $2
	`
	args := []interface{}{
		req.GrowingSpotID,
		req.DeviceID,
	}

	// Make the appropiate SQL Call
	if err := db.Postgres.Get(&result, SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetGrowingSpotDeviceResponse{
		Device: result,
	}, nil
}

func (db *Persistence) GetGrowingSpotDeviceWithTransaction(tx *sqlx.Tx, req routemodels.GetGrowingSpotDeviceRequest) (*routemodels.GetGrowingSpotDeviceResponse, error) {
	var result routemodels.Device

	SQL := `
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
	WHERE grow_spots_devices.growing_group = $1
		AND grow_spots_devices.device = $2
	`
	args := []interface{}{
		req.GrowingSpotID,
		req.DeviceID,
	}

	// Make the appropiate SQL Call
	if err := tx.Get(&result, SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	return &routemodels.GetGrowingSpotDeviceResponse{
		Device: result,
	}, nil
}

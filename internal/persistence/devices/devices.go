// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package devices

import (
	routemodels "totally-legit-grow-management/v1/internal/route-models"

	"github.com/jmoiron/sqlx"
)

// CreateDevice will insert a new Device into the persistence layer and return the created ID
func (db *DevicePersistence) CreateDevice(req *routemodels.CreateDeviceRequest) (*routemodels.CreateDeviceResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/

	var result string

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/

	args := []interface{}{
		req.DisplayName,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := db.Postgres.QueryRow(CREATE_DEVICE_SQL, args...).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.CreateDeviceResponse{
		ID: result,
	}, nil
}

// CreateDeviceWithTransaction will insert a new Device into the persistence layer within a transaction and return the created ID
func (db *DevicePersistence) CreateDeviceWithTransaction(tx *sqlx.Tx, req *routemodels.CreateDeviceRequest) (*routemodels.CreateDeviceResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result string

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/

	args := []interface{}{
		req.DisplayName,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := tx.QueryRow(CREATE_DEVICE_SQL, args...).Scan(result); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.CreateDeviceResponse{
		ID: result,
	}, nil
}

// CreateDeviceAction will insert a new Device Action into the persistence layer
func (db *DevicePersistence) CreateDeviceAction(req *routemodels.CreateDeviceActionRequest) (*routemodels.CreateDeviceActionResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result string

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/

	args := []interface{}{
		req.ID,
		req.DisplayName,
		req.InvokeRoute,
		req.Information,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if _, err := db.Postgres.Exec(CREATE_DEVICE_ACTIONS_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.CreateDeviceActionResponse{
		ID: result,
	}, nil
}

// CreateDeviceActionWithTransaction will insert a new Device Action into the persistence layer within a transaction
func (db *DevicePersistence) CreateDeviceActionWithTransaction(tx *sqlx.Tx, req *routemodels.CreateDeviceActionRequest) (*routemodels.CreateDeviceActionResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result string

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/

	args := []interface{}{
		req.ID,
		req.DisplayName,
		req.InvokeRoute,
		req.Information,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if _, err := tx.Exec(CREATE_DEVICE_ACTIONS_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.CreateDeviceActionResponse{
		ID: result,
	}, nil
}

// CreateGrowingGroupDevice will insert a new Growing Group Device into the persistence layer
func (db *DevicePersistence) CreateGrowingGroupDevice(req *routemodels.CreateGrowingGroupDeviceRequest) (*routemodels.CreateGrowingGroupDeviceResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/

	args := []interface{}{
		req.DeviceID,
		req.GrowingGroupID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if _, err := db.Postgres.Exec(CREATE_GROWING_GROUP_DEVICES_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.CreateGrowingGroupDeviceResponse{
		ID: "",
	}, nil
}

// CreateGrowingGroupDeviceWithTransaction will insert a new Growing Group Device into the persistence layer within a transaction
func (db *DevicePersistence) CreateGrowingGroupDeviceWithTransaction(tx *sqlx.Tx, req *routemodels.CreateGrowingGroupDeviceRequest) (*routemodels.CreateGrowingGroupDeviceResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/

	args := []interface{}{
		req.DeviceID,
		req.GrowingGroupID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if _, err := tx.Exec(CREATE_GROWING_GROUP_DEVICES_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.CreateGrowingGroupDeviceResponse{
		ID: "",
	}, nil
}

// CreateGrowingLocationDevice will insert a new Growing Location Device into the persistence layer
func (db *DevicePersistence) CreateGrowingLocationDevice(req *routemodels.CreateGrowingLocationDeviceRequest) (*routemodels.CreateGrowingLocationDeviceResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/

	args := []interface{}{
		req.DeviceID,
		req.GrowingLocationID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if _, err := db.Postgres.Exec(CREATE_GROWING_LOCATION_DEVICES_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.CreateGrowingLocationDeviceResponse{
		ID: "",
	}, nil
}

// CreateGrowingLocationDeviceWithTransaction will insert a new Growing Location Device into the persistence layer within a transaction
func (db *DevicePersistence) CreateGrowingLocationDeviceWithTransaction(tx *sqlx.Tx, req *routemodels.CreateGrowingLocationDeviceRequest) (*routemodels.CreateGrowingLocationDeviceResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/

	args := []interface{}{
		req.DeviceID,
		req.GrowingLocationID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if _, err := db.Postgres.Exec(CREATE_GROWING_LOCATION_DEVICES_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.CreateGrowingLocationDeviceResponse{
		ID: "",
	}, nil
}

// CreateGrowingLevelDevice will insert a new Growing Level Device into the persistence layer
func (db *DevicePersistence) CreateGrowingLevelDevice(req *routemodels.CreateGrowingLevelDeviceRequest) (*routemodels.CreateGrowingLevelDeviceResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/

	args := []interface{}{
		req.DeviceID,
		req.GrowingLevelID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if _, err := db.Postgres.Exec(CREATE_GROWING_LEVEL_DEVICES_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.CreateGrowingLevelDeviceResponse{
		ID: "",
	}, nil
}

// CreateGrowingLevelDevice will insert a new Growing Level Device into the persistence layer within a transaction
func (db *DevicePersistence) CreateGrowingLevelDeviceWithTransaction(sql *sqlx.Tx, req *routemodels.CreateGrowingLevelDeviceRequest) (*routemodels.CreateGrowingLevelDeviceResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/

	args := []interface{}{
		req.DeviceID,
		req.GrowingLevelID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if _, err := db.Postgres.Exec(CREATE_GROWING_LEVEL_DEVICES_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.CreateGrowingLevelDeviceResponse{
		ID: "",
	}, nil
}

// CreateGrowingSpotDevice will insert a new Growing Spot Device into the persistence layer
func (db *DevicePersistence) CreateGrowingSpotDevice(req *routemodels.CreateGrowingSpotDeviceRequest) (*routemodels.CreateGrowingSpotDeviceResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/

	args := []interface{}{
		req.DeviceID,
		req.GrowingLevelID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if _, err := db.Postgres.Exec(CREATE_GROW_SPOT_DEVICES_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.CreateGrowingSpotDeviceResponse{
		ID: "",
	}, nil
}

// CreateGrowingSpotDeviceWithTransaction will insert a new Growing Spot Device into the persistence layer within a transaction
func (db *DevicePersistence) CreateGrowingSpotDeviceWithTransaction(tx *sqlx.Tx, req *routemodels.CreateGrowingSpotDeviceRequest) (*routemodels.CreateGrowingSpotDeviceResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/

	args := []interface{}{
		req.DeviceID,
		req.GrowingLevelID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if _, err := db.Postgres.Exec(CREATE_GROW_SPOT_DEVICES_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.CreateGrowingSpotDeviceResponse{
		ID: "",
	}, nil
}

// DeleteDevice will set the Device to be archived
func (db *DevicePersistence) DeleteDevice(req *routemodels.DeleteDeviceRequest) error {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/
	args := []interface{}{
		req.ID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if _, err := db.Postgres.Exec(DELETE_DEVICE_SQL, args...); err != nil {
		// handle err
		return err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return nil
}

// DeleteGrowingGroupDevice will set the Growing Group Device to be archived
func (db *DevicePersistence) DeleteGrowingGroupDevice(req *routemodels.DeleteGrowingGroupDeviceRequest) error {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/
	args := []interface{}{
		req.DeviceID,
		req.GrowingGroupID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if _, err := db.Postgres.Exec(DELETE_GROWING_GROUP_DEVICE_SQL, args...); err != nil {
		// handle err
		return err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return nil
}

// DeleteGrowingLocationDevice will set the Growing Location Device to be archived
func (db *DevicePersistence) DeleteGrowingLocationDevice(req *routemodels.DeleteGrowingLocationDeviceRequest) error {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/
	args := []interface{}{
		req.DeviceID,
		req.GrowingLocationID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if _, err := db.Postgres.Exec(DELETE_GROWING_LOCATION_DEVICE_SQL, args...); err != nil {
		// handle err
		return err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return nil
}

// DeleteGrowingLevelDevice will set the Growing Level Device to be archived
func (db *DevicePersistence) DeleteGrowingLevelDevice(req *routemodels.DeleteGrowingLevelDeviceRequest) error {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/
	args := []interface{}{
		req.DeviceID,
		req.GrowingLevelID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if _, err := db.Postgres.Exec(DELETE_GROWING_LEVEL_DEVICE_SQL, args...); err != nil {
		// handle err
		return err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return nil
}

// DeleteGrowingSpotDevice will set the Growing Spot Device to be archived
func (db *DevicePersistence) DeleteGrowingSpotDevice(req *routemodels.DeleteGrowingSpotDeviceRequest) error {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/
	args := []interface{}{
		req.DeviceID,
		req.GrowingSpotID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if _, err := db.Postgres.Exec(DELETE_GROW_SPOT_DEVICE_SQL, args...); err != nil {
		// handle err
		return err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return nil
}

// EditDevice is not currently implemented
func (db *DevicePersistence) EditDevice(req *routemodels.EditDeviceRequest) (*routemodels.EditDeviceResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/

	//args := []interface{}{}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return nil, nil
}

// GetDevice will get a Device's information from the persistence layer
func (db *DevicePersistence) GetDevice(req *routemodels.GetDeviceRequest) (*routemodels.GetDeviceResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.Device

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/
	args := []interface{}{
		req.ID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := db.Postgres.Get(&result, GET_DEVICE_BY_ID_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetDeviceResponse{
		Device: result,
	}, nil
}

// GetDeviceActions will get a Device's Actions from the persistence layer
func (db *DevicePersistence) GetDeviceActions(req *routemodels.GetDeviceActionsRequest) (*routemodels.GetDeviceActionsResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result []routemodels.DeviceAction

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/
	args := []interface{}{
		req.DeviceID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := db.Postgres.Select(&result, GET_DEVICE_ACTIONS_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetDeviceActionsResponse{
		Actions: result,
	}, nil
}

// GetDeviceActionsWithTransaction will get a Device's Actions from the persistence layer within a transaction
func (db *DevicePersistence) GetDeviceActionsWithTransaction(tx *sqlx.Tx, req *routemodels.GetDeviceActionsRequest) (*routemodels.GetDeviceActionsResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result []routemodels.DeviceAction

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/
	args := []interface{}{
		req.DeviceID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := tx.Select(&result, GET_DEVICE_ACTIONS_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetDeviceActionsResponse{
		Actions: result,
	}, nil
}

// GetAllDevices will get all Devices from the persistence layer
func (db *DevicePersistence) GetAllDevices(req *routemodels.GetAllDevicesRequest) (*routemodels.GetAllDevicesResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result []routemodels.Device

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := db.Postgres.Select(&result, GET_ALL_DEVICES_SQL); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetAllDevicesResponse{
		Devices: result,
	}, nil
}

// GetAllGrowingGroupDevices will get all Growing Group Devices from the persistence layer
func (db *DevicePersistence) GetAllGrowingGroupDevices(req *routemodels.GetAllGrowingGroupDevicesRequest) (*routemodels.GetAllGrowingGroupDevicesResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result []routemodels.Device

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/
	args := []interface{}{
		req.GrowingGroupID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := db.Postgres.Select(&result, GET_ALL_GROWING_GROUP_DEVICES_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetAllGrowingGroupDevicesResponse{
		Devices: result,
	}, nil
}

// GetAllGrowingGroupDeviceWithTransaction will get all Growing Group Devices from the persistence layer within a transaction
func (db *DevicePersistence) GetAllGrowingGroupDeviceWithTransaction(tx *sqlx.Tx, req *routemodels.GetAllGrowingGroupDevicesRequest) (*routemodels.GetAllGrowingGroupDevicesResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result []routemodels.Device

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/
	args := []interface{}{
		req.GrowingGroupID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := tx.Select(&result, GET_ALL_GROWING_GROUP_DEVICES_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetAllGrowingGroupDevicesResponse{
		Devices: result,
	}, nil
}

// GetAllGrowingLocationDevices will get all Growing Location Devices from the persistence layer
func (db *DevicePersistence) GetAllGrowingLocationDevices(req *routemodels.GetAllGrowingLocationDevicesRequest) (*routemodels.GetAllGrowingLocationDevicesResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result []routemodels.Device

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/
	args := []interface{}{
		req.GrowingLocationID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := db.Postgres.Select(&result, GET_ALL_GROWING_LOCATION_DEVICES_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetAllGrowingLocationDevicesResponse{
		Devices: result,
	}, nil
}

// GetAllGrowingLocationDeviceWithTransaction will get all Growing Location Devices from the persistence layer within a transaction
func (db *DevicePersistence) GetAllGrowingLocationDeviceWithTransaction(tx *sqlx.Tx, req *routemodels.GetAllGrowingLocationDevicesRequest) (*routemodels.GetAllGrowingLocationDevicesResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result []routemodels.Device

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/
	args := []interface{}{
		req.GrowingLocationID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := tx.Select(&result, GET_ALL_GROWING_LOCATION_DEVICES_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetAllGrowingLocationDevicesResponse{
		Devices: result,
	}, nil
}

// GetAllGrowingLevelDevices will get all Growing Level Devices from the persistence layer
func (db *DevicePersistence) GetAllGrowingLevelDevices(req *routemodels.GetAllGrowingLevelDevicesRequest) (*routemodels.GetAllGrowingLevelDevicesResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result []routemodels.Device

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/
	args := []interface{}{
		req.GrowingLevelID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := db.Postgres.Get(&result, GET_ALL_GROWING_LEVEL_DEVICES_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetAllGrowingLevelDevicesResponse{
		Devices: result,
	}, nil
}

// GetAllGrowingLevelDeviceWithTransaction will get all Growing Level Devices from the persistence layer within a transaction
func (db *DevicePersistence) GetAllGrowingLevelDeviceWithTransaction(tx *sqlx.Tx, req *routemodels.GetAllGrowingLevelDevicesRequest) (*routemodels.GetAllGrowingLevelDevicesResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result []routemodels.Device

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/
	args := []interface{}{
		req.GrowingLevelID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := tx.Select(&result, GET_ALL_GROWING_LEVEL_DEVICES_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetAllGrowingLevelDevicesResponse{
		Devices: result,
	}, nil
}

// GetAllGrowingSpotDevices will get all Growing Spot Devices from the persistence layer
func (db *DevicePersistence) GetAllGrowingSpotDevices(req *routemodels.GetAllGrowingSpotDevicesRequest) (*routemodels.GetAllGrowingSpotDevicesResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result []routemodels.Device

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/
	args := []interface{}{
		req.GrowingSpotID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := db.Postgres.Select(&result, GET_ALL_GROWING_SPOT_DEVICES_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetAllGrowingSpotDevicesResponse{
		Devices: result,
	}, nil
}

// GetAllGrowingSpotDeviceWithTransaction will get all Growing Spot Devices from the persistence layer within a transaction
func (db *DevicePersistence) GetAllGrowingSpotDeviceWithTransaction(tx *sqlx.Tx, req *routemodels.GetAllGrowingSpotDevicesRequest) (*routemodels.GetAllGrowingSpotDevicesResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result []routemodels.Device

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/
	args := []interface{}{
		req.GrowingSpotID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := tx.Select(&result, GET_ALL_GROWING_SPOT_DEVICES_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetAllGrowingSpotDevicesResponse{
		Devices: result,
	}, nil
}

// GetGrowingGroupDevice will get a Growing Group Device from the persistence layer
func (db *DevicePersistence) GetGrowingGroupDevice(req *routemodels.GetGrowingGroupDeviceRequest) (*routemodels.GetGrowingGroupDeviceResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.Device

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/
	args := []interface{}{
		req.GrowingGroupID,
		req.DeviceID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := db.Postgres.Get(&result, GET_GROWING_GROUP_DEVICE_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetGrowingGroupDeviceResponse{
		Device: result,
	}, nil
}

// GetGrowingGroupDeviceWithTransaction will get a Growing Group Device from the persistence layer within a transaction
func (db *DevicePersistence) GetGrowingGroupDeviceWithTransaction(tx *sqlx.Tx, req *routemodels.GetGrowingGroupDeviceRequest) (*routemodels.GetGrowingGroupDeviceResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.Device

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/
	args := []interface{}{
		req.GrowingGroupID,
		req.DeviceID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := tx.Get(&result, GET_GROWING_GROUP_DEVICE_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetGrowingGroupDeviceResponse{
		Device: result,
	}, nil
}

// GetGrowingLocationDevice will get a Growing Location Device from the persistence layer
func (db *DevicePersistence) GetGrowingLocationDevice(req *routemodels.GetGrowingLocationDeviceRequest) (*routemodels.GetGrowingLocationDeviceResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.Device

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/
	args := []interface{}{
		req.GrowingLocationID,
		req.DeviceID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := db.Postgres.Get(&result, GET_GROWING_LOCATION_DEVICE_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetGrowingLocationDeviceResponse{
		Device: result,
	}, nil
}

// GetGrowingLocationDeviceWithTransaction will get a Growing Location Device from the persistence layer within a transaction
func (db *DevicePersistence) GetGrowingLocationDeviceWithTransaction(tx *sqlx.Tx, req *routemodels.GetGrowingLocationDeviceRequest) (*routemodels.GetGrowingLocationDeviceResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.Device

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/
	args := []interface{}{
		req.GrowingLocationID,
		req.DeviceID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := db.Postgres.Get(&result, GET_GROWING_LOCATION_DEVICE_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetGrowingLocationDeviceResponse{
		Device: result,
	}, nil
}

// GetGrowingLevelDevice will get a Growing Level Device from the persistence layer
func (db *DevicePersistence) GetGrowingLevelDevice(req *routemodels.GetGrowingLevelDeviceRequest) (*routemodels.GetGrowingLevelDeviceResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.Device

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/
	args := []interface{}{
		req.GrowingLevelID,
		req.DeviceID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := db.Postgres.Get(&result, GET_GROWING_LEVEL_DEVICE_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetGrowingLevelDeviceResponse{
		Device: result,
	}, nil
}

// GetGrowingLevelDeviceWithTransaction will get a Growing Level Device from the persistence layer within a transaction
func (db *DevicePersistence) GetGrowingLevelDeviceWithTransaction(tx *sqlx.Tx, req *routemodels.GetGrowingLevelDeviceRequest) (*routemodels.GetGrowingLevelDeviceResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.Device

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/
	args := []interface{}{
		req.GrowingLevelID,
		req.DeviceID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := tx.Get(&result, GET_GROWING_LEVEL_DEVICE_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetGrowingLevelDeviceResponse{
		Device: result,
	}, nil
}

// GetGrowingSpotDevice will get a Growing Spot Device from the persistence layer
func (db *DevicePersistence) GetGrowingSpotDevice(req *routemodels.GetGrowingSpotDeviceRequest) (*routemodels.GetGrowingSpotDeviceResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.Device

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/
	args := []interface{}{
		req.GrowingSpotID,
		req.DeviceID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := db.Postgres.Get(&result, GET_GROWING_SPOT_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetGrowingSpotDeviceResponse{
		Device: result,
	}, nil
}

// GetGrowingSpotDeviceWithTransaction will get a Growing Spot Device from the persistence layer within a transaction
func (db *DevicePersistence) GetGrowingSpotDeviceWithTransaction(tx *sqlx.Tx, req *routemodels.GetGrowingSpotDeviceRequest) (*routemodels.GetGrowingSpotDeviceResponse, error) {
	/**********************************************************************
	/
	/	State Stuff to Return
	/
	/**********************************************************************/
	var result routemodels.Device

	/**********************************************************************
	/
	/	Define Arguments For SQL Call
	/
	/**********************************************************************/
	args := []interface{}{
		req.GrowingSpotID,
		req.DeviceID,
	}

	/**********************************************************************
	/
	/	Do The SQL Call
	/
	/**********************************************************************/

	if err := tx.Get(&result, GET_GROWING_SPOT_SQL, args...); err != nil {
		// handle err
		return nil, err
	}

	/**********************************************************************
	/
	/	Return Expected Response
	/
	/**********************************************************************/
	return &routemodels.GetGrowingSpotDeviceResponse{
		Device: result,
	}, nil
}

// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routemodels

type CreateDeviceRequest struct {
	DisplayName string         `json:"displayName"`
	Actions     []DeviceAction `json:"actions"`
}

type CreateDeviceResponse struct {
	ID string `json:"id" db:"id"`
}

type CreateDeviceActionRequest struct {
	DisplayName string `json:"displayName"`
	DeviceAction
}

type CreateDeviceActionResponse struct {
	ID string `json:"id" db:"id"`
}

type CreateGrowingGroupDeviceRequest struct {
	DeviceID       string `json:"deviceID"`
	GrowingGroupID string `json:"growingGroupID"`
}

type CreateGrowingGroupDeviceResponse struct {
	ID string `json:"id" db:"id"`
}

type CreateGrowingLocationDeviceRequest struct {
	DeviceID          string `json:"deviceID"`
	GrowingLocationID string `json:"growingLocationID"`
}

type CreateGrowingLocationDeviceResponse struct {
	ID string `json:"id" db:"id"`
}

type CreateGrowingLevelDeviceRequest struct {
	DeviceID       string `json:"deviceID"`
	GrowingLevelID string `json:"growingLocationID"`
}

type CreateGrowingLevelDeviceResponse struct {
	ID string `json:"id" db:"id"`
}

type CreateGrowingSpotDeviceRequest struct {
	DeviceID       string `json:"deviceID"`
	GrowingLevelID string `json:"growingLevelID"`
}

type CreateGrowingSpotDeviceResponse struct {
	ID string `json:"id" db:"id"`
}

type DeleteDeviceRequest struct {
	ID string `json:"id"`
}

type DeleteGrowingGroupDeviceRequest struct {
	DeviceID       string `json:"deviceID"`
	GrowingGroupID string `json:"growingGroupID"`
}

type DeleteGrowingLocationDeviceRequest struct {
	DeviceID          string `json:"deviceID"`
	GrowingLocationID string `json:"growingLocationID"`
}

type DeleteGrowingLevelDeviceRequest struct {
	DeviceID       string `json:"deviceID"`
	GrowingLevelID string `json:"growingLevelID"`
}

type DeleteGrowingSpotDeviceRequest struct {
	DeviceID      string `json:"deviceID"`
	GrowingSpotID string `json:"growingSpotID"`
}

type EditDeviceRequest struct {
	DisplayName string `json:"displayName"`
}

type EditDeviceResponse struct {
	Device
}

type GetDeviceRequest struct {
	ID string `json:"id"`
}

type GetDeviceResponse struct {
	Device
}

type GetDeviceActionsRequest struct {
	DeviceID string `json:"deviceID"`
}

type GetDeviceActionsResponse struct {
	Actions []DeviceAction `json:"actions"`
}

type GetAllDevicesRequest struct {
	OrganizationID string `json:"organizationID"`
}

type GetAllDevicesResponse struct {
	Devices []Device
}

type GetGrowingGroupDeviceRequest struct {
	DeviceID       string `json:"deviceID"`
	GrowingGroupID string `json:"growingGroupID"`
}

type GetGrowingGroupDeviceResponse struct {
	Device Device `json:"device"`
}

type GetAllGrowingGroupDevicesRequest struct {
	GrowingGroupID string `json:"growingGroupID"`
}

type GetAllGrowingGroupDevicesResponse struct {
	Devices []Device `json:"devices"`
}

type GetGrowingLocationDeviceRequest struct {
	DeviceID          string `json:"deviceID"`
	GrowingLocationID string `json:"growingLocationID"`
}

type GetGrowingLocationDeviceResponse struct {
	Device Device `json:"device"`
}

type GetAllGrowingLocationDevicesRequest struct {
	GrowingLocationID string `json:"growingLocationID"`
}

type GetAllGrowingLocationDevicesResponse struct {
	Devices []Device `json:"devices"`
}

type GetGrowingLevelDeviceRequest struct {
	DeviceID       string `json:"deviceID"`
	GrowingLevelID string `json:"growingLevelID"`
}

type GetGrowingLevelDeviceResponse struct {
	Device Device `json:"device"`
}

type GetAllGrowingLevelDevicesRequest struct {
	GrowingLevelID string `json:"growingLevelID"`
}

type GetAllGrowingLevelDevicesResponse struct {
	Devices []Device `json:"devices"`
}

type GetGrowingSpotDeviceRequest struct {
	DeviceID      string `json:"deviceID"`
	GrowingSpotID string `json:"growingSpotID"`
}

type GetGrowingSpotDeviceResponse struct {
	Device Device `json:"device"`
}

type GetAllGrowingSpotDevicesRequest struct {
	GrowingSpotID string `json:"growingSpotID"`
}

type GetAllGrowingSpotDevicesResponse struct {
	Devices []Device `json:"devices"`
}

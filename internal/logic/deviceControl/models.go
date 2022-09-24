// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package devicecontrol

import (
	devicepersistence "totally-legit-grow-management/v1/internal/persistence/devices"
	routemodels "totally-legit-grow-management/v1/internal/route-models"
)

//
type IDeviceLogic interface {
	CreateDevice(*routemodels.CreateDeviceRequest) (*routemodels.CreateDeviceResponse, error)
	CreateDeviceAction(*routemodels.CreateDeviceActionRequest) (*routemodels.CreateDeviceActionResponse, error)
	CreateGrowingGroupDevice(*routemodels.CreateGrowingGroupDeviceRequest) (*routemodels.CreateGrowingGroupDeviceResponse, error)
	CreateGrowingLocationDevice(*routemodels.CreateGrowingLocationDeviceRequest) (*routemodels.CreateGrowingLocationDeviceResponse, error)
	CreateGrowingLevelDevice(*routemodels.CreateGrowingLevelDeviceRequest) (*routemodels.CreateGrowingLevelDeviceResponse, error)
	CreateGrowingSpotDevice(*routemodels.CreateGrowingSpotDeviceRequest) (*routemodels.CreateGrowingSpotDeviceResponse, error)
	DeleteDevice(*routemodels.DeleteDeviceRequest) error
	DeleteGrowingGroupDevice(*routemodels.DeleteGrowingGroupDeviceRequest) error
	DeleteGrowingLocationDevice(*routemodels.DeleteGrowingLocationDeviceRequest) error
	DeleteGrowingLevelDevice(*routemodels.DeleteGrowingLevelDeviceRequest) error
	DeleteGrowingSpotDevice(*routemodels.DeleteGrowingSpotDeviceRequest) error
	EditDevice(*routemodels.EditDeviceRequest) (*routemodels.EditDeviceResponse, error)
	GetAllDevices(*routemodels.GetAllDevicesRequest) (*routemodels.GetAllDevicesResponse, error)
	GetAllGrowingGroupDevices(*routemodels.GetAllGrowingGroupDevicesRequest) (*routemodels.GetAllGrowingGroupDevicesResponse, error)
	GetAllGrowingLocationDevices(*routemodels.GetAllGrowingLocationDevicesRequest) (*routemodels.GetAllGrowingLocationDevicesResponse, error)
	GetAllGrowingLevelDevices(*routemodels.GetAllGrowingLevelDevicesRequest) (*routemodels.GetAllGrowingLevelDevicesResponse, error)
	GetAllGrowingSpotDevices(*routemodels.GetAllGrowingSpotDevicesRequest) (*routemodels.GetAllGrowingSpotDevicesResponse, error)
	GetDevice(*routemodels.GetDeviceRequest) (*routemodels.GetDeviceResponse, error)
	GetDeviceActions(*routemodels.GetDeviceActionsRequest) (*routemodels.GetDeviceActionsResponse, error)
	GetGrowingGroupDevice(*routemodels.GetGrowingGroupDeviceRequest) (*routemodels.GetGrowingGroupDeviceResponse, error)
	GetGrowingLocationDevice(*routemodels.GetGrowingLocationDeviceRequest) (*routemodels.GetGrowingLocationDeviceResponse, error)
	GetGrowingLevelDevice(*routemodels.GetGrowingLevelDeviceRequest) (*routemodels.GetGrowingLevelDeviceResponse, error)
	GetGrowingSpotDevice(*routemodels.GetGrowingSpotDeviceRequest) (*routemodels.GetGrowingSpotDeviceResponse, error)
}

type DeviceControl struct {
	Persistence devicepersistence.IDevicesDB
}

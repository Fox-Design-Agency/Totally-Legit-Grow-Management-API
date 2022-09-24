// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package devicehandlers

import (
	"net/http"
	devicecontrol "totally-legit-grow-management/v1/internal/logic/deviceControl"
)

var _ IDeviceHandler = &DeviceHandler{}

type IDeviceHandler interface {
	CreateDeviceAction(w http.ResponseWriter, r *http.Request)
	CreateDevice(w http.ResponseWriter, r *http.Request)
	CreateGrowingGroupDevice(w http.ResponseWriter, r *http.Request)
	CreateGrowingLevelDevice(w http.ResponseWriter, r *http.Request)
	CreateGrowingLocationDevice(w http.ResponseWriter, r *http.Request)
	CreateGrowingSpotDevice(w http.ResponseWriter, r *http.Request)
	DeleteDevice(w http.ResponseWriter, r *http.Request)
	DeleteGrowingGroupDevice(w http.ResponseWriter, r *http.Request)
	DeleteGrowingLevelDevice(w http.ResponseWriter, r *http.Request)
	DeleteGrowingLocationDevice(w http.ResponseWriter, r *http.Request)
	DeleteGrowingSpotDevice(w http.ResponseWriter, r *http.Request)
	EditDevice(w http.ResponseWriter, r *http.Request)
	GetAllDevices(w http.ResponseWriter, r *http.Request)
	GetAllGrowingGroupDevices(w http.ResponseWriter, r *http.Request)
	GetAllGrowingLevelDevices(w http.ResponseWriter, r *http.Request)
	GetAllGrowingLocationDevices(w http.ResponseWriter, r *http.Request)
	GetAllGrowingSpotDevices(w http.ResponseWriter, r *http.Request)
	GetDeviceActions(w http.ResponseWriter, r *http.Request)
	GetDevice(w http.ResponseWriter, r *http.Request)
	GetGrowingGroupDevice(w http.ResponseWriter, r *http.Request)
	GetGrowingLevelDevice(w http.ResponseWriter, r *http.Request)
	GetGrowingLocationDevice(w http.ResponseWriter, r *http.Request)
	GetGrowingSpotDevice(w http.ResponseWriter, r *http.Request)
}

type DeviceHandler struct {
	Logic devicecontrol.IDeviceLogic
}

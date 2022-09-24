// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routes

import (
	"net/http"
	"totally-legit-grow-management/v1/pkg/server"

	"github.com/gorilla/mux"
)

func deviceRoutes(r *mux.Router, svr *server.Server) {
	r.HandleFunc("/api/v1/device", svr.DeviceHandlers.CreateDevice).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/device/action", svr.DeviceHandlers.CreateDeviceAction).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/device/growing-group", svr.DeviceHandlers.CreateGrowingGroupDevice).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/device/growing-location", svr.DeviceHandlers.CreateGrowingLocationDevice).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/device/growing-level", svr.DeviceHandlers.CreateGrowingLevelDevice).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/device/growing-spot", svr.DeviceHandlers.CreateGrowingSpotDevice).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/device", svr.DeviceHandlers.DeleteDevice).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/device/growing-group", svr.DeviceHandlers.DeleteGrowingGroupDevice).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/device/growing-location", svr.DeviceHandlers.DeleteGrowingLocationDevice).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/device/growing-level", svr.DeviceHandlers.DeleteGrowingLevelDevice).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/device/growing-spot", svr.DeviceHandlers.DeleteGrowingSpotDevice).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/device", svr.DeviceHandlers.EditDevice).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/device", svr.DeviceHandlers.GetDevice).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/device/actions", svr.DeviceHandlers.GetDeviceActions).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/devices", svr.DeviceHandlers.GetAllDevices).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/device/growing-group", svr.DeviceHandlers.GetGrowingGroupDevice).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/device/growing-location", svr.DeviceHandlers.GetGrowingLocationDevice).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/device/growing-level", svr.DeviceHandlers.GetGrowingLevelDevice).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/device/growing-spot", svr.DeviceHandlers.GetGrowingSpotDevice).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/devices/growing-group", svr.DeviceHandlers.GetAllGrowingGroupDevices).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/devices/growing-location", svr.DeviceHandlers.GetAllGrowingLocationDevices).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/devices/growing-level", svr.DeviceHandlers.GetAllGrowingLevelDevices).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/devices/growing-spot", svr.DeviceHandlers.GetAllGrowingSpotDevices).Methods(http.MethodGet)
}

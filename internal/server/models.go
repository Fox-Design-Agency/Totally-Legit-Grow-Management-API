// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package server

import (
	"totally-legit-grow-management/v1/internal/logic"
	devicehandlers "totally-legit-grow-management/v1/internal/server/device_handlers"
)

type Server struct {
	DeviceHandlers devicehandlers.IDeviceHandler
	Logic          logic.ILogic
}

// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package devicehandlers

import devicecontrol "totally-legit-grow-management/v1/internal/logic/deviceControl"

func InitDeviceHandler(deviceLogic devicecontrol.IDeviceLogic) *DeviceHandler {
	return &DeviceHandler{
		Logic: deviceLogic,
	}
}

// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package devicecontrol

import "totally-legit-grow-management/v1/internal/persistence"

func InitDeviceLogic(deviceDB persistence.IDevicesDB) *DeviceControl {
	return &DeviceControl{
		Persistence: deviceDB,
	}
}

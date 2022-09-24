package devicehandlers

var _ IDeviceHandler = &DeviceHandler{}

func InitDeviceHandler() *DeviceHandler {
	return &DeviceHandler{}
}

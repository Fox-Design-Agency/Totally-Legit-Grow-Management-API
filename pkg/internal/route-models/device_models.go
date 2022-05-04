package routemodels

type SlimDevice struct {
	ID          string `json:"id" db:"device_id"`
	DisplayName string `json:"displayName" db:"display_name"`
}

type SlimDeviceAction struct {
	ID          string `json:"id" db:"action_id"`
	DisplayName string `json:"displayName" db:"action_name"`
}

type DeviceAction struct {
	SlimDeviceAction
	InvokeRoute string `json:"invokeRoute" db:"invoke_route"`
	Information string `json:"information" db:"action_info"`
}

type Device struct {
	SlimDevice
	Actions []DeviceAction `json:"actions"`
	CreatedAtMember
	UpdatedAtMember
}

package service

import (
	"context"
	"encoding/json"
	"netedit/utility/utilerr"

	"github.com/Wifx/gonetworkmanager/v2"
)

type IDevice interface {
	GetDevice(ctx context.Context) (device []map[string]interface{}, err error)
}

var deviceService = deviceImpl{}

func Device() IDevice {
	return &deviceService
}

type deviceImpl struct{}

func (s *deviceImpl) GetDevice(ctx context.Context) (device []map[string]interface{}, err error) {
	nm, err := gonetworkmanager.NewNetworkManager()
	utilerr.ErrIsNil(ctx, err, "NewNetworkManager failed")
	/* Get devices */
	nmdevices, err := nm.GetPropertyAllDevices()
	utilerr.ErrIsNil(ctx, err, "GetPropertyAllDevices failed")

	/* Show each device path and interface name */
	for _, nmdevice := range nmdevices {

		deviceInterface, err := nmdevice.GetPropertyInterface()
		utilerr.ErrIsNil(ctx, err, "device.GetPropertyInterface failed")
		deviceInfo, err := nmdevice.MarshalJSON()
		utilerr.ErrIsNil(ctx, err, "device.MarshalJSON failed")
		var deviceInfoUnmarshaled interface{}
		err = json.Unmarshal(deviceInfo, &deviceInfoUnmarshaled)
		utilerr.ErrIsNil(ctx, err, "device.UnmarshalJSON failed")
		device = append(device, map[string]interface{}{
			deviceInterface: deviceInfoUnmarshaled,
		})
	}

	return
}

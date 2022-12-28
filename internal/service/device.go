package service

import (
	"context"
	"netedit/utility/utilerr"

	"github.com/Wifx/gonetworkmanager/v2"
)

type IDevice interface {
	GetDevice(ctx context.Context) (deviceList []string, err error)
}

var deviceService = deviceImpl{}

func Device() IDevice {
	return &deviceService
}

type deviceImpl struct{}

func (s *deviceImpl) GetDevice(ctx context.Context) (deviceList []string, err error) {
	nm, err := gonetworkmanager.NewNetworkManager()
	utilerr.ErrIsNil(ctx, err, "NewNetworkManager failed")
	/* Get devices */
	nmdevices, err := nm.GetAllDevices()
	utilerr.ErrIsNil(ctx, err, "GetPropertyAllDevices failed")
	for _, device := range nmdevices {
		deviceName, err := device.GetPropertyInterface()
		utilerr.ErrIsNil(ctx, err, "device.GetPropertyInterface failed")
		deviceList = append(deviceList, deviceName)
	}

	return
}

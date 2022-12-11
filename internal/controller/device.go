package controller

import (
	"context"
	"netedit/internal/service"
	"netedit/utility/utilerr"

	"netedit/api/v1"
)

var (
	Device = cDevice{}
)

type cDevice struct{}

func (c *cDevice) Device(ctx context.Context, req *v1.DeviceReq) (res *v1.DeviceRes, err error) {
	deviceList, err := service.Device().GetDevice(ctx)
	utilerr.ErrIsNil(ctx, err, "get deviceList failed")
	res = &v1.DeviceRes{
		DeviceList: deviceList,
	}
	return
}

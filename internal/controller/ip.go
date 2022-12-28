package controller

import (
	"context"
	"netedit/internal/service"
	"netedit/utility/utilerr"

	"netedit/api/v1"
)

var (
	IP = cIP{}
)

type cIP struct{}

func (c *cIP) IP(ctx context.Context, req *v1.IPReq) (res *v1.IPRes, err error) {
	ipList, err := service.IP(req.DeviceName).GetIP(ctx)
	utilerr.ErrIsNil(ctx, err, "get ip failed")
	res = &v1.IPRes{
		IPList: ipList,
	}
	return
}

func (c *cIP) UpdateIP(ctx context.Context, req *v1.UpdateIPReq) (res *v1.UpdateIPRes, err error) {
	err = service.IP(req.DeviceName).UpdateIP(ctx, req.OldIP, req.NewIP)
	utilerr.ErrIsNil(ctx, err, "update ip failed")
	res = &v1.UpdateIPRes{}
	return
}

func (c *cIP) AddIP(ctx context.Context, req *v1.AddIPReq) (res *v1.AddIPRes, err error) {
	err = service.IP(req.DeviceName).AddIP(ctx, req.IP)
	utilerr.ErrIsNil(ctx, err, "add ip failed")
	res = &v1.AddIPRes{}
	return
}

func (c *cIP) DeleteIP(ctx context.Context, req *v1.DeleteIPReq) (res *v1.DeleteIPRes, err error) {
	err = service.IP(req.DeviceName).DeleteIP(ctx, req.IP)
	utilerr.ErrIsNil(ctx, err, "delete ip failed")
	res = &v1.DeleteIPRes{}
	return
}

package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type DeviceReq struct {
	g.Meta `path:"/device" tags:"Device" method:"get" summary:"nm device list"`
}

type DeviceRes struct {
	g.Meta     `mime:"application/json"`
	DeviceList []map[string]interface{}
}

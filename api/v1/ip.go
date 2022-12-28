package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type IPReq struct {
	g.Meta `path:"/ip" tags:"IP" method:"get" summary:"hostname -I"`
	DeviceName string `p:"device" v:"required"`
}

type IPRes struct {
	g.Meta     `mime:"application/json"`
	IPList []string `json:"iplist"`
}

type UpdateIPReq struct {
	g.Meta `path:"/ip" tags:"IP" method:"put" summary:"update ip"`
	DeviceName string `p:"device" v:"required"`
	OldIP string `p:"oldip" v:"required|ipv4"`
	NewIP string `p:"newip" v:"required|ipv4"`
}

type UpdateIPRes struct {
	g.Meta     `mime:"application/json"`
}

type AddIPReq struct {
	g.Meta `path:"/ip" tags:"IP" method:"post" summary:"add ip"`
	DeviceName string `p:"device" v:"required"`
	IP string `p:"ip" v:"required|ipv4"`
}

type AddIPRes struct {
	g.Meta     `mime:"application/json"`
}

type DeleteIPReq struct {
	g.Meta `path:"/ip" tags:"IP" method:"delete" summary:"delete ip"`
	DeviceName string `p:"device" v:"required"`
	IP string `p:"ip" v:"required|ipv4"`
}

type DeleteIPRes struct {
	g.Meta     `mime:"application/json"`
}

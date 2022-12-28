package utilerr

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

func ErrIsNil(ctx context.Context, err error, msg ...string) {
	if !g.IsNil(err) {
		if len(msg) > 0 {
			g.Log().Error(ctx, err.Error())
			g.Log().Error(ctx, msg[0])
		} else {
			g.Log().Error(ctx, err.Error())
		}
	}
}

func ValueIsNil(ctx context.Context,value interface{}, msg string) {
	if g.IsNil(value) {
		g.Log().Error(ctx, msg)
	}
}

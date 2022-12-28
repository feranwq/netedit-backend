package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"

	"netedit/internal/controller"
)

type cMain struct {
	g.Meta `name:"main" brief:"start http server"`
}

type cMainHttpInput struct {
	g.Meta `name:"http" brief:"start http server"`
	Port   int    `v:"required" short:"p" name:"port"  brief:"port of http server"`
}

type cMainHttpOutput struct{}

func (c *cMain) Http(ctx context.Context, in cMainHttpInput) (out *cMainHttpOutput, err error) {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		group.Bind(
			controller.Device,
			controller.IP,
		)
	})
	s.SetPort(in.Port)
	s.Run()
	return
}


func Main() {
	cmd, err := gcmd.NewFromObject(cMain{})
	if err != nil {
		panic(err)
	}
	cmd.Run(gctx.New())
}

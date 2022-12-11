package main

import (
	_ "netedit/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"netedit/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}

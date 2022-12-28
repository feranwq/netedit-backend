package utilexec

import (
	"bytes"
	"context"
	"os"
	"os/exec"

	"github.com/gogf/gf/v2/frame/g"
)

//Cmd is exec on os ,no return
func Cmd(ctx context.Context, name string, arg ...string) {
	g.Log().Info(ctx, "[os]exec cmd is : ", name, arg)
	cmd := exec.Command(name, arg[:]...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		g.Log().Error(ctx, "[os]os call error.", err)
	}
}

//String is exec on os , return result
func String(ctx context.Context, name string, arg ...string) string {
	g.Log().Info(ctx, "[os]exec cmd is : ", name, arg)
	cmd := exec.Command(name, arg[:]...)
	cmd.Stdin = os.Stdin
	var b bytes.Buffer
	cmd.Stdout = &b
	cmd.Stderr = &b
	if err := cmd.Run(); err != nil {
		g.Log().Error(ctx, "[os]os call error.", err)
		return ""
	}
	return b.String()
}

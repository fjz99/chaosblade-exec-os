package nginx

import (
	"context"
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
)

func NewNginxCommandSpec() spec.ExpModelCommandSpec {
	return &NginxCommandSpec{
		spec.BaseExpModelCommandSpec{
			ExpActions: []spec.ExpActionCommandSpec{
				NewCrashActionSpec(),
				NewRestartActionSpec(),
				NewConfigActionSpec(),
				NewResponseActionSpec(),
			},
			ExpFlags: []spec.ExpFlagSpec{},
		},
	}
}

func testNginxExists(channel spec.Channel, ctx context.Context) *spec.Response {
	return nil
}

func killNginx(channel spec.Channel, ctx context.Context) *spec.Response {
	return nil
}

func runNginxCommand(channel spec.Channel, ctx context.Context, args string) *spec.Response {
	return channel.Run(ctx, "nginx", args)
}

func restoreConfigFile(channel spec.Channel, ctx context.Context, backup, activeFile string) *spec.Response {
	return nil
}

func backupConfigFile(channel spec.Channel, ctx context.Context, backup string, activeFile string, newFile string, remove bool) *spec.Response {
	return nil
}

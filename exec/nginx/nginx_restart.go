package nginx

import (
	"context"
	"fmt"
	"github.com/chaosblade-io/chaosblade-exec-os/exec/category"
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
)

const NginxRestartBin = "chaos_nginxrestart"

type RestartActionSpec struct {
	spec.BaseExpActionCommandSpec
}

func NewRestartActionSpec() spec.ExpActionCommandSpec {
	return &RestartActionSpec{
		spec.BaseExpActionCommandSpec{
			ActionMatchers: []spec.ExpFlagSpec{},
			ActionFlags:    []spec.ExpFlagSpec{},
			ActionExecutor: &NginxRestartExecutor{},
			ActionExample: `
# Block outgoing connection to the specific domain on port 80
blade create network drop --destination-port 80 --string-pattern baidu.com --network-traffic out
`,
			ActionPrograms:   []string{NginxRestartBin},
			ActionCategories: []string{category.Middleware},
		},
	}
}

func (*RestartActionSpec) Name() string {
	return "restart"
}

func (*RestartActionSpec) Aliases() []string {
	return []string{}
}

func (*RestartActionSpec) ShortDesc() string {
	return "Restart experiment"
}

func (d *RestartActionSpec) LongDesc() string {
	if d.ActionLongDesc != "" {
		return d.ActionLongDesc
	}
	return "Nginx restart"
}

type NginxRestartExecutor struct {
	channel spec.Channel
}

func (*NginxRestartExecutor) Name() string {
	return "restart"
}

func (ng *NginxRestartExecutor) Exec(suid string, ctx context.Context, model *spec.ExpModel) *spec.Response {
	commands := []string{"kill"}
	if response, ok := ng.channel.IsAllCommandsAvailable(ctx, commands); !ok {
		return response
	}

	if _, ok := spec.IsDestroy(ctx); ok {
		return spec.ReturnFail(spec.OsCmdExecFailed, "cancel 'nginx restart' is meaningless")
	}
	return ng.start(ctx)
}

func (ng *NginxRestartExecutor) start(ctx context.Context) *spec.Response {
	allPid, response := getNginxPid(ng.channel, ctx)
	if response != nil {
		return response
	}
	for _, pid := range allPid {
		response = ng.channel.Run(ctx, fmt.Sprintf("kill -9 %d", pid), "")
		if !response.Success {
			return response
		}
	}

	response = ng.channel.Run(ctx, "nginx", "")
	return response
}

func (ng *NginxRestartExecutor) SetChannel(channel spec.Channel) {
	ng.channel = channel
}

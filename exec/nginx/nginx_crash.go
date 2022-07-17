package nginx

import (
	"context"
	"fmt"
	"github.com/chaosblade-io/chaosblade-exec-os/exec/category"
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
)

const NginxCrashBin = "chaos_nginxcrash"

type CrashActionSpec struct {
	spec.BaseExpActionCommandSpec
}

func NewCrashActionSpec() spec.ExpActionCommandSpec {
	return &CrashActionSpec{
		spec.BaseExpActionCommandSpec{
			ActionMatchers: []spec.ExpFlagSpec{},
			ActionFlags:    []spec.ExpFlagSpec{},
			ActionExecutor: &NginxCrashExecutor{},
			ActionExample: `
# Block outgoing connection to the specific domain on port 80
blade create network drop --destination-port 80 --string-pattern baidu.com --network-traffic out
`,
			ActionPrograms:   []string{NginxCrashBin},
			ActionCategories: []string{category.Middleware},
		},
	}
}

func (*CrashActionSpec) Name() string {
	return "crash"
}

func (*CrashActionSpec) Aliases() []string {
	return []string{}
}

func (*CrashActionSpec) ShortDesc() string {
	return "Crash experiment"
}

func (d *CrashActionSpec) LongDesc() string {
	if d.ActionLongDesc != "" {
		return d.ActionLongDesc
	}
	return "Nginx crash"
}

type NginxCrashExecutor struct {
	channel spec.Channel
}

func (*NginxCrashExecutor) Name() string {
	return "crash"
}

func (ng *NginxCrashExecutor) Exec(suid string, ctx context.Context, model *spec.ExpModel) *spec.Response {
	commands := []string{"kill"}
	if response, ok := ng.channel.IsAllCommandsAvailable(ctx, commands); !ok {
		return response
	}

	if _, ok := spec.IsDestroy(ctx); ok {
		return ng.stop(ctx)
	}
	return ng.start(ctx)
}

func (ng *NginxCrashExecutor) start(ctx context.Context) *spec.Response {
	allPid, response := getNginxPid(ng.channel, ctx)
	if response != nil {
		return response
	}
	fmt.Println(allPid)
	for _, pid := range allPid {
		response = ng.channel.Run(ctx, fmt.Sprintf("kill -9 %d", pid), "")
		fmt.Println(response)
		if !response.Success {
			return response
		}
	}
	return response
}

func (ng *NginxCrashExecutor) stop(ctx context.Context) *spec.Response {
	response := ng.channel.Run(ctx, "nginx", "")
	return response
}

func (ng *NginxCrashExecutor) SetChannel(channel spec.Channel) {
	ng.channel = channel
}

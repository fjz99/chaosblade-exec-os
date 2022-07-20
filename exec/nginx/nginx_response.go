package nginx

import (
	"context"
	"github.com/chaosblade-io/chaosblade-exec-os/exec/category"
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
)

const NginxResponseBin = "chaos_nginxresponse"

type ResponseActionSpec struct {
	spec.BaseExpActionCommandSpec
}

func NewResponseActionSpec() spec.ExpActionCommandSpec {
	return &ResponseActionSpec{
		spec.BaseExpActionCommandSpec{
			ActionMatchers:   []spec.ExpFlagSpec{},
			ActionFlags:      []spec.ExpFlagSpec{},
			ActionExecutor:   &NginxResponseExecutor{},
			ActionExample:    ``,
			ActionPrograms:   []string{NginxResponseBin},
			ActionCategories: []string{category.Middleware},
		},
	}
}

func (*ResponseActionSpec) Name() string {
	return "response"
}

func (*ResponseActionSpec) Aliases() []string {
	return []string{}
}

func (*ResponseActionSpec) ShortDesc() string {
	return "Response experiment"
}

func (d *ResponseActionSpec) LongDesc() string {
	if d.ActionLongDesc != "" {
		return d.ActionLongDesc
	}
	return "Nginx response experiment"
}

type NginxResponseExecutor struct {
	channel spec.Channel
}

func (*NginxResponseExecutor) Name() string {
	return "response"
}

func (ng *NginxResponseExecutor) Exec(suid string, ctx context.Context, model *spec.ExpModel) *spec.Response {

	if _, ok := spec.IsDestroy(ctx); ok {
		return ng.stop(ctx)
	}
	return ng.start(ctx)
}

func (ng *NginxResponseExecutor) start(ctx context.Context) *spec.Response {
	return nil
}

func (ng *NginxResponseExecutor) stop(ctx context.Context) *spec.Response {
	return nil
}

func (ng *NginxResponseExecutor) SetChannel(channel spec.Channel) {
	ng.channel = channel
}

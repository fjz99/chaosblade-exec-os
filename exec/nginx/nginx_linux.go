package nginx

import (
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
)

func NewNginxCommandSpec() spec.ExpModelCommandSpec {
	return &NginxCommandSpec{
		spec.BaseExpModelCommandSpec{
			ExpActions: []spec.ExpActionCommandSpec{
				NewCrashActionSpec(),
				NewRestartActionSpec(),
				NewConfigActionSpec(),
			},
			ExpFlags: []spec.ExpFlagSpec{},
		},
	}
}

package nginx

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/chaosblade-io/chaosblade-exec-os/exec/category"
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
	"github.com/chaosblade-io/chaosblade-spec-go/util"
)

const (
	NginxConfigBin   = "chaos_nginxconfig"
	configBackupName = "nginx.conf.chaosblade.back"
)

type ConfigActionSpec struct {
	spec.BaseExpActionCommandSpec
}

func NewConfigActionSpec() spec.ExpActionCommandSpec {
	return &ConfigActionSpec{
		spec.BaseExpActionCommandSpec{
			ActionMatchers: []spec.ExpFlagSpec{
				&spec.ExpFlag{
					Name: "file",
					Desc: "The new nginx.conf file",
				},
				&spec.ExpFlag{
					Name:   "force",
					Desc:   "If force == true, ChaosBlade will delete config backup file if it exists",
					NoArgs: true,
				},
			},
			ActionFlags:    []spec.ExpFlagSpec{},
			ActionExecutor: &NginxConfigExecutor{},
			ActionExample: `
# Block outgoing connection to the specific domain on port 80
blade create nginx config --file newConfig.conf --string-pattern baidu.com --network-traffic out
`,
			ActionPrograms:   []string{NginxConfigBin},
			ActionCategories: []string{category.Middleware},
		},
	}
}

func (*ConfigActionSpec) Name() string {
	return "config"
}

func (*ConfigActionSpec) Aliases() []string {
	return []string{}
}

func (*ConfigActionSpec) ShortDesc() string {
	return "Config experiment"
}

func (d *ConfigActionSpec) LongDesc() string {
	if d.ActionLongDesc != "" {
		return d.ActionLongDesc
	}
	return "Nginx config"
}

type NginxConfigExecutor struct {
	channel spec.Channel
}

func (*NginxConfigExecutor) Name() string {
	return "config"
}

func (ng *NginxConfigExecutor) Exec(suid string, ctx context.Context, model *spec.ExpModel) *spec.Response {
	_, response := getNginxPid(ng.channel, ctx) // conf nginx process
	if response != nil {
		return response
	}

	activeFile, response := getNginxConfigLocation(ng.channel, ctx)
	if response != nil {
		return response
	}
	dir := activeFile[:strings.LastIndex(activeFile, string(os.PathSeparator))+1]
	backup := dir + configBackupName

	if _, ok := spec.IsDestroy(ctx); ok {
		return ng.stop(ctx, dir, activeFile, backup, model)
	}
	return ng.start(ctx, dir, activeFile, backup, model)
}

func (ng *NginxConfigExecutor) start(ctx context.Context, dir, activeFile, backup string, model *spec.ExpModel) *spec.Response {
	if util.IsExist(backup) {
		force := model.ActionFlags["force"] == "true"
		if force {
			if response := ng.channel.Run(ctx, fmt.Sprintf("rm %s", backup), ""); !response.Success {
				return response
			}
		} else {
			return spec.ReturnFail(spec.OsCmdExecFailed,
				fmt.Sprintf("cannot change config due to backup file %s exists", backup))
		}
	}

	newFile := model.ActionFlags["file"]
	if newFile == "" {
		//创建配置文件
		newFile = createNewConfig()
	} else {
		if !util.IsExist(newFile) || util.IsDir(newFile) {
			return spec.ReturnFail(spec.OsCmdExecFailed, fmt.Sprintf("config file %s not exists", newFile))
		}
	}
	if response := testNginxConfig(ng.channel, ctx, newFile, dir); response != nil {
		return response
	}

	cmd := fmt.Sprintf("cp %s %s && cp -f %s %s", activeFile, backup, newFile, activeFile)
	if model.ActionFlags["file"] == "" {
		// remove auto generated config
		cmd += fmt.Sprintf(" && rm %s", newFile)
	}
	cmd += " && nginx -s reload"
	response := ng.channel.Run(ctx, cmd, "")
	if !response.Success {
		return response
	}

	return spec.ReturnSuccess("nginx config changed")
}

func createNewConfig() string {
	newFile := "nginx.conf.tmp"
	panic("not impl")
	return newFile
}

//通用
func (ng *NginxConfigExecutor) stop(ctx context.Context, dir, activeFile, backup string, model *spec.ExpModel) *spec.Response {
	if !util.IsExist(backup) || util.IsDir(backup) {
		return spec.ReturnFail(spec.OsCmdExecFailed, fmt.Sprintf("backup file %s not exists", backup))
	}

	response := ng.channel.Run(ctx, fmt.Sprintf("mv -f %s %s && nginx -s reload", backup, activeFile), "")
	if !response.Success {
		return response
	}
	return spec.ReturnSuccess("nginx config restored")
}

func (ng *NginxConfigExecutor) SetChannel(channel spec.Channel) {
	ng.channel = channel
}

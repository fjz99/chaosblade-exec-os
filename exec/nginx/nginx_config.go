package nginx

import (
	"context"
	"fmt"
	"github.com/chaosblade-io/chaosblade-exec-os/exec/nginx/parser"
	"os"
	"strconv"
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
					Name: "block-id",
					Desc: "",
				},
				&spec.ExpFlag{
					Name: "set-config",
					Desc: "set multiple key-value config paris for specified block-id",
				},
			},
			ActionFlags: []spec.ExpFlagSpec{
				&spec.ExpFlag{
					Name:   "force",
					Desc:   "ChaosBlade will delete config backup file if it exists",
					NoArgs: true,
				},
				&spec.ExpFlag{
					Name:   "list",
					Desc:   "List all nginx config blocks",
					NoArgs: true,
				},
				//&spec.ExpFlag{
				//	Name:   "pretty",
				//	Desc:   "Print all nginx config blocks in pretty format",
				//	NoArgs: true,
				//},
			},
			ActionExecutor: &NginxConfigExecutor{},
			ActionExample: `
# Block outgoing connection to the specific domain on port 80
blade create nginx config --file newConfig.conf --string-pattern baidu.com --network-traffic out

chaos_os create nginx config --force=true --set-config 'xx=xx;xx=xx'

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
	for k, v := range model.ActionFlags {
		fmt.Println(k, v)
	}
	if true {
		result := parser.ListResult{Block: &parser.Block{Header: "ffff"}, Header: "server", Type: "server", Id: 1}
		return spec.ReturnResultIgnoreCode(result)
	}
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
	var config *parser.Config
	if model.ActionFlags["list"] == "true" {
		config, _ = parser.LoadConfig(activeFile)
		config.ListAllBlocks()
		return spec.Success()
	}

	newFile := model.ActionFlags["file"]
	if newFile == "" {
		//create new config
		if config == nil {
			config, _ = parser.LoadConfig(activeFile)
		}
		var response *spec.Response
		newFile, response = createNewConfig(config, model.ActionFlags["block-id"], model.ActionFlags["set-config"])
		if response != nil {
			return response
		}
		return response
	} else {
		if !util.IsExist(newFile) || util.IsDir(newFile) {
			return spec.ReturnFail(spec.OsCmdExecFailed, fmt.Sprintf("config file %s not exists", newFile))
		}
	}
	if response := testNginxConfig(ng.channel, ctx, newFile, dir); response != nil {
		return response
	}

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

func createNewConfig(config *parser.Config, id string, newKV string) (string, *spec.Response) {
	blockId, err := strconv.Atoi(id)
	if err != nil {
		return "", spec.ReturnFail(spec.OsCmdExecFailed, fmt.Sprintf("block-id %s is not valid", id))
	}
	blocksList := config.GetBlocksList()
	for _, kv := range strings.Split(newKV, ";") {
		arr := strings.Split(strings.Trim(kv, " "), "=")
		if len(arr) != 2 {
			return "", spec.ReturnFail(spec.OsCmdExecFailed, fmt.Sprintf("set-config %s is not valid", newKV))
		}
		k := strings.Trim(arr[0], " ")
		v := strings.Trim(arr[1], " ")
		blocksList[blockId].Block.Statements[k] = parser.Statement{Key: k, Value: v}
	}
	name := "nginx.chaosblade.tmp.conf"
	err = config.EasyDumpToFile(name)
	if err != nil {
		return "", spec.ReturnFail(spec.OsCmdExecFailed, err.Error())
	}
	return name, nil
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

func listAllBlocks() string {
	newFile := "nginx.conf.tmp"
	panic("not impl")
	return newFile
}

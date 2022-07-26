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
	fileMode         = "file"
	cmdMode          = "cmd"
)

//TODO 支持版本链
type ConfigActionSpec struct {
	spec.BaseExpActionCommandSpec
}

func NewConfigActionSpec() spec.ExpActionCommandSpec {
	return &ConfigActionSpec{
		spec.BaseExpActionCommandSpec{
			ActionMatchers: []spec.ExpFlagSpec{
				&spec.ExpFlag{
					Name: "mode",
					Desc: fmt.Sprintf("The configuration change mode (%s or %s)", fileMode, cmdMode),
				},
				&spec.ExpFlag{
					Name: "file",
					Desc: "The new nginx.conf file",
				},
				&spec.ExpFlag{
					Name: "block-id",
					Desc: "target block id for config modification",
				},
				&spec.ExpFlag{
					Name: "set-config",
					Desc: "set multiple key-value config paris for specified block-id",
				},
			},
			ActionFlags: []spec.ExpFlagSpec{
				// &spec.ExpFlag{
				// 	Name:   "force",
				// 	Desc:   "ChaosBlade will delete config backup file if it exists",
				// 	NoArgs: true,
				// },
				&spec.ExpFlag{
					Name:   "list",
					Desc:   "List all nginx config blocks",
					NoArgs: true,
				},
			},
			ActionExecutor: &NginxConfigExecutor{},
			ActionExample: `
# List all nginx.conf blocks
blade create nginx config --list

# Change config file to my.conf
blade create nginx config --mode file --file my.conf

//!!!
# Change config file to my.conf, and delete nginx conf backup file if it exists
blade create nginx config --file my.conf --force
//!!!

# Change 'server' (assuming block id = 1) exposed on port 8899
blade create nginx config --mode cmd --block-id 3 --set-config='listen=8899'

# Set 'location /' (assuming block id = 3) proxy_pass to www.baidu.com
blade create nginx config --mode cmd --block-id 3 --set-config='proxy_pass=www.baidu.com'

# Revert config change, uid = xxx
blade destroy nginx config --uid 

# Revert config change to the oldest config file
blade destroy nginx config --force
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
	// for k, v := range model.ActionFlags {
	// fmt.Println(k, v)
	// }
	// if true {
	// 	result := parser.ListResult{Block: &parser.Block{Header: "ffff"}, Header: "server", Type: "server", Id: 1}
	// 	return spec.ReturnResultIgnoreCode(result)
	// }
	_, response := getNginxPid(ng.channel, ctx) // nginx process
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
	mode := model.ActionFlags["mode"]
	newFile := model.ActionFlags["file"]
	switch mode {
	case fileMode:
		if newFile == "" || !util.IsExist(newFile) || util.IsDir(newFile) {
			return spec.ReturnFail(spec.OsCmdExecFailed, fmt.Sprintf("config file '%s' not exists", newFile))
		}
	case cmdMode:
		if config == nil {
			config, _ = parser.LoadConfig(activeFile)
		}
		var resp *spec.Response
		newFile, resp = createNewConfig(config, model.ActionFlags["block-id"], model.ActionFlags["set-config"])
		if resp != nil {
			return resp
		}
	default:
		return spec.ReturnFail(spec.OsCmdExecFailed, fmt.Sprintf("invalid --mode argument, which must be '%s' or '%s'", fileMode, cmdMode))
	}
	// if mode == "file" {
	// 	if newFile == "" || !util.IsExist(newFile) || util.IsDir(newFile) {
	// 		return spec.ReturnFail(spec.OsCmdExecFailed, fmt.Sprintf("config file %s not exists", newFile))
	// 	}
	// } else if mode == "" {
	// 	//create new config
	// 	if config == nil {
	// 		config, _ = parser.LoadConfig(activeFile)
	// 	}
	// 	newFile, _ = createNewConfig(config, model.ActionFlags["block-id"], model.ActionFlags["set-config"])
	// } else {
	// 	return spec.ReturnFail(spec.OsCmdExecFailed, "mode cannot be null")
	// }
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
	blocksList := config.GetBlocksList()
	blockId, err := strconv.Atoi(id)
	fmt.Println("id=", id, err)
	if err != nil || blockId-1 >= len(blocksList) || blockId < 0 {
		return "", spec.ReturnFail(spec.OsCmdExecFailed, fmt.Sprintf("--block-id='%s' is not valid, expect %d-%d", id, 0, len(blocksList)))
	}
	for _, kv := range strings.Split(newKV, ";") {
		arr := strings.Split(strings.TrimSpace(kv), "=")
		if newKV == "" || len(arr) != 2 {
			return "", spec.ReturnFail(spec.OsCmdExecFailed, fmt.Sprintf("--set-config='%s' is not valid", newKV))
		}
		k := strings.TrimSpace(arr[0])
		v := strings.TrimSpace(arr[1])
		if blockId == 0 {
			config.Statements[k] = parser.Statement{Key: k, Value: v}
		} else {
			blocksList[blockId-1].Block.Statements[k] = parser.Statement{Key: k, Value: v}
		}
	}
	name := "nginx.chaosblade.tmp.conf"
	err = config.EasyDumpToFile(name)
	if err != nil {
		return "", spec.ReturnFail(spec.OsCmdExecFailed, err.Error())
	}
	return name, nil
}

func (ng *NginxConfigExecutor) stop(ctx context.Context, dir, activeFile, backup string, model *spec.ExpModel) *spec.Response {
	mode := model.ActionFlags["mode"]
	if mode != "" {
		return spec.ReturnFail(spec.OsCmdExecFailed, fmt.Sprintf("--mode cannot be %s when destroying Nginx config experiment", mode))
	}
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

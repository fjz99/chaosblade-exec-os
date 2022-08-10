package nginx

import (
	"context"
	"fmt"
	"github.com/chaosblade-io/chaosblade-exec-os/exec/nginx/parser"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/chaosblade-io/chaosblade-exec-os/exec/category"
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
	"github.com/chaosblade-io/chaosblade-spec-go/util"
)

const (
	NginxConfigBin = "chaos_nginxconfig"
	fileMode       = "file"
	cmdMode        = "cmd"
)

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

# Change 'server' (assuming block id = 3) exposed on port 8899
blade create nginx config --mode cmd --block-id 3 --set-config='listen=8899'

# Set 'location /' (assuming block id = 4) proxy_pass to www.baidu.com
blade create nginx config --mode cmd --block-id 4 --set-config='proxy_pass=www.baidu.com'


//!!!!!!
//!!!
# Revert config change, suid = xxx
blade destroy nginx config --suid

# Revert config change to the oldest config file
blade destroy nginx config --force
//!!!
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
	// } // nginx process
	if response := testNginxExists(ng.channel, ctx); response != nil {
		return response
	}

	_, activeFile, _, response := getNginxConfigLocation(ng.channel, ctx)
	if response != nil {
		return response
	}

	if _, ok := spec.IsDestroy(ctx); ok {
		return ng.stop(ctx, model)
	}
	return ng.start(ctx, activeFile, model)
}

func (ng *NginxConfigExecutor) start(ctx context.Context, activeFile string, model *spec.ExpModel) *spec.Response {
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
			return spec.ResponseFailWithFlags(spec.FileNotExist, fmt.Sprintf("config file '%s'", newFile))
		}
		newFile, _ = filepath.Abs(newFile)
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
		return spec.ResponseFailWithFlags(spec.ParameterInvalid, "--mode", mode, fmt.Sprintf("invalid --mode argument, which must be '%s' or '%s'", fileMode, cmdMode))
	}

	return swapNginxConfig(ng.channel, ctx, newFile, model)
}

func createNewConfig(config *parser.Config, id string, newKV string) (string, *spec.Response) {
	blocksList := config.GetBlocksList()
	blockId, err := strconv.Atoi(id)
	if err != nil || blockId-1 >= len(blocksList) || blockId < 0 {
		return "", spec.ResponseFailWithFlags(spec.ParameterInvalid, "--block-id", id, fmt.Sprintf("--block-id='%s' is not valid, expect %d-%d", id, 0, len(blocksList)))
	}
	for _, kv := range strings.Split(newKV, ";") {
		arr := strings.Split(strings.TrimSpace(kv), "=")
		if newKV == "" || len(arr) != 2 {
			return "", spec.ResponseFailWithFlags(spec.OsCmdExecFailed, "--set-config", newKV, fmt.Sprintf("--set-config='%s' is not valid", newKV))
		}
		k := strings.TrimSpace(arr[0])
		v := strings.TrimSpace(arr[1])
		if blockId == 0 {
			config.Statements = parser.SetStatement(config.Statements, k, v, false)
		} else {
			statements := blocksList[blockId-1].Block.Statements
			blocksList[blockId-1].Block.Statements = parser.SetStatement(statements, k, v, false)
		}
	}
	name := "nginx.chaosblade.tmp.conf"
	err = config.EasyDumpToFile(name)
	if err != nil {
		return "", spec.ReturnFail(spec.OsCmdExecFailed, err.Error())
	}
	return name, nil
}

func (ng *NginxConfigExecutor) stop(ctx context.Context, model *spec.ExpModel) *spec.Response {
	mode := model.ActionFlags["mode"]
	if mode != "" {
		return spec.ResponseFailWithFlags(spec.ParameterInvalid, "--mode", mode, fmt.Sprintf("--mode cannot be %s when destroying Nginx config experiment", mode))
	}
	return reloadNginxConfig(ng.channel, ctx)
}

func (ng *NginxConfigExecutor) SetChannel(channel spec.Channel) {
	ng.channel = channel
}

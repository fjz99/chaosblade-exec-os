package nginx

import (
	"context"
	"fmt"
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
	"github.com/chaosblade-io/chaosblade-spec-go/util"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

const configBackupName = "nginx.conf.chaosblade.back"

type NginxCommandSpec struct {
	spec.BaseExpModelCommandSpec
}

func (*NginxCommandSpec) Name() string {
	return "nginx"
}

func (*NginxCommandSpec) ShortDesc() string {
	return "Nginx experiment"
}

func (*NginxCommandSpec) LongDesc() string {
	return "Nginx experiment"
}

func startNginx(channel spec.Channel, ctx context.Context) *spec.Response {
	return runNginxCommand(channel, ctx, "")
}

//dir, activeFile, backup
func getNginxConfigLocation(channel spec.Channel, ctx context.Context) (string, string, string, *spec.Response) {
	response := runNginxCommand(channel, ctx, "-t")
	if !response.Success {
		return "", "", "", response
	}
	result := response.Result.(string)
	if !strings.Contains(result, "successful") {
		return "", "", "", spec.ReturnFail(spec.OsCmdExecFailed, `your nginx.conf has something wrong, please run 'nginx -t' to test it.`)
	}
	regex := regexp.MustCompile("file (.*) test is successful")
	location := regex.FindStringSubmatch(result)[1]
	dir := location[:strings.LastIndex(location, string(os.PathSeparator))+1]
	return dir, location, dir + configBackupName, nil
}

// nginx.conf may have 'include mime.types;' etc.
func testNginxConfig(channel spec.Channel, ctx context.Context, file, dir string) *spec.Response {
	file, _ = filepath.Abs(file)
	tmpFile := fmt.Sprintf("%snginx_chaosblade_temp_%v.conf", dir, time.Now().Unix())
	response := channel.Run(ctx, fmt.Sprintf("cp %s %s", file, tmpFile), "")
	if !response.Success {
		return response
	}
	response = runNginxCommand(channel, ctx, fmt.Sprintf("-t -c %s", tmpFile))
	_ = channel.Run(ctx, fmt.Sprintf("rm %s", tmpFile), "") //ignore response
	if !response.Success || !strings.Contains(response.Result.(string), "successful") {
		return response
	}
	return nil
}

func parseMultipleKvPairs(newKV string) [][]string {
	if newKV == "" {
		return nil
	}
	var pairs [][]string
	newKV = strings.TrimSpace(newKV)
	newKV = strings.TrimRight(newKV, ";")
	for _, kv := range strings.Split(newKV, ";") {
		arr := strings.Split(strings.TrimSpace(kv), "=")
		if len(arr) != 2 {
			return nil
		}
		k := strings.TrimSpace(arr[0])
		v := strings.TrimSpace(arr[1])
		pairs = append(pairs, []string{k, v})
	}
	return pairs
}

func reloadNginxConfig(channel spec.Channel, ctx context.Context) *spec.Response {
	_, activeFile, backup, response := getNginxConfigLocation(channel, ctx)
	if response != nil {
		return response
	}

	if !util.IsExist(backup) || util.IsDir(backup) {
		return spec.ReturnFail(spec.OsCmdExecFailed, fmt.Sprintf("backup file %s not exists", backup))
	}

	if response := restoreConfigFile(channel, ctx, backup, activeFile); !response.Success  {
		return response
	}
	if response := runNginxCommand(channel, ctx, "-s reload"); !response.Success {
		return response
	}
	return spec.ReturnSuccess("nginx config restored")
}

func swapNginxConfig(channel spec.Channel, ctx context.Context, newFile string, model *spec.ExpModel) *spec.Response {
	dir, activeFile, backup, response := getNginxConfigLocation(channel, ctx)
	if response != nil {
		return response
	}
	if response := testNginxConfig(channel, ctx, newFile, dir); response != nil {
		return response
	}

	if model.ActionFlags["mode"] == cmdMode {
		// remove auto generated config
		if response := backupConfigFile(channel, ctx, backup, activeFile, newFile, true); !response.Success {
			return response
		}
	} else {
		if response := backupConfigFile(channel, ctx, backup, activeFile, newFile, false); !response.Success {
			return response
		}
	}

	if response := runNginxCommand(channel, ctx, "-s reload"); !response.Success {
		return response
	}
	return spec.ReturnSuccess("nginx config changed")
}

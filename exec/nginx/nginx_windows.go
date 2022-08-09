package nginx

import (
	"context"
	"fmt"
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
	"github.com/chaosblade-io/chaosblade-spec-go/util"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func startDaemon(ctx context.Context, script, args string) *spec.Response {
	cmd := exec.CommandContext(ctx, "cmd", "/C", script+` `+args)
	if err := cmd.Start(); err != nil {
		return spec.ResponseFailWithFlags(spec.OsCmdExecFailed, cmd, err)
	} else {
		return spec.Success()
	}
}

// nginx.conf may have 'include mime.types;' etc.
func testNginxConfig(channel spec.Channel, ctx context.Context, file, dir string) *spec.Response {
	file, _ = filepath.Abs(file)
	tmpFile := fmt.Sprintf("%snginx_chaosblade_temp_%v.conf", dir, time.Now().Unix())
	response := channel.Run(ctx, fmt.Sprintf(`copy /Y %s %s`, file, tmpFile), "")
	if !response.Success {
		return response
	}
	response = runNginxCommand(channel, ctx, fmt.Sprintf("-t -c %s", tmpFile))
	_ = channel.Run(ctx, fmt.Sprintf("del %s", tmpFile), "") //ignore response
	if !response.Success || !strings.Contains(response.Result.(string), "successful") {
		return response
	}
	return nil
}

func testNginxExists(channel spec.Channel, ctx context.Context) *spec.Response {
	response := channel.Run(ctx, "tasklist", "")
	if response.Success {
		processes := response.Result.(string)
		if strings.Contains(processes, "nginx.exe") {
			return nil
		} else {
			return spec.ReturnFail(spec.ProcessIdByNameFailed, "cannot find nginx process")
		}
	} else {
		return response
	}
}

func killNginx(channel spec.Channel, ctx context.Context) *spec.Response {
	response := channel.Run(ctx, "taskkill /im nginx.exe /F", "")
	if !response.Success {
		return response
	} else {
		return nil
	}
}

func runNginxCommand(channel spec.Channel, ctx context.Context, args string) *spec.Response {
	//find nginx location, NGINX_HOME or PATH
	dir := os.Getenv("NGINX_HOME")
	if dir == "" {
		//loc := os.Getenv("PATH")
		//TODO PATH
		return spec.ReturnFail(spec.OsCmdExecFailed, "cannot find nginx location, check your PATH and NGINX_HOME")
	}
	if args == "" {
		//start nginx daemon
		return startDaemon(ctx, fmt.Sprintf("cd /d %s && nginx", dir), "")
	}
	return channel.Run(ctx, fmt.Sprintf("cd /d %s && nginx %s", dir, args), "")
}

func restoreConfigFile(channel spec.Channel, ctx context.Context, backup, activeFile string) *spec.Response {
	return channel.Run(ctx, fmt.Sprintf("move /Y %s %s", backup, activeFile), "")
}

func backupConfigFile(channel spec.Channel, ctx context.Context, backup string, activeFile string, newFile string, remove bool) *spec.Response {
	cmd := ""
	if util.IsExist(backup) {
		//don't create backup
		cmd = fmt.Sprintf("copy /Y %s %s", newFile, activeFile)
	} else {
		cmd = fmt.Sprintf("copy %s %s && copy /Y %s %s", activeFile, backup, newFile, activeFile)
	}
	if remove {
		cmd += fmt.Sprintf(" && del %s", newFile)
	}
	return channel.Run(ctx, cmd, "")
}

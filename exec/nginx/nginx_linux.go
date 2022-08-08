package nginx

import (
	"context"
	"fmt"
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
	"github.com/chaosblade-io/chaosblade-spec-go/util"
	"strconv"
	"strings"
)

func NewNginxCommandSpec() spec.ExpModelCommandSpec {
	return &NginxCommandSpec{
		spec.BaseExpModelCommandSpec{
			ExpActions: []spec.ExpActionCommandSpec{
				NewCrashActionSpec(),
				NewRestartActionSpec(),
				NewConfigActionSpec(),
				NewResponseActionSpec(),
			},
			ExpFlags: []spec.ExpFlagSpec{},
		},
	}
}

func testNginxExists(channel spec.Channel, ctx context.Context) *spec.Response {
	_, response := getNginxPid(channel, ctx)
	if response != nil {
		return response
	}
	return nil
}

func killNginx(channel spec.Channel, ctx context.Context) *spec.Response {
	commands := []string{"kill"}
	if response, ok := channel.IsAllCommandsAvailable(ctx, commands); !ok {
		return response
	}

	allPid, response := getNginxPid(channel, ctx)
	if response != nil {
		return response
	}
	// fmt.Println(allPid)
	for _, pid := range allPid {
		response = channel.Run(ctx, fmt.Sprintf("kill -9 %d", pid), "")
		// fmt.Println(response)
		if !response.Success {
			return response
		}
	}
	return response
}

func runNginxCommand(channel spec.Channel, ctx context.Context, args string) *spec.Response {
	return channel.Run(ctx, "nginx", args)
}

func getNginxPid(channel spec.Channel, ctx context.Context) ([]int, *spec.Response) {
	response := channel.Run(ctx,
		`ps aux | grep -v grep | egrep 'nginx: ' | awk '{print $2}'`, "")
	if !response.Success {
		return []int{}, response
	}
	result := response.Result.(string)
	count := strings.Count(result, "\n")
	if count == 0 {
		return []int{}, spec.ReturnFail(spec.ProcessIdByNameFailed, "cannot find nginx process")
	}
	var allPid []int
	for _, s := range strings.Split(strings.Trim(result, "\n"), "\n") {
		pid, err := strconv.Atoi(s)
		if err != nil {
			return []int{}, spec.ReturnFail(spec.ProcessIdByNameFailed, "cannot find nginx process")
		}
		allPid = append(allPid, pid)
	}

	return allPid, nil
}

func restoreConfigFile(channel spec.Channel, ctx context.Context, backup, activeFile string) *spec.Response {
	return channel.Run(ctx, fmt.Sprintf("mv -f %s %s", backup, activeFile), "")
}

func backupConfigFile(channel spec.Channel, ctx context.Context, backup string, activeFile string, newFile string, remove bool) *spec.Response {
	cmd := ""
	if util.IsExist(backup) {
		//don't create backup
		cmd = fmt.Sprintf("cp -f %s %s", newFile, activeFile)
	} else {
		cmd = fmt.Sprintf("cp %s %s && cp -f %s %s", activeFile, backup, newFile, activeFile)
	}
	if remove {
		cmd += fmt.Sprintf(" && rm %s", newFile)
	}
	return channel.Run(ctx, cmd, "")
}

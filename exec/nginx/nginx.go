package nginx

import (
	"context"
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
	"regexp"
	"strconv"
	"strings"
)

const BurnCpuBin = "chaos_burncpu"

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

func getNginxConfigLocation(channel spec.Channel, ctx context.Context) (string, *spec.Response) {
	response := channel.Run(ctx, `nginx -t`, "")
	if !response.Success {
		return "", response
	}
	result := response.Result.(string)
	if !strings.Contains(result, "successful") {
		return "", spec.ReturnFail(spec.OsCmdExecFailed, `your nginx.conf has something wrong, please run 'nginx -t' to test it.`)
	}
	regex := regexp.MustCompile("file (.*) test is successful")
	location := regex.FindStringSubmatch(result)[1]
	//location = location[:strings.LastIndex(location, "/")]
	return location, nil
}

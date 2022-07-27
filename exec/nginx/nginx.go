package nginx

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/chaosblade-io/chaosblade-spec-go/spec"
)

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

//dir, activeFile
func getNginxConfigLocation(channel spec.Channel, ctx context.Context) (string, string, *spec.Response) {
	response := channel.Run(ctx, `nginx -t`, "")
	if !response.Success {
		return "", "", response
	}
	result := response.Result.(string)
	if !strings.Contains(result, "successful") {
		return "", "", spec.ReturnFail(spec.OsCmdExecFailed, `your nginx.conf has something wrong, please run 'nginx -t' to test it.`)
	}
	regex := regexp.MustCompile("file (.*) test is successful")
	location := regex.FindStringSubmatch(result)[1]
	dir := location[:strings.LastIndex(location, string(os.PathSeparator))+1]
	//location = location[:strings.LastIndex(location, "/")]
	return dir, location, nil
}

// nginx.conf may have 'include mime.types;' etc.
func testNginxConfig(channel spec.Channel, ctx context.Context, file, dir string) *spec.Response {
	file, _ = filepath.Abs(file)
	tmpFile := fmt.Sprintf("%snginx_chaosblade_temp_%v.conf", dir, time.Now().Unix())
	response := channel.Run(ctx, fmt.Sprintf("cp %s %s && nginx -t -c %s", file, tmpFile, tmpFile), "")
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
	pairs := [][]string{}
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

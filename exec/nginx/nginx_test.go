package nginx

import (
	"context"
	"fmt"
	"github.com/chaosblade-io/chaosblade-spec-go/channel"
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
	"regexp"
	"strings"
	"testing"
)

func TestPid(t *testing.T) {
	localChannel := channel.LocalChannel{}
	response := localChannel.Run(context.Background(),
		`ps aux | grep -v grep | egrep 'nginx: master' | awk '{print $2}'`, "")
	fmt.Println(response)
}

func TestCmd(t *testing.T) {
	localChannel := channel.LocalChannel{}
	response := localChannel.Run(context.Background(),
		`ps aux | grep -v grep | egrep -o 'nginx: master.*' | egrep -o ' [^ ]*nginx.*'`, "")
	fmt.Println(response)
}

func TestRegex(t *testing.T) {
	regex := regexp.MustCompile("file (.*) test is successful")
	location := regex.FindStringSubmatch(`nginx: the configuration file /etc/nginx/nginx.conf syntax is ok
	nginx: configuration file /etc/nginx/nginx.conf test is successful`)[1]
	location = location[:strings.LastIndex(location, "/")]
	fmt.Println(location)

	pid, response := getNginxPid(channel.NewLocalChannel(), context.Background())
	fmt.Println(pid, response)

	loc, res := getNginxConfigLocation(channel.NewLocalChannel(), context.Background())
	fmt.Println(loc, res)
}

func TestCrash(t *testing.T) {
	executor := NginxCrashExecutor{channel: channel.NewLocalChannel()}
	model := spec.ExpModel{}
	response := executor.Exec("", context.Background(), &model)
	fmt.Println(*response)

	//cancel
	response = executor.Exec("dsadsad2", context.WithValue(context.Background(), "suid", "dasdsa"), &model)
	fmt.Println(*response)
}

func TestRestart(t *testing.T) {
	executor := NginxRestartExecutor{channel: channel.NewLocalChannel()}
	model := spec.ExpModel{}
	response := executor.Exec("", context.Background(), &model)
	fmt.Println(*response)

	//cancel
	response = executor.Exec("dsadsad2", context.WithValue(context.Background(), "suid", "dasdsa"), &model)
	fmt.Println(*response)
}

func TestConfigChange(t *testing.T) {
	s := NewConfigActionSpec()
	executor := s.Executor()
	executor.SetChannel(channel.NewLocalChannel())
	model := spec.ExpModel{}
	model.ActionFlags = make(map[string]string)
	// model.ActionFlags["file"] = "conf/ok.conf"
	model.ActionFlags["file"] = "conf/wrong.conf"
	response := executor.Exec("", context.Background(), &model)
	fmt.Println(*response)

	//cancel
	// response = executor.Exec("dsadsad2", context.WithValue(context.Background(), "suid", "dasdsa"), &model)
	// fmt.Println(*response)
}

func TestConfigChangeRevert(t *testing.T) {
	s := NewConfigActionSpec()
	executor := s.Executor()
	executor.SetChannel(channel.NewLocalChannel())
	model := spec.ExpModel{}
	model.ActionFlags = make(map[string]string)

	//cancel
	response := executor.Exec("dsadsad2", context.WithValue(context.Background(), "suid", "dasdsa"), &model)
	fmt.Println(*response)
}

func TestListBlock(t *testing.T) {
	s := NewConfigActionSpec()
	executor := s.Executor()
	executor.SetChannel(channel.NewLocalChannel())
	model := spec.ExpModel{}
	model.ActionFlags = make(map[string]string)
	model.ActionFlags["list"] = "true"

	//cancel
	response := executor.Exec("dsadsad2", context.Background(), &model)
	fmt.Println(*response)
}

func TestKVChange(t *testing.T) {
	s := NewConfigActionSpec()
	executor := s.Executor()
	executor.SetChannel(channel.NewLocalChannel())
	model := spec.ExpModel{}
	model.ActionFlags = make(map[string]string)
	// model.ActionFlags["list"] = "true"
	model.ActionFlags["set-config"]="a=b;c=d"
	model.ActionFlags["block-id"]="0"

	//cancel
	response := executor.Exec("dsadsad2", context.Background(), &model)
	fmt.Println(*response)
}

func TestTmp(t *testing.T) {

}

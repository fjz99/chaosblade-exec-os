//go:build windows
// +build windows

package nginx

import (
	"context"
	"fmt"
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
	"regexp"
	"strings"
	"testing"
)

func TestRegex(t *testing.T) {
	regex := regexp.MustCompile("file (.*) test is successful")
	location := regex.FindStringSubmatch(`nginx: the configuration file /etc/nginx/nginx.conf syntax is ok
	nginx: configuration file /etc/nginx/nginx.conf test is successful`)[1]
	location = location[:strings.LastIndex(location, "/")]
	fmt.Println(location)

	response := testNginxExists(NewCmdChannel(), context.Background())
	fmt.Println(response)

	dir, loc, backup, res := getNginxConfigLocation(NewCmdChannel(), context.Background())
	fmt.Println(dir, loc, backup, res)
}

func TestCrash(t *testing.T) {
	executor := NginxCrashExecutor{channel: NewCmdChannel()}
	model := spec.ExpModel{}
	response := executor.Exec("", context.Background(), &model)
	fmt.Println(*response)

	//cancel
	//response := executor.Exec("dsadsad2", context.WithValue(context.Background(), "suid", "dasdsa"), &model)
	//fmt.Println(*response)
}

func TestRestart(t *testing.T) {
	executor := NginxRestartExecutor{channel: NewCmdChannel()}
	model := spec.ExpModel{}
	response := executor.Exec("", context.Background(), &model)
	fmt.Println(*response)

	//cancel
	// response := executor.Exec("dsadsad2", context.WithValue(context.Background(), "suid", "dasdsa"), &model)
	// fmt.Println(*response)
}

func TestConfigChange(t *testing.T) {
	s := NewConfigActionSpec()
	executor := s.Executor()
	executor.SetChannel(NewCmdChannel())
	model := spec.ExpModel{}
	model.ActionFlags = make(map[string]string)
	model.ActionFlags["mode"] = "file"
	model.ActionFlags["file"] = "conf/ok.conf"
	// model.ActionFlags["file"] = "conf/wrong.conf"
	response := executor.Exec("", context.Background(), &model)
	fmt.Println(*response)
}

func TestConfigChangeRevert(t *testing.T) {
	s := NewConfigActionSpec()
	executor := s.Executor()
	executor.SetChannel(NewCmdChannel())
	model := spec.ExpModel{}
	model.ActionFlags = make(map[string]string)

	//cancel
	response := executor.Exec("dsadsad2", context.WithValue(context.Background(), "suid", "dasdsa"), &model)
	fmt.Println(*response)
}

func TestListBlock(t *testing.T) {
	s := NewConfigActionSpec()
	executor := s.Executor()
	executor.SetChannel(NewCmdChannel())
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
	executor.SetChannel(NewCmdChannel())
	model := spec.ExpModel{}
	model.ActionFlags = make(map[string]string)
	// model.ActionFlags["list"] = "true"

	model.ActionFlags["mode"] = "xxx"

	// model.ActionFlags["set-config"] = "listen=9999"
	// model.ActionFlags["block-id"] = "3"

	model.ActionFlags["set-config"] = "proxy_pass=https://www.taobao.com"
	model.ActionFlags["block-id"] = "4"

	response := executor.Exec("dsadsad2", context.Background(), &model)
	fmt.Println(response)
}

func TestCancelKVChange(t *testing.T) {
	s := NewConfigActionSpec()
	executor := s.Executor()
	executor.SetChannel(NewCmdChannel())
	model := spec.ExpModel{}
	model.ActionFlags = make(map[string]string)
	// model.ActionFlags["mode"] = "file"

	response := executor.Exec("dsadsad2", context.WithValue(context.Background(), "suid", "dasdsa"), &model)
	fmt.Println(response)
}

func TestChangeResponse(t *testing.T) {
	s := NewResponseActionSpec()
	executor := s.Executor()
	executor.SetChannel(NewCmdChannel())
	model := spec.ExpModel{}
	model.ActionFlags = make(map[string]string)
	model.ActionFlags["type"] = "json"
	model.ActionFlags["path"] = "/test"
	model.ActionFlags["code"] = "200"
	model.ActionFlags["header"] = ""
	model.ActionFlags["body"] = `{"a":1}`
	// model.ActionFlags["body"] = "hello!"

	response := executor.Exec("dsadsad2", context.Background(), &model)
	fmt.Println(response)
}

func TestCancelResponseChange(t *testing.T) {
	s := NewResponseActionSpec()
	executor := s.Executor()
	executor.SetChannel(NewCmdChannel())
	model := spec.ExpModel{}
	model.ActionFlags = make(map[string]string)

	response := executor.Exec("dsadsad2", context.WithValue(context.Background(), "suid", "dasdsa"), &model)
	fmt.Println(response)
}

package nginx

import (
	"context"
	"fmt"
	"github.com/chaosblade-io/chaosblade-spec-go/log"
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
	"github.com/chaosblade-io/chaosblade-spec-go/util"
	"github.com/shirou/gopsutil/process"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type PowerShellChannel struct {
}

// NewPowerShellChannel returns a local channel for invoking the host command
func NewPowerShellChannel() spec.Channel {
	return &PowerShellChannel{}
}

func (l *PowerShellChannel) Name() string {
	return "powershell"
}

func (l *PowerShellChannel) Run(ctx context.Context, script, args string) *spec.Response {
	return execScript(ctx, script, args)
}

func (l *PowerShellChannel) GetScriptPath() string {
	return util.GetProgramPath()
}

func (l *PowerShellChannel) GetPidsByProcessCmdName(processName string, ctx context.Context) ([]string, error) {
	return nil, nil
}

func (l *PowerShellChannel) GetPidsByProcessName(processName string, ctx context.Context) ([]string, error) {
	return nil, nil
}

func (l *PowerShellChannel) GetPsArgs(ctx context.Context) string {
	var psArgs = "-eo user,pid,ppid,args"
	if l.IsAlpinePlatform(ctx) {
		psArgs = "-o user,pid,ppid,args"
	}
	return psArgs
}

func (l *PowerShellChannel) IsAlpinePlatform(ctx context.Context) bool {
	var osVer = ""
	if util.IsExist("/etc/os-release") {
		response := l.Run(ctx, "awk", "-F '=' '{if ($1 == \"ID\") {print $2;exit 0}}' /etc/os-release")
		if response.Success {
			osVer = response.Result.(string)
		}
	}
	return strings.TrimSpace(osVer) == "alpine"
}

func (l *PowerShellChannel) IsAllCommandsAvailable(ctx context.Context, commandNames []string) (*spec.Response, bool) {
	return nil, true
}

func (l *PowerShellChannel) IsCommandAvailable(ctx context.Context, commandName string) bool {
	response := l.Run(ctx, "command", fmt.Sprintf("-v %s", commandName))
	return response.Success
}

func (l *PowerShellChannel) ProcessExists(pid string) (bool, error) {
	p, err := strconv.Atoi(pid)
	if err != nil {
		return false, err
	}
	return process.PidExists(int32(p))
}

func (l *PowerShellChannel) GetPidUser(pid string) (string, error) {
	p, err := strconv.Atoi(pid)
	if err != nil {
		return "", err
	}
	process, err := process.NewProcess(int32(p))
	if err != nil {
		return "", err
	}
	return process.Username()
}

func (l *PowerShellChannel) GetPidsByLocalPorts(ctx context.Context, localPorts []string) ([]string, error) {
	if localPorts == nil || len(localPorts) == 0 {
		return nil, fmt.Errorf("the local port parameter is empty")
	}
	var result = make([]string, 0)
	for _, port := range localPorts {
		pids, err := l.GetPidsByLocalPort(ctx, port)
		if err != nil {
			return nil, fmt.Errorf("failed to get pid by %s, %v", port, err)
		}
		log.Infof(ctx, "get pids by %s port returns %v", port, pids)
		if pids != nil && len(pids) > 0 {
			result = append(result, pids...)
		}
	}
	return result, nil
}

func (l *PowerShellChannel) GetPidsByLocalPort(ctx context.Context, localPort string) ([]string, error) {
	return nil, nil
}

// execScript invokes exec.CommandContext
func execScript(ctx context.Context, script, args string) *spec.Response {
	isBladeCommand := isBladeCommand(script)
	if isBladeCommand && !util.IsExist(script) {
		return spec.ResponseFailWithFlags(spec.ChaosbladeFileNotFound, script)
	}
	newCtx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	if ctx == context.Background() {
		ctx = newCtx
	}
	log.Debugf(ctx, "Command: %s %s", script, args)
	cmd := exec.CommandContext(ctx, "powershell", "-c", script+" "+args)
	output, err := cmd.CombinedOutput()
	outMsg := string(output)
	log.Debugf(ctx, "Command Result, output: %v, err: %v", outMsg, err)
	if strings.TrimSpace(outMsg) != "" {
		resp := spec.Decode(outMsg, nil)
		if resp.Code != spec.ResultUnmarshalFailed.Code {
			return resp
		}
	}
	if err == nil {
		return spec.ReturnSuccess(outMsg)
	}
	outMsg += " " + err.Error()
	return spec.ResponseFailWithFlags(spec.OsCmdExecFailed, cmd, outMsg)
}

func isBladeCommand(script string) bool {
	return strings.HasSuffix(script, util.GetProgramPath())
}

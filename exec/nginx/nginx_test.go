package nginx

import (
	"context"
	"fmt"
	"github.com/chaosblade-io/chaosblade-spec-go/channel"
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

package nginx

import (
	"context"
	"fmt"
	"testing"
)

//test channel, PowerShell

//TestWindows
func TestWindows(t *testing.T) {
	localChannel := NewPowerShellChannel()
	response := localChannel.Run(context.Background(),
		`Start-Process -FilePath D:\nginx-1.9.9\nginx -WorkingDirectory D:\nginx-1.9.9 -ArgumentList "-t" -NoNewWindow`, "")
	fmt.Println(response)
}

func Test2(t *testing.T) {
	localChannel := NewPowerShellChannel()
	location, s, b, response := getNginxConfigLocation(localChannel, context.Background())
	fmt.Println(location, s, b, *response)
}

package telegraf

import (
	"fmt"
	"testing"
	"time"
)

func TestStartX(t *testing.T) {
	go Start()
	time.Sleep(30 * time.Second)
	fmt.Printf("\n ============== 重启 ============== \n")
	Restart()
	select {}
}

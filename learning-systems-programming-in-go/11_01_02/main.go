package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Printf("プロセスID: %d\n", os.Getpid())
	fmt.Printf("親プロセスID: %d\n", os.Getppid())

	// プロセス確認する時間を確保するためにSleepを入れる
	time.Sleep(time.Minute * 3)
}

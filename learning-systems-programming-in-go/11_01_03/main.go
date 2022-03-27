package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func main() {
	sid, _ := syscall.Getsid(os.Getpid())
	fmt.Fprintf(os.Stderr, " グループID: %d セッションID: %d\n",
		syscall.Getpgrp(), sid)

	// プロセス確認する時間を確保するためにSleepを入れる
	time.Sleep(time.Minute * 3)
}

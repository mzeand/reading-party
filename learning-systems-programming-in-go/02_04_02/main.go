package main

import (
	"os"
)

func main() {
	os.Stdout.Write([]byte("こんにちは、世界!!\n"))
}

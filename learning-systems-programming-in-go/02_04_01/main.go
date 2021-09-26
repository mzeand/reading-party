package main

import (
	"os"
)

func main() {
	file, err := os.Create("text.txt")
	if err != nil {
		panic(nil)
	}
	file.Write([]byte("こんにちは、世界\n"))
	file.Close()
}

package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer
	buffer.Write([]byte("こんにちは、世界!!\n"))
	fmt.Println(buffer.String())
}

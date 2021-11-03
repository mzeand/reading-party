package main

import (
	"bufio"
	"os"
)

func main() {
	buffer := bufio.NewWriter(os.Stdout)
	buffer.WriteString("buffer.Writer ")
	buffer.Flush()
	buffer.WriteString("example\n")
	buffer.Flush()
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	var builer strings.Builder
	builer.Write([]byte("string.Builder example\n"))
	fmt.Println(builer.String())
}

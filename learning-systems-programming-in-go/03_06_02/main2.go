package main

import (
	"fmt"
	"strings"
)

var source = "123, 1.123, 1.0e4, test"

func main() {
	reader := strings.NewReader(source)
	var i int
	var f, g float64
	var s string
	fmt.Fscanf(reader, "%v, %v, %v, %v", &i, &f, &g, &s)
	fmt.Printf("i=%#v f=%#f g=%#v s=%#v\n", i, f, g, s)
}

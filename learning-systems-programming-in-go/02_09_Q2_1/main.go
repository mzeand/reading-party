package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %d\n", 1000)
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %s\n", time.Now().Month())
	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %f\n", 10.0/1.0)
}

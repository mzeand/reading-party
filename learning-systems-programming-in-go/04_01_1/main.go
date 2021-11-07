package main

import (
	"fmt"
	"time"
)

func sub() {
	fmt.Println("sub() is runnning")
	time.Sleep(time.Second)
	fmt.Printf("sub() is finished")
}

func main() {
	fmt.Println("start sub()")
	go sub()
	time.Sleep(2 * time.Second)
}

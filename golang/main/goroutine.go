package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		panic("just panic")

	}()

	time.Sleep(10)
	if r := recover(); r != nil {
		fmt.Println("recovered")
	}
}

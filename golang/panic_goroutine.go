package golang

import (
	"fmt"
	"time"
)

func PanicAndGoroutine() {
	go func(para interface{}) {
		str := para.(string)
		fmt.Printf("str is %v\n", str)
	}(1)

	time.Sleep(2 * time.Second)
	fmt.Println("main goroutine is fine!")
}

func PanicAndGorWithDpeth() {

	go func() {
		go func(para interface{}) {
			str := para.(string)
			fmt.Printf("str is %v\n", str)
		}(1)
		time.Sleep(2 * time.Second)
		fmt.Println("child goroutine is fine!")
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("main goroutine is fine!")
}

func Goroutines() {
	go func() {
		defer fmt.Println("child go routine is out")
		for {
			fmt.Println("child go routine is running")
			time.Sleep(2 * time.Second)
		}
	}()
	time.Sleep(20 * time.Second)
}

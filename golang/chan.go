package golang

import (
	"fmt"
	"math/rand"
	"time"
)

func DestroyGoroutine() {
	rand.Seed(time.Now().Unix())

	go func() {
		fmt.Println("this msg is from a child goroutine")
		go func() {
			for {
				fmt.Println("this msg is from a child child goroutine")
				ranndom := rand.Intn(5)
				time.Sleep(time.Duration(ranndom) * time.Second)
			}
		}()

	}()
	fmt.Println("child go routine has been destroied")
	time.Sleep(30 * time.Second)
}

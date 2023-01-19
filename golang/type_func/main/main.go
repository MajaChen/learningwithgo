package main

import (
	"fmt"
	"learning/golang/type_func"
)

func test(f type_func.Factory) {
	// do nothing
	fmt.Println(f.B(1.0))
}

func main() {
	i := type_func.NewImpl(type_func.WithImplB(func(f float32) float32 {
		return f + 1
	}))
	test(i)
}

package golang

import "fmt"

func Demo() {
	var i interface{} = nil
	fmt.Println(i)

	arr := []int{
		1, 2, 3}
	fmt.Println(arr)
}

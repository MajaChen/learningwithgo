package template

import (
	"bufio"
	"fmt"
	"os"
)

func ScanDemo() {
	//读取单个字符
	var a1 int
	fmt.Scan(&a1)
	// 读取多个字符到数组
	nums := make([]int, 0, a1)
	for i := 0; i < a1; i++ {
		var j int
		fmt.Scan(&j)
		nums = append(nums, j)
	}
	fmt.Scanln()

	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')

	fmt.Print(str)
}

// Launch 相当于main
func Launch() {
	ScanDemo()
}

func InputDemo() {
	//reader := bufio.NewReader(os.Stdin)
	//str,_ := reader.ReadString('\n')

}

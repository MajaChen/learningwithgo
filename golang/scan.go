package golang

import (
	"bufio"
	"fmt"
	"os"
)

func Scan() error {
	var m, n int
	var i, j, k int

	if _, err := fmt.Scanln(&m, &n); err != nil {
		fmt.Errorf("err is %v\n", err)
		return err
	}

	if _, err := fmt.Scanln(&i, &j, &k); err != nil {
		fmt.Errorf("err is %v\n", err)
		return err
	}

	fmt.Printf("inputs are %v,%v\n", m, n)
	fmt.Printf("inputs are %v,%v,%v\n", i, j, k)
	return nil
}

func Scan2() error {
	var j int
	for i := 0; i < 5; i++ {
		if _, err := fmt.Scan(&j); err != nil {
			return err
		}
		fmt.Printf("j is %v\n", j)
	}
	return nil
}

func Reader() error {
	reader := bufio.NewReader(os.Stdin)
	var text string
	var err error
	if text, err = reader.ReadString('\n'); err != nil {
		return err
	}
	fmt.Printf("text is %v", text)
	return nil
}

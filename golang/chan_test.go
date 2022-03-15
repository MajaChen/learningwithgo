package golang

import (
	"testing"
)

func TestChan(t *testing.T) {
	testChan := make(chan int)
	close(testChan)
	go func() {
		for {
			select {
			case i := <-testChan:
				t.Errorf("i is %v", i)
			}
		}
	}()
}

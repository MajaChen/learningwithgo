package golang

import "fmt"

func SliceInFunction(nums []int) {
	nums[0] += 1
}

func SliceDemo() {
	nums := make([]int, 0, 2)
	nums[0] = 0
	nums[1] = 1

	fmt.Printf("nums is %v", nums)
}

func ArrayAndSlice() {
	s1 := []int{1, 2, 3}
	s2 := s1[1:]
	for i := range s2 {
		s2[i] += 10
	}

	fmt.Println(s2)
	s2 = append(s2, 4)
	for i := range s2 {
		s2[i] += 10
	}
	fmt.Println(s2)
}

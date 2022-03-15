package leetcode

import "fmt"

func myPow(x float64, n int) float64 {

	if x == 0 {
		return float64(0)
	}

	var flagX, flagN bool
	if x < 0 {
		flagX = true
		x = 0 - x
	}
	if n < 0 {
		flagN = true
		n = 0 - n
	}

	multi, an := float64(1), n
	for count := 1; an > 0; {
		aMulti := float64(1)
		for ax := x; count <= an; count *= 2 {
			aMulti *= ax
			ax = aMulti
		}
		multi *= aMulti
		an -= count / 2
		count = 1
	}

	if flagX {
		if n%2 != 0 {
			multi = 0 - multi
		}
	}

	if flagN {
		return 1 / multi
	}
	return multi
}

var reached bool
var reachMapping []int

func canJumpRe(startIndex int, nums []int) int {

	if reachMapping[startIndex] != 0 {
		return reachMapping[startIndex]
	}

	span := nums[startIndex]
	if span+startIndex >= len(nums)-1 {
		reached = true
		reachMapping[startIndex] = 1
		return 1
	}

	for i := startIndex + 1; i <= startIndex+span && !reached; i++ {
		if reachMapping[i] < 0 {
			continue
		}
		if j := canJumpRe(i, nums); j > 0 {
			reachMapping[startIndex] = 1
			return 1
		} else {
			reachMapping[i] = -1
		}
	}
	reachMapping[startIndex] = -1
	return -1
}

func canJump(nums []int) bool {
	reached = false
	reachMapping = make([]int, len(nums))
	canJumpRe(0, nums)
	fmt.Printf("%v", reachMapping)
	return reached
}

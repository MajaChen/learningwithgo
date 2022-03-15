package leetcode

import (
	"math"
	"sort"
)

func singleNumberV2(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res = res ^ nums[i]
	}
	return res
}

type dividePair struct {
	num   int
	count int
}

func divide2(dividend int, divisor int) int {

	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	var flag1, flag2 int
	if dividend < 0 {
		dividend = 0 - dividend
		flag1 = 1
	}
	if divisor < 0 {
		divisor = 0 - divisor
		flag2 = 1
	}

	count := 0
	for dividend >= divisor {
		var pair dividePair
		initCount, initDivisor := 1, divisor
		for dividend < initDivisor { // 指数式增长
			pair = dividePair{num: initDivisor, count: initCount}
			initDivisor += initDivisor
			initCount += initCount
		}
		count += pair.count
		dividend -= pair.num
	}

	if flag1^flag2 == 0 {
		return count
	} else {
		return 0 - count
	}
}

type permutationPair struct {
	leftIndex  int
	rightIndex int
}

func nextPermutation(nums []int) {
	lastPair := permutationPair{leftIndex: -1, rightIndex: -1}
	for i := 0; i < len(nums)-1; i++ {
		pair := permutationPair{leftIndex: -1, rightIndex: -1}
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i] {
				pair.leftIndex = i
				if pair.rightIndex < 0 || nums[pair.rightIndex] > nums[j] {
					pair.rightIndex = j
				}
			}
		}
		if pair.leftIndex > lastPair.leftIndex {
			lastPair.leftIndex = pair.leftIndex
			lastPair.rightIndex = pair.rightIndex
		}

	}
	if lastPair.leftIndex < 0 {
		sort.Ints(nums)
	} else { // "交换"leftIndex与rightIndex处的值
		t := nums[lastPair.rightIndex]
		for i := lastPair.rightIndex - 1; i >= lastPair.leftIndex; i-- {
			nums[i+1] = nums[i]
		}
		nums[lastPair.leftIndex] = t
		sort.Ints(nums[lastPair.leftIndex+1:])
	}
}

package leetcode

import (
	"sort"
	"strconv"
)

func reverseBits(num uint32) uint32 {

	str := strconv.FormatInt(int64(num), 2)
	for i := len(str); i <= 32; i++ {
		str = "0" + str
	}
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; {
		k := runes[i]
		runes[i] = runes[j]
		runes[j] = k
		i++
		j--
	}
	i, _ := strconv.ParseUint(string(runes), 2, 64)
	return uint32(i)
}

type interval [][]int

func (interval interval) Len() int {
	return len(interval)
}

func (interval interval) Less(i, j int) bool {
	return interval[i][0]-interval[j][0] < 0
}

func (interval interval) Swap(i, j int) {
	nums := interval[i]
	interval[i] = interval[j]
	interval[j] = nums
	return
}

func canMerge(inteA, inteB []int) (newInte []int, flag bool) {
	newInte, flag = make([]int, 2), true
	if inteA[0] == inteB[0] {
		newInte[0] = inteA[0]
		if newInte[1] = inteA[1]; inteA[1] < inteB[1] {
			newInte[1] = inteB[1]
		}
	} else if inteA[0] < inteB[0] {
		if inteA[1] >= inteB[0] {
			newInte[0] = inteA[0]
			if newInte[1] = inteA[1]; inteA[1] < inteB[1] {
				newInte[1] = inteB[1]
			}
		} else {
			flag = false
		}
	}
	return
}

func merge2(intervals [][]int) [][]int {
	sort.Sort(interval(intervals))
	mergedIntervals := make([][]int, 0)
	target, i := intervals[0], 1
	for ; i < len(intervals); i++ {
		if newInte, canMerge := canMerge(target, intervals[i]); canMerge {
			target = newInte
		} else {
			mergedIntervals = append(mergedIntervals, target)
			target = intervals[i]
		}
	}
	mergedIntervals = append(mergedIntervals, target)
	return mergedIntervals
}

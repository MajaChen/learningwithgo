package leetcode

//leetcode submit region begin(Prohibit modification and deletion)

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

func insert(intervals [][]int, newInterval []int) [][]int {

	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	newIntervals, flag := make([][]int, 0, len(intervals)), true
	for index, interval := range intervals {
		if newInterval[1] < interval[0] {
			newIntervals = append(newIntervals, newInterval)
			newIntervals = append(newIntervals, intervals[index:]...)
			flag = false
			break
		} else if newInterval[0] > interval[1] {
			newIntervals = append(newIntervals, interval)
			continue
		}
		newInterval[0], newInterval[1] = min(newInterval[0], interval[0]), max(newInterval[1], interval[1])
	}

	if flag {
		newIntervals = append(newIntervals, newInterval)
	}

	return newIntervals
}

//leetcode submit region end(Prohibit modification and deletion)

package leetcode

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func videoStitching(clips [][]int, time int) int {
	sort.Slice(clips, func(i, j int) bool {
		return clips[i][0] < clips[j][0] ||
			(clips[i][0] == clips[j][0] &&
				clips[i][1] > clips[j][1])
	})
	ans, n := 0, len(clips)
	for i, j := 0, 0; i < time; {
		ans++
		max := -1
		for j < n {
			if clips[j][0] <= i && clips[j][1] >= i {
				if clips[j][1] > max {
					max = clips[j][1]
				}
			} else if clips[j][0] > i{
				break
			}
			j++
		}
		if max < 0 {
			return -1
		}
		i = max
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

package leetcode

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func arrayRankTransform(arr []int) []int {
	arrCopy, mapping := make([]int, len(arr)), make(map[int]int)
	copy(arrCopy, arr)
	sort.Ints(arrCopy)
	for i, j, k := 0, 1, 0; i < len(arrCopy); {
		mapping[arrCopy[i]] = j
		for k = i; k < len(arr)&& arrCopy[k] == arrCopy[i]; k++ {}
		i = k
		j += 1
	}
	for i := 0; i < len(arr); i++ {
		arr[i] = mapping[arr[i]]
	}
	return arr
}
//leetcode submit region end(Prohibit modification and deletion)

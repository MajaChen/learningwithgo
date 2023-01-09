package leetcode

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)

func insertAt(result [][]int, target []int, n int) {
	index := target[1]
	for i := n - 2; i >= index; i-- {
		result[i+1] = result[i]
	}
	result[index] = target
}

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})

	n := len(people)
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		insertAt(result, people[i], n)
	}

	return result
}

//leetcode submit region end(Prohibit modification and deletion)

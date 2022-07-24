package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func distanceBetweenBusStops(distance []int, start int, destination int) int {
	var x, y int
	m := len(distance)
	for i := start; i != destination; {
		x += distance[i]
		i = (i + 1) % m
	}
	for j := start; j != destination; {
		if j == 0 {
			j = m - 1
		} else {
			j -= 1
		}
		y += distance[j]
	}

	if x > y {
		return y
	}
	return x
}

//leetcode submit region end(Prohibit modification and deletion)

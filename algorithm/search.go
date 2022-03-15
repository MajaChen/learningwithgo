package algorithm

func k2IJ(k int) (int, int) {
	seed, sum, row, col := 1, 0, 1, 1
	for {
		if sum += seed; sum >= k {
			break
		} else {
			row++
			seed += 1
		}
	}
	col = row - (sum - k)
	return row, col
}

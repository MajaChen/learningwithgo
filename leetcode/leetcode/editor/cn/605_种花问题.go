package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	l, c := len(flowerbed), 0
	for i := 0; i < l; i++ {
		if flowerbed[i] == 0 {
			if (i == 0 || flowerbed[i-1] == 0) && (i == l-1 || flowerbed[i+1] == 0) {
				flowerbed[i] = 1
				if c++; c >= n {
					return true
				}
			}
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

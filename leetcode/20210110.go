package leetcode

func canFinish(piles []int, speed, h int) bool {

	for i, sum := 0, 0; i < len(piles); i++ {
		if sum += (piles[i] + speed - 1) / speed; sum > h {
			return false
		}
	}

	return true
}

func findMaxIndex(piles []int) int {

	max, maxIndex := piles[0], 0
	for i := 1; i < len(piles); i++ {
		if piles[i] > max {
			max = piles[i]
			maxIndex = i
		}
	}

	return maxIndex
}

func minEatingSpeed(piles []int, h int) int {

	var l, r int
	for l, r = 1, piles[findMaxIndex(piles)]; l < r; {
		if mid := (l + r) / 2; canFinish(piles, mid, h) {
			r = mid // 继续试探一个更小的速度
		} else {
			l = mid + 1
		}
	}

	return l
}

func maxProfit3(prices []int) int {

	sum := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			sum += prices[i] - prices[i-1]
		}
	}
	return sum
}

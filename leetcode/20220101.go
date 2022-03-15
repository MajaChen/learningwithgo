package leetcode

import "math"

var winMapping []int

func sum(n int) int {

	s := 0
	for i := 0; i <= n; i++ {
		s += i
	}
	return s
}

func canIWinRe(maxChoosableInteger, pickedSet, desiredTotal int) bool {

	if winMapping[pickedSet] >= 0 {
		return winMapping[pickedSet] == 1
	}

	for i := 1; i <= maxChoosableInteger; i++ {
		j := 1 << (i - 1)
		if pickedSet&j > 0 { // 未被选择
			if desiredTotal-i <= 0 || !canIWinRe(maxChoosableInteger, pickedSet^j, desiredTotal-i) {
				winMapping[pickedSet] = 1
				return true
			}
		}
	}

	winMapping[pickedSet] = 0
	return false
}

func canIWin(maxChoosableInteger int, desiredTotal int) bool {

	if maxChoosableInteger >= desiredTotal {
		return true
	}

	if sum(maxChoosableInteger) < desiredTotal {
		return false
	}

	m, n := (1<<maxChoosableInteger)-1, desiredTotal
	winMapping = make([]int, 0)
	for i := 0; i <= m; i++ {
		winMapping = append(winMapping, -1)
	}

	return canIWinRe(maxChoosableInteger, m, n)
}

var targetSubWaysCounter int

func findTargetSumWaysRe(nums []int, target, sum, index int) {

	if index >= len(nums) {
		if sum == target {
			targetSubWaysCounter++
		}
		return
	}

	// 对index所指元素加+号
	findTargetSumWaysRe(nums, target, sum+nums[index], index+1)
	// 对index所指元素加-号
	findTargetSumWaysRe(nums, target, sum-nums[index], index+1)
}

func findTargetSumWays(nums []int, target int) int {

	targetSubWaysCounter = 0
	findTargetSumWaysRe(nums, target, 0, 0)
	return targetSubWaysCounter
}

//1.max那加了点号，微小的但却是致命的错误
//2.dp[i+1][j-1],i的遍历方向是向左，j的遍历方向是向右
func longestPalindromeSubseq(s string) int {
	dp := make([][]int, 0, len(s))
	for i := 0; i < len(s); i++ {
		dp = append(dp, make([]int, len(s)))
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			//dp[i][j]
			if i == j {
				dp[i][j] = 1
			} else if s[i:i+1] == s[j:j+1] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				var max int
				if max = dp[i+1][j]; max < dp[i][j-1] {
					max = dp[i][j-1]
				}
				dp[i][j] = max
			}
		}
	}
	return dp[0][len(s)-1]
}

func findBottomLeftValue(root *TreeNode) int {

	res := make([][]int, 0)
	var q Queue = Queue{elems: make([]interface{}, 0)}
	q.Push(root)
	for !q.IsEmpty() {
		size, count, tmpRes := q.Size(), 0, make([]int, 0, q.Size())
		for count < size {
			node := q.Pop().(*TreeNode)
			tmpRes = append(tmpRes, node.Val)
			if node.Left != nil {
				q.Push(node.Left)
			}
			if node.Right != nil {
				q.Push(node.Right)
			}
			count++
		}
		count = 0
		res = append(res, tmpRes)
	}

	return res[len(res)-1][0]
}

func findMin(nums []int) int {

	var l, r int
	for l, r = 0, len(nums)-1; l <= r; {
		if nums[l] <= nums[r] {
			return nums[l]
		}
		if mid := (l + r) / 2; nums[mid] > nums[l] { //nums[mid] > nums[l]
			l = mid + 1
		} else { // nums[mid] <= nums[l]
			r = mid // 死循环的问题->r往左移
		}
	}
	return nums[l]
}

var changeCount int

func changeRe(amount, index int, coins []int) {

	if amount == 0 {
		changeCount++
		return
	}

	for i := index; i < len(coins); i++ {
		if amount-coins[i] >= 0 {
			continue
		}
		changeRe(amount-coins[i], i, coins)
	}
}

func change2(amount int, coins []int) int {

	changeCount = 0
	changeRe(amount, 0, coins)
	return changeCount
}

func change(amount int, coins []int) int {

	m, n := len(coins), amount
	dp := make([][]int, 0, m)
	for i := 0; i < m; i++ {
		arr := make([]int, n+1)
		for i := 0; i <= n; i++ {
			arr = append(arr, -1)
		}
		dp = append(dp, arr)
	}

	for j := 0; j <= n; j++ {
		if j%coins[0] == 0 {
			dp[0][j] = 1
		}
	}

	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}

	for i := 1; i < m; i++ {
		for j := 0; j <= n; j++ {
			dp[i][j] = dp[i-1][j]
			for aj := j; aj-coins[i] >= 0; aj -= coins[i] {
				if dp[i-1][aj-coins[i]] >= 0 {
					dp[i][j] += dp[i-1][aj-coins[i]]
				}
			}
		}
	}
	return dp[m-1][n]
}

func friendCycle(matrix [][]int) int {

	if len(matrix) <= 0 {
		return len(matrix)
	}

	n, set := len(matrix), InitDisjointSet()
	for i := 0; i < n; i++ {
		set.Add(i, i)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 1 {
				set.Union(i, j)
			}
		}
	}

	return set.RootCount()
}

func subarraySum(nums []int, k int) int {

	mapping, sum, count := make(map[int]int), 0, 0
	for i := 0; i < len(nums); i++ {
		if sum += nums[i]; sum == k { // 检查子序列的和
			count++
		}

		if _, ok := mapping[sum-k]; ok { // 检查中间序列的和
			count += mapping[sum-k]
		}

		if _, ok := mapping[sum]; ok {
			i := mapping[sum] + 1
			mapping[sum] = i
		} else {
			mapping[sum] = 1
		}
	}

	return count
}

func maxSubArray(nums []int) int {

	maximalSubArray, minimalSubArray, sum := math.MinInt32, 0, 0
	for i := 0; i < len(nums); i++ {
		if sum += nums[i]; sum-minimalSubArray > maximalSubArray {
			maximalSubArray = sum - minimalSubArray
		}

		if sum < minimalSubArray {
			minimalSubArray = sum
		}
	}

	return maximalSubArray
}

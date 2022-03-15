package leetcode

import (
	"math"
	"strconv"
)

var squareMapping map[int]int

func numSquaresRe(n int) int {

	if n == 0 {
		return 0
	}

	if _, ok := squareMapping[n]; ok {
		return squareMapping[n]
	}

	minimalCount := math.MaxInt32
	for i := int(math.Sqrt(float64(n))); i >= 1; i-- {
		an := n - int(math.Pow(float64(i), float64(2)))
		if count := 1 + numSquaresRe(an); count > 0 && count < minimalCount {
			minimalCount = count
		}
	}

	if minimalCount == math.MaxInt32 {
		minimalCount = -1
	}

	squareMapping[n] = minimalCount
	return minimalCount
}

func numSquares(n int) int {

	squareMapping = make(map[int]int)
	if i := numSquaresRe(n); i == math.MaxInt32 {
		return 0
	} else {
		return i
	}
}

func numSquares2(n int) int {

	squares := make([]int, n+1)
	for i := 1; i <= len(squares); i++ {
		squares[i] = n
	}

	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			if squares[i] < squares[i-j*j]+1 {
				squares[i] = squares[i-j*j] + 1
			}
		}
	}
	return squares[n]
}

var profitMapping []int

func maxProfitRe2(prices []int, startIndex int) int {

	if startIndex >= len(prices) {
		return 0
	}

	if profitMapping[startIndex] >= 0 {
		return profitMapping[startIndex]
	}

	globalMaxProfit := 0
	// 从何时买入
	for i := startIndex; i < len(prices); i++ {
		localMaxProfit := 0
		// 到何时卖出
		for j := i + 1; j < len(prices); j++ {
			var profit int
			if profit = prices[j] - prices[i]; profit <= 0 {
				continue
			}
			if j+2 < len(prices)-1 {
				profit += maxProfitRe2(prices, j+2)
			}
			if profit > localMaxProfit {
				localMaxProfit = profit
			}
		}
		if localMaxProfit > globalMaxProfit {
			globalMaxProfit = localMaxProfit
		}
	}

	profitMapping[startIndex] = globalMaxProfit
	return globalMaxProfit
}

func maxProfit2(prices []int) int {

	profitMapping = make([]int, 0, len(prices))
	for i := 0; i < len(prices); i++ {
		profitMapping = append(profitMapping, -1)
	}

	return maxProfitRe2(prices, 0)
}

func isLetter(s string) bool {
	return s != "[" && s != "]" && !isNumber(s)
}

func isNumber(s string) bool {
	return []rune(s)[0] >= 48 && []rune(s)[0] <= 57
}

func expand(letter string, k int) string {

	str := ""
	for i := 0; i < k; i++ {
		str += letter
	}
	return str
}

func decodeString(s string) string {

	stack := Stack{elems: make([]interface{}, 0)}
	for i := 0; i < len(s); {
		if str := s[i : i+1]; str == "[" {
			stack.Push(str)
			i++
		} else if str == "]" { //遭遇右括号
			target := ""
			for !stack.IsEmpty() && stack.GetTop().(string) != "[" {
				target = stack.Pop().(string) + target
			}
			stack.Pop()            //弹出左括号
			k := stack.Pop().(int) //弹出重复次数
			target = expand(target, k)
			stack.Push(target)
			i++
		} else if isNumber(str) {
			var j int
			for j = i + 1; j < len(s) && isNumber(s[j:j+1]); j++ {
			}
			str = s[i:j]
			k, _ := strconv.Atoi(str)
			stack.Push(k)
			i = j
		} else {
			var j int
			for j = i + 1; j < len(s) && isLetter(s[j:j+1]); j++ {
			}
			stack.Push(s[i:j])
			i = j
		}
	}

	res := ""
	for !stack.IsEmpty() {
		res = stack.Pop().(string) + res
	}

	return res
}

func coinChange(coins []int, amount int) int {

	dps := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		minimalAmount := math.MaxInt32
		for j := 0; j < len(coins); j++ {
			if amount-coins[j] >= 0 && dps[amount-coins[j]] >= 0 {
				if 1+dps[amount-coins[j]] < minimalAmount {
					minimalAmount = 1 + dps[amount-coins[j]]
				}
			}
		}
		if minimalAmount == math.MaxInt32 {
			minimalAmount = -1
		}
		dps[i] = minimalAmount
	}
	return dps[amount]
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	oddListHead := &ListNode{}
	var p, q *ListNode
	for p, q = head, oddListHead; p != nil && p.Next != nil; {
		target := p.Next // 摘
		next := target.Next
		target.Next = nil
		q.Next = target //接
		q = target

		p.Next = next // 承
		if p.Next != nil {
			p = p.Next
		} else {
			break
		}
	}

	p.Next = oddListHead.Next

	return head
}

func increasingTriplet(nums []int) bool {

	minimal, middle := math.MaxInt32, math.MaxInt32
	for i := 0; i < len(nums); i++ {
		if nums[i] < minimal {
			minimal = nums[i]
		} else if nums[i] < middle {
			middle = nums[i]
		} else {
			return true
		}
	}
	return false

}

func intersection(nums1 []int, nums2 []int) []int {

	mapping := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		if _, ok := mapping[nums1[i]]; !ok {
			mapping[nums1[i]] = 0
		}
	}

	intersets := make([]int, 0)
	for i := 0; i < len(nums2); i++ {
		if _, ok := mapping[nums2[i]]; ok {
			delete(mapping, nums2[i])
			intersets = append(intersets, nums2[i])
		}
	}

	return intersets
}

var integerMapping []int

func integerBreakRe(n int) int {

	if integerMapping[n] >= 0 {
		return integerMapping[n]
	}

	// 有些情况不需要分割
	max := n
	for i := 1; i <= n/2; i++ {
		if j := integerBreakRe(i) * integerBreakRe(n-i); j > max {
			max = j
		}
	}

	integerMapping[n] = max
	return max
}

func integerBreak(n int) int {

	if n == 2 || n == 3 {
		return 2
	}

	integerMapping = make([]int, n+1)
	for i := 0; i <= n; i++ {
		integerMapping[i] = -1
	}
	return integerBreakRe(n)
}

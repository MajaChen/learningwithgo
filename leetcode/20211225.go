package leetcode

import (
	"sort"
	"strconv"
)

func nextGreaterElement3(nums []int) []int {

	elements, stack := make([]int, len(nums)), Stack{elems: make([]interface{}, 0)}

	for i := 0; i < len(elements); i++ {
		elements[i] = -1
	}

	for i := 0; i < len(nums); i++ {
		for !stack.IsEmpty() && nums[stack.GetTop().(int)] <= nums[i] {
			elements[stack.GetTop().(int)] = i
			stack.Pop()
		}
		stack.Push(i)
	}
	return elements
}

func nextGreaterElement4(nums []int) []int {

	elements, stack := make([]int, len(nums)), Stack{elems: make([]interface{}, 0)}

	for i := 0; i < len(elements); i++ {
		elements[i] = -1
	}

	for i, n := 0, len(nums); i < n*2-1; i++ {
		for !stack.IsEmpty() && nums[stack.GetTop().(int)] < nums[i%n] {
			elements[stack.GetTop().(int)] = i % n
			stack.Pop()
		}
		stack.Push(i % n)
	}
	return elements
}

func nextGreaterElement5(nums []int) []int {

	elements, stack := make([]int, len(nums)), Stack{elems: make([]interface{}, 0)}

	for i := 0; i < len(elements); i++ {
		elements[i] = -1
	}

	for i := 0; i < len(nums); i++ {
		for !stack.IsEmpty() && nums[stack.GetTop().(int)] < nums[i] {
			elements[stack.GetTop().(int)] = i - stack.GetTop().(int)
			stack.Pop()
		}
		stack.Push(i)
	}
	return elements
}

func findKthLargest(nums []int, k int) int {

	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	return nums[k]
}

func min(i, j, k int) int {
	s := []int{i, j, k}
	sort.Ints(s)
	return s[0]
}

func maximalSquare(matrix [][]byte) int {

	if len(matrix) <= 0 {
		return 0
	}

	m, n, maxSquare := len(matrix), len(matrix[0]), 0
	distance := make([][]int, 0, m)
	for i := 0; i < m; i++ {
		distance = append(distance, make([]int, n))
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				di, dj, dk := 0, 0, 0
				if i-1 >= 0 {
					di = distance[i-1][j]
				}
				if j-1 >= 0 {
					dj = distance[i][j-1]
				}
				if i-1 >= 0 && j-1 >= 0 {
					dk = distance[i-1][j-1]
				}

				if distance[i][j] = 1 + min(di, dj, dk); distance[i][j]*distance[i][j] > maxSquare {
					maxSquare = distance[i][j] * distance[i][j]
				}
			}
		}
	}

	return maxSquare
}

func findKthLargest2(nums []int, k int) int {

	anums := []int{0}
	anums = append(anums, nums...)

	N := len(nums)
	buildHeap(anums, N)
	var target int
	for i := 0; i < k; i++ {
		target = anums[1]
		swap(1, N-i, anums)
		adjust_Big(anums, 1, N-i-1)
	}

	return target
}

func isNum(s string) bool {

	if len(s) != 1 {
		return false
	}

	return []rune(s)[0] >= 48 && []rune(s)[0] <= 57
}

func add(a, b int) int      { return a + b }
func subtract(a, b int) int { return b - a }
func multiply(a, b int) int { return a * b }
func divide(a, b int) int   { return b / a }

func calculate(s string) int {

	priority := map[string]int{
		"+": 3,
		"-": 3,
		"*": 7,
		"/": 7,
		".": 1,
	}

	funcs := map[string]func(a, b int) int{
		"+": add,
		"-": subtract,
		"*": multiply,
		"/": divide,
	}

	as := s + "."
	operStack, numStack := Stack{elems: make([]interface{}, 0)}, Stack{elems: make([]interface{}, 0)}
	operStack.Push(".")
	for i := 0; i < len(as) && !operStack.IsEmpty(); {
		if as[i:i+1] == " " {
			i++
			continue
		}

		if isNum(as[i : i+1]) {
			var j int
			for j = i + 1; j < len(as) && isNum(as[j:j+1]); j++ {
			}
			num, _ := strconv.Atoi(as[i:j])
			numStack.Push(num)
			i = j
		} else {
			if priority[operStack.GetTop().(string)] < priority[as[i:i+1]] {
				// 入栈
				operStack.Push(as[i : i+1])
				i++
			} else {
				// 弹栈加计算
				var oper string
				if oper = operStack.Pop().(string); oper == "." {
					continue
				}

				function := funcs[oper]
				numStack.Push(function(numStack.Pop().(int), numStack.Pop().(int)))
			}
		}
	}

	return numStack.Pop().(int)
}

func majorityElement2(nums []int) []int {

	mapping := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := mapping[nums[i]]; ok {
			count := mapping[nums[i]]
			mapping[nums[i]] = count + 1
		} else {
			mapping[nums[i]] = 1
		}
	}

	elements := make([]int, 0)
	for num, count := range mapping {
		if count > len(nums)/3 {
			elements = append(elements, num)
		}
	}

	return elements
}

var target int

func traverseNodes(root *TreeNode, k int) int {

	if root == nil {
		return k
	}

	if k = traverseNodes(root.Left, k); k == 0 {
		return k
	}

	if k -= 1; k == 0 {
		target = root.Val
		return k
	}

	return traverseNodes(root.Right, k)

}

func kthSmallest2(root *TreeNode, k int) int {

	traverseNodes(root, k)
	return target
}

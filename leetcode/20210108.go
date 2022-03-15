package leetcode

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	// 寻找nums1、nums2长度最小者
	if len(nums1) > len(nums2) {
		num := nums1
		nums1 = nums2
		nums2 = num
	}

	// 寻找i的位置
	m, n := len(nums1), len(nums2)
	// 如果m+n是偶数，则左边元素个数等于右边元素个数，
	// 如果m+n是奇数，则左边元素个数比右边元素个数多一个
	pos := (m + n + 1) / 2
	for l, r, i, j := 0, m, 0, 0; l <= r; {
		mid := (l + r) / 2
		i = mid
		j = pos - i
		if i > 0 && j < n && nums1[i-1] > nums2[j] { // 左移
			r = mid - 1
		} else if j > 0 && i < m && nums1[i] < nums2[j-1] { // 右移
			l = mid + 1
		} else {
			// 寻找左边的最大值
			var leftMax int
			if i == 0 {
				leftMax = nums2[j-1]
			} else if j == 0 {
				leftMax = nums1[i-1]
			} else {
				leftMax = int(math.Max(float64(nums1[i-1]), float64(nums2[j-1])))
			}

			// m+n长度是否为奇数
			if (m+n)%2 != 0 {
				return float64(leftMax)
			}

			// 寻找右边的最小值
			var rightMin int
			if i == m {
				rightMin = nums2[j]
			} else if j == n {
				rightMin = nums1[i]
			} else {
				rightMin = int(math.Min(float64(nums1[i]), float64(nums2[j])))
			}

			return (float64(leftMax) + float64(rightMin)) / 2
		}
	}

	return 0
}

func trap(height []int) int {

	sum := 0
	for left, right, leftMax, rightMax := 0, len(height)-1, 0, 0; left <= right; {
		if height[left] < height[right] {
			if leftMax > height[left] {
				sum += leftMax - height[left]
			} else {
				leftMax = height[left]
			}
			left++
		} else {
			if rightMax > height[right] {
				sum += rightMax - height[right]
			} else {
				rightMax = height[right]
			}
			right--
		}
	}

	return sum
}

func longestValidParentheses(s string) int {

	stack, longestLength := Stack{elems: make([]interface{}, 0)}, 0
	stack.Push(-1)
	for i := 0; i < len(s); i++ {
		if s[i:i+1] == "(" {
			stack.Push(i)
		} else {
			stack.Pop()
			if stack.IsEmpty() {
				stack.Push(i)
			} else {
				if i-stack.GetTop().(int) > longestLength {
					longestLength = i - stack.GetTop().(int)
				}
			}
		}
	}

	return longestLength
}

func reachNumber(target int) int {

	target = int(math.Abs(float64(target)))
	var step int
	for step = 1; target > 0; step++ {
		target -= step
	}

	target = -target
	step -= 1
	if target == 0 || target%2 == 0 {
		return step
	}

	step += 1
	if target += step; target%2 == 0 {
		return step
	}

	return step + 1

}

func findNotGreater(matrix [][]int, target int) int {

	colCount, colIndex, count := len(matrix[0]), len(matrix[0])-1, 0
	for i := 0; i < len(matrix); i++ {
		if matrix[i][colIndex] <= target {
			count += colCount
		} else {
			var j int
			for j = colIndex; j >= 0 && matrix[i][j] > target; j-- {
			}
			count += j + 1
		}
	}
	return count
}

func kthSmallest(matrix [][]int, k int) int {

	var l, r int
	for l, r = matrix[0][0], matrix[len(matrix)-1][len(matrix[0])-1]; l < r; {
		mid := l + ((r - l) >> 2) //这种求middle的方法可以避免middle(-5,-4)=-4情况的发生
		if count := findNotGreater(matrix, mid); count < k {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l
}

//问题抽象：从一个元素集合中，选出n个元素，关注元素排列的顺序([1,2,3] != [1,3,2])，可以重复选
func arrange1(nums []int, n int) [][]int {

	res := make([][]int, 1)
	for i := 0; i < n; i++ {
		var sres [][]int
		for j := 0; j < len(nums); j++ {
			for _, re := range res {
				pre := []int{nums[j]}
				pre = append(pre, re...)
				sres = append(sres, pre)
			}
		}
		res = sres
	}
	return res
}

func arrangeRe2(nums []int, n int, visited []int) [][]int {
	if n == 0 {
		return make([][]int, 1)
	}

	globalRes := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if visited[i] > 0 {
			continue
		}
		visited[i] = 1
		localRes := arrangeRe2(nums, n-1, visited)
		visited[i] = -1
		for index, re := range localRes {
			re = append(re, nums[i])
			localRes[index] = re
		}
		globalRes = append(globalRes, localRes...)
	}
	return globalRes
}

// 从一个元素集合中，选出n个元素，关注元素排列的顺序([1,2,3] != [1,3,2])，不可以重复选
func arrange2(nums []int, n int) [][]int {

	visited := make([]int, len(nums))
	for i := 0; i < len(visited); i++ {
		visited[i] = -1
	}

	return arrangeRe2(nums, n, visited)
}

func arrangeRe3(nums []int, n, startIndex int) [][]int {

	if n == 1 {
		res := make([][]int, 0, len(nums))
		for i := startIndex; i < len(nums); i++ {
			arr := []int{nums[i]}
			res = append(res, arr)
		}
		return res
	}

	globalRes := make([][]int, 0)
	for i := startIndex; i < len(nums) && i+n <= len(nums); i++ {
		localRes := arrangeRe3(nums, n-1, i+1)
		for index, re := range localRes {
			re = append(re, nums[i])
			localRes[index] = re
		}
		globalRes = append(globalRes, localRes...)
	}
	return globalRes
}

// 从一个元素集合中，选出n个元素，不关注元素排列的顺序([1,2,3] == [1,3,2])，不可以重复选
func arrange3(nums []int, n int) [][]int {

	return arrangeRe3(nums, n, 0)
}

func permuteUniqueRe(nums []int, n int, visited []int) [][]int {
	if n == 0 {
		return make([][]int, 1)
	}

	globalRes := make([][]int, 0)
	for i := 0; i < len(nums); {
		if visited[i] > 0 {
			i++
			continue
		}
		visited[i] = 1
		localRes := permuteUniqueRe(nums, n-1, visited)
		visited[i] = -1
		for index, re := range localRes {
			re = append(re, nums[i])
			localRes[index] = re
		}
		globalRes = append(globalRes, localRes...)

		for i += 1; i < len(nums) && nums[i] == nums[i-1]; i++ {
		}
	}
	return globalRes
}

func permuteUnique(nums []int) [][]int {

	visited := make([]int, len(nums))
	for i := 0; i < len(visited); i++ {
		visited[i] = -1
	}
	return permuteUniqueRe(nums, len(nums), visited)
}

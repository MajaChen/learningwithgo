package leetcode

//leetcode submit region begin(Prohibit modification and deletion)

import "learning/common"

func largestRectangleAreaV2(heights []byte) int {

	stack, largestArea := common.Stack{Elems: make([]interface{}, 0)}, 0
	heights = append(heights, 0)
	for i := 0; i < len(heights); i++ {
		for !stack.IsEmpty() && heights[stack.GetTop().(int)] > heights[i] {
			var l, r, h int
			h = int(heights[stack.Pop().(int)])
			if stack.IsEmpty() {
				l = -1
			} else {
				l = stack.GetTop().(int)
			}
			r = i
			if area := (r - l - 1) * h; area > largestArea {
				largestArea = area
			}
		}
		stack.Push(i)
	}
	return largestArea
}

func maximalRectangle(matrix [][]byte) int {

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[i][j] -= 48
		}
	}

	heights, largestRetangle := matrix[0], 0
	largestRetangle = largestRectangleAreaV2(heights)

	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				heights[j] = 0
			} else {
				heights[j] += 1
			}
		}
		if area := largestRectangleAreaV2(heights); area > largestRetangle {
			largestRetangle = area
		}
	}

	return largestRetangle
}

//leetcode submit region end(Prohibit modification and deletion)

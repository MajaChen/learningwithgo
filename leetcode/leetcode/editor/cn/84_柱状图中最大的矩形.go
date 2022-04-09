package leetcode

//leetcode submit region begin(Prohibit modification and deletion)

func largestRectangleArea(heights []int) int {

	stack, largestArea := Stack{elems: make([]interface{}, 0)}, 0
	heights = append(heights, 0)
	for i := 0; i < len(heights); i++ {
		for !stack.IsEmpty() && heights[stack.GetTop().(int)] > heights[i] {
			var l, r, h int
			h = heights[stack.Pop().(int)]
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

//leetcode submit region end(Prohibit modification and deletion)

package leetcode

//leetcode submit region begin(Prohibit modification and deletion)

type Stack struct {
	elems []interface{}
}

func (s *Stack) IsEmpty() bool {
	return len(s.elems) == 0
}

func (s *Stack) Len() int {
	return len(s.elems)
}

func (s *Stack) Pop() interface{} {
	elem := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]
	return elem
}

func (s *Stack) Push(elem interface{}) {
	s.elems = append(s.elems, elem)
}

func (s *Stack) GetTop() interface{} {
	return s.elems[len(s.elems)-1]
}

func largestRectangleAreaV2(heights []byte) int {

	stack, largestArea := Stack{elems: make([]interface{}, 0)}, 0
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

	heights, ans := matrix[0], 0
	ans = largestRectangleAreaV2(heights)

	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				heights[j] = 0
			} else {
				heights[j] += 1
			}
		}
		if area := largestRectangleAreaV2(heights); area > ans {
			ans = area
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

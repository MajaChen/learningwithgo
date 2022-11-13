package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func countStudents(students []int, sandwiches []int) int {
	s0, s1 := 0, 0
	for _, student := range students {
		if student == 0 {
			s0++
		} else if student == 1 {
			s1++
		}
	}
	for _, sandwiche := range sandwiches {
		if sandwiche == 0 {
			if s0 > 0 {
				s0--
			} else {
				break
			}
		}
		if sandwiche == 1 {
			if s1 > 0 {
				s1--
			} else {
				break
			}
		}
	}

	return s0+s1

}
//leetcode submit region end(Prohibit modification and deletion)

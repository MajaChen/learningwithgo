package leetcode

import "strings"

//leetcode submit region begin(Prohibit modification and deletion)
func reformatNumber(number string) string {
	nums := make([]rune, 0)
	for _, n := range number {
		if n >= 48 && n <= 57 {
			nums = append(nums, n)
		}
	}
	s := strings.Builder{}
	var i int
	for i = 0; i < len(nums)-4; i += 3 {
		for j := i; j < i+3; j++ {
			s.WriteRune(nums[j])
		}
		s.WriteString("-")
	}

	nums = nums[i:]
	if len(nums) == 4 {
		for i = 0; i < len(nums); i += 2 {
			for j := i; j < i+2; j++ {
				s.WriteRune(nums[j])
			}
			s.WriteString("-")
		}
	} else {
		for j := 0; j < len(nums); j++ {
			s.WriteRune(nums[j])
		}
	}

	ans := s.String()
	if strings.HasSuffix(ans, "-") {
		ans = ans[:len(ans)-1]
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

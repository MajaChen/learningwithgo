package leetcode

func plusOne(digits []int) []int {
	flag := 1
	for i := len(digits) - 1; i >= 0 && flag != 0; i++ {
		var sum int
		if sum = digits[i] + flag; sum >= 10 {
			sum -= 10
			flag = 1
		} else {
			flag = 0
		}
		digits[i] = sum
	}
	if flag == 1 {
		newDigits := make([]int, 0, len(digits)+1)
		newDigits = append(newDigits, 1)
		newDigits = append(newDigits, digits...)
		digits = newDigits
	}
	return digits
}

func lengthOfLongestSubstring(s string) int {

	if len(s) <= 1 {
		return len(s)
	}

	mapping, l, r, maxLen := make(map[string]int), 0, 0, 1
	mapping[s[0:1]] = 0
	for i := 1; i < len(s); i++ {
		r = i
		if _, ok := mapping[s[i:i+1]]; ok {
			var index int
			for index = l; index < mapping[s[i:i+1]]; index++ {
				delete(mapping, s[index:index+1])
			}
			l = mapping[s[i:i+1]] + 1
		}
		mapping[s[i:i+1]] = i

		if r-l+1 > maxLen {
			maxLen = r - l + 1
		}
	}

	return maxLen
}

func longestPalindrome(s string) string {

	if len(s) <= 1 {
		return s
	}

	n, maxStartPoint, maxLen := len(s), 0, 1
	mapping := make([][]bool, 0, n)
	for i := 0; i < n; i++ {
		mapping = append(mapping, make([]bool, n))
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				mapping[i][j] = true
			}
		}
	}

	for span := 0; span < n; span++ {
		flag, startPoint := false, 0
		for i := 0; i+span < n; i++ {
			if l, r := i, i+span; s[l:l+1] == s[r:r+1] && (r-l <= 1 || mapping[l+1][r-1] == true) {
				mapping[l][r] = true
				flag, startPoint = true, i
			}
		}
		if flag {
			maxLen = span
			maxStartPoint = startPoint
		}
	}
	return s[maxStartPoint : maxStartPoint+maxLen+1]
}

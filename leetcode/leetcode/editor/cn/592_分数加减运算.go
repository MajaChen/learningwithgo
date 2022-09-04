package leetcode

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)

func extract(expression string) ([]int, []int, []rune) {
	numerators, denominators, operators := make([]int, 0), make([]int, 0), make([]rune, 0)
	if !strings.HasPrefix(expression, "-") {
		expression = fmt.Sprintf("%s%s", "+", expression)
	}
	isDenominator := true
	for i := 0; i < len(expression); {
		if expression[i] == 43 || expression[i] == 45 {
			operators = append(operators, rune(expression[i]))
			i++
		} else if expression[i] >= 48 && expression[i] <= 57 {
			var j int
			for j = i; j < len(expression) && expression[j] >= 48 && expression[j] <= 57; j++ {
			}
			num, _ := strconv.Atoi(expression[i:j])
			if isDenominator {
				denominators = append(denominators, num)
				isDenominator = false
			} else {
				numerators = append(numerators, num)
				isDenominator = true
			}
			i = j
		} else {
			i++
		}
	}

	return numerators, denominators, operators
}

func findMinCommonMultipleOfTwo(i, j int) int {
	d := findMaxCommonDivisor(i, j)
	return (i / d) * j
}

func findMinCommonMultiple(nums []int) int {
	x, y := nums[0], nums[1]
	for i := 1; i < len(nums); i++ {
		y = nums[i]
		x = findMinCommonMultipleOfTwo(x, y)
	}
	return x
}

func eucid(i, j int) int {
	if j == 0 {
		return i
	}
	return eucid(j, i%j)

}

func findMaxCommonDivisor(i, j int) int {
	return eucid(i, j)
}

func fractionAddition(expression string) string {
	denominators, numerators, operators := extract(expression)
	if len(denominators) <= 1 {
		return expression
	}
	m := findMinCommonMultiple(denominators)
	s := 0
	for i, numerator := range numerators {
		numerator *= m / denominators[i]
		if operators[i] == 45 {
			numerator = -numerator
		}
		s += numerator
	}
	var isNegtive bool
	if s < 0 {
		s = int(math.Abs(float64(s)))
		isNegtive = true
	}
	d := findMaxCommonDivisor(m, s)
	m /= d
	s /= d
	ans := fmt.Sprintf("%v/%v", s, m)
	if isNegtive {
		ans = fmt.Sprintf("%s%s", "-", ans)
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

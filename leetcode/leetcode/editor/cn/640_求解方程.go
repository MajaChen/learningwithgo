package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)

func parseExpression(expression string) (int, int) {
	if !strings.HasPrefix(expression, "+") || !strings.HasPrefix(expression, "-") {
		expression = "+" + expression
	}
	counter, sum := 0, 0
	for i := 0; i < len(expression); {
		if expression[i] == 43 || expression[i] == 45 {
			if i+1 < len(expression) && expression[i+1] == 120 {
				if expression[i] == 43 {
					counter += 1
				} else {
					counter -= 1
				}
				i += 2
			} else {
				var j int
				for j = i + 1; j < len(expression) && (expression[j] >= 48 && expression[j] <= 57); j++ {
				}
				num, _ := strconv.Atoi(expression[i+1 : j])
				if j < len(expression) && expression[j] == 120 {
					if expression[i] == 43 {
						counter += num
					} else {
						counter -= num
					}
					i = j + 1
				} else {
					if expression[i] == 43 {
						sum += num
					} else {
						sum -= num
					}
					i = j
				}
			}
		}
	}
	return counter, sum
}

func solveEquation(equation string) string {
	ss := strings.Split(equation, "=")
	i, m := parseExpression(ss[0])
	j, n := parseExpression(ss[1])
	if i == j {
		if m-n == 0 {
			return "Infinite solutions"
		} else {
			return "No solution"
		}
	}
	ans := (n - m) / (i - j)
	return fmt.Sprintf("x=%v", ans)
}

//leetcode submit region end(Prohibit modification and deletion)

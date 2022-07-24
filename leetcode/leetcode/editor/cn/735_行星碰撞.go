package leetcode

import "learning/common"

//leetcode submit region begin(Prohibit modification and deletion)

func asteroidCollision(asteroids []int) []int {
	alive, stack, n :=
		make([]bool, len(asteroids)),
		common.Stack{Elems: make([]interface{}, 0, len(asteroids))},
		len(asteroids)

	for i := 0; i < n; i++ {
		alive[i] = true
	}

	for i, a := range asteroids {
		if a > 0 {
			stack.Push(i)
		} else {
			stackLoop:
			for !stack.IsEmpty() {
				j := stack.GetTop().(int)
				switch {
				case asteroids[i]+asteroids[j] == 0:
					stack.Pop()
					alive[i] = false
					alive[j] = false
					break stackLoop
				case asteroids[i]+asteroids[j] > 0:
					alive[i] = false
					break stackLoop
				default:
					stack.Pop()
					alive[j] = false
				}
			}
		}
	}

	ans := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if alive[i] {
			ans = append(ans, asteroids[i])
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

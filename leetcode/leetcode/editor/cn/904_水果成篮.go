package leetcode

import "math"

//leetcode submit region begin(Prohibit modification and deletion)

func deleteFromFruitsMapping(fruitsMapping map[int]int) int {
	mk, mv := 0, math.MaxInt32
	for k, v := range fruitsMapping {
		if v < mv {
			mv = v
			mk = k
		}
	}
	delete(fruitsMapping, mk)
	return mv
}

func totalFruit(fruits []int) int {
	ans, l, r, fruitsMapping := 1, 0, 0, make(map[int]int)
	fruitsMapping[fruits[0]] = 0
	for i := 1; i < len(fruits); i++ {
		if len(fruitsMapping) < 2 {
			fruitsMapping[fruits[i]] = i
			r = i
		} else {
			if _, ok := fruitsMapping[fruits[i]]; ok {
				fruitsMapping[fruits[i]] = i
				r = i
			} else {
				j := deleteFromFruitsMapping(fruitsMapping)
				l = j+1
				r = i
				fruitsMapping[fruits[i]] = i
			}
		}
		if r-l+1 > ans {
			ans = r-l+1
		}
	}
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)

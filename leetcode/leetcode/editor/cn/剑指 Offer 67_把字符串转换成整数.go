package leetcode

import "math"

//leetcode submit region begin(Prohibit modification and deletion)

const START = 0
const SIGNED = 1
const NUMBER = 2
const END = 3

var ans int64
var currentState int
var stateMapping [][]int
var isPositive bool

func col2State(r rune) int {
	if r == 32 {
		return START
	} else if r == 43 || r == 45 {
		return SIGNED
	} else if r >= 48 && r <= 57 {
		return NUMBER
	} else {
		return END
	}
}

func getState(r rune) {
	if currentState = stateMapping[currentState][col2State(r)]; currentState == NUMBER {
		if isPositive {
			ans = ans*10 + (int64(r) - 48)
			ans = int64(math.Min(float64(ans), float64(math.MaxInt32)))
		} else {
			ans = ans*10 - (int64(r) - 48)
			ans = int64(math.Max(float64(ans), float64(math.MinInt32)))
		}
	} else if currentState == SIGNED {
		if r == 45 {
			isPositive = false
		}
	}
}

func strToInt(str string) int {

	ans = 0
	isPositive = true
	currentState = START
	stateMapping = make([][]int, 0, 4)
	stateMapping = append(stateMapping, []int{START, SIGNED, NUMBER, END})
	stateMapping = append(stateMapping, []int{END, END, NUMBER, END})
	stateMapping = append(stateMapping, []int{END, END, NUMBER, END})
	stateMapping = append(stateMapping, []int{END, END, END, END})

	runes := []rune(str)
	for _, r := range runes {
		getState(r)
	}

	return int(ans)
}

//leetcode submit region end(Prohibit modification and deletion)

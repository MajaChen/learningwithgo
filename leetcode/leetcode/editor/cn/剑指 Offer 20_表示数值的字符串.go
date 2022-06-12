package leetcode

//leetcode submit region begin(Prohibit modification and deletion)

const (
	BEGIN   = 0
	SIGN    = 1
	NUM     = 2
	POINT   = 3
	SCIENCE = 4
	STOP    = 5
	ABORT   = 6
)

var (
	states       [][]int
	curState     int
	pointExist   bool
	scienceExist bool
	numberExist  bool
)

func rune2State(r rune) int {
	if r == 32 {
		return BEGIN
	} else if r == 43 || r == 45 {
		return SIGN
	} else if r >= 48 && r <= 57 {
		return NUM
	} else if r == 46 {
		if pointExist || scienceExist {
			return ABORT
		}
		return POINT
	} else if r == 69 || r == 101 {
		if scienceExist || !numberExist {
			return ABORT
		}
		return SCIENCE
	} else {
		return ABORT
	}

}

func isNumber(s string) bool {

	states = make([][]int, 0, 7)
	states = append(states, []int{BEGIN, SIGN, NUM, POINT, ABORT, ABORT, ABORT})
	states = append(states, []int{ABORT, ABORT, NUM, POINT, ABORT, ABORT, ABORT})
	states = append(states, []int{STOP, ABORT, NUM, POINT, SCIENCE, ABORT, ABORT})
	states = append(states, []int{STOP, ABORT, NUM, ABORT, SCIENCE, ABORT, ABORT})
	states = append(states, []int{ABORT, SIGN, NUM, ABORT, ABORT, ABORT, ABORT})
	states = append(states, []int{STOP, ABORT, ABORT, ABORT, ABORT, ABORT, ABORT})
	states = append(states, []int{ABORT, ABORT, ABORT, ABORT, ABORT, ABORT, ABORT})
	curState = BEGIN
	pointExist = false
	scienceExist = false
	numberExist = false

	runes := []rune(s)
	runes = append(runes, 32)
	for _, r := range runes {
		if curState = states[curState][rune2State(r)]; curState == NUM {
			numberExist = true
		} else if curState == POINT {
			pointExist = true
		} else if curState == SCIENCE {
			scienceExist = true
		}
	}

	return curState == STOP && numberExist
}

//leetcode submit region end(Prohibit modification and deletion)

package leetcode

import (
	"learning/common"
	"strconv"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)

type stackFrame struct {
	fun   int
	typ   string
	time  int
	extra int
}

func exclusiveTime(n int, logs []string) []int {
	frames := make([]*stackFrame, 0, len(logs))
	for i := 0; i < len(logs); i++ {
		ss := strings.Split(logs[i], ":")
		f, _ := strconv.Atoi(ss[0])
		t, _ := strconv.Atoi(ss[2])
		frames = append(frames, &stackFrame{fun: f, typ: ss[1], time: t})
	}

	stack, funcs := common.Stack{Elems: make([]interface{}, 0)}, make([]int, n)
	for _, frame := range frames {
		if frame.typ == "start" {
			stack.Push(frame)
		} else {
			f := stack.Pop().(*stackFrame)
			timeLen := frame.time - f.time + 1
			if !stack.IsEmpty() {
				f := stack.GetTop().(*stackFrame)
				f.extra += timeLen
			}
			timeLen -= f.extra
			funcs[f.fun] += timeLen
		}
	}
	return funcs
}

//leetcode submit region end(Prohibit modification and deletion)

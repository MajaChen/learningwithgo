package leetcode

import (
	"math"
)

//leetcode submit region begin(Prohibit modification and deletion)

func findParent(label int) int {
	layer := int(math.Floor(math.Log2(float64(label)))) + 1

	curBase := int(math.Pow(2, float64(layer-1)))
	parentBase := int(math.Pow(2, float64(layer-2)))
	parentNum := parentBase
	return parentBase + (parentNum - ((label-curBase)/2 + 1))

}

func pathInZigZagTree(label int) []int {

	res := []int{label}
	for label != 1 {
		parent := findParent(label)
		res = append(res, parent)
		label = parent
	}

	for i := 0; i < len(res)/2; i++ {
		t := res[i]
		res[i] = res[len(res)-i-1]
		res[len(res)-i-1] = t
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)

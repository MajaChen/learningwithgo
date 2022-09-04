package leetcode

import "math"

//leetcode submit region begin(Prohibit modification and deletion)

func distance(x, y []int) float64 {
	return math.Sqrt(math.Pow(math.Abs(float64(y[0]-x[0])), 2) +
		math.Pow(math.Abs(float64(y[1]-x[1])), 2))
}

func equals(p1, p2 []int) bool {
	return p1[0] == p2[0] && p1[1] == p2[1]
}

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	if equals(p1, p2) ||
		equals(p1, p3) ||
		equals(p1, p4) ||
		equals(p2, p3) ||
		equals(p2, p4) ||
		equals(p3, p4) {
		return false
	}
	a, b, c := distance(p1, p2), distance(p1, p3), distance(p1, p4)
	var A, B, C, D []int
	var x, z float64
	if a == b && a != c {
		A, B, C, D = p1, p2, p3, p4
		x, z = a, c
	} else if a == c && a != b {
		A, B, C, D = p1, p2, p4, p3
		x, z = a, b
	} else if b == c && b != a {
		A, B, C, D = p1, p3, p4, p2
		x, z = b, a
	}

	if len(A) == 0 {
		return false
	}

	a, b, c = distance(D, B), distance(D, C), distance(B, C)
	return a == b && a == x && c == z
}

//leetcode submit region end(Prohibit modification and deletion)

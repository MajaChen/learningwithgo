package algorithm

import "testing"

func TestOptimalBST(t *testing.T) {

	p := []float64{0, 0.15, 0.1, 0.05, 0.10, 0.20}
	q := []float64{0.05, 0.1, 0.05, 0.05, 0.05, 0.1}
	const n = 5
	cost, roots := optimalBST(p, q, 5)
	t.Errorf("cost is %v,roots is %v", cost, roots)
}

func TestDoubleArr(t *testing.T) {
	const n = 2
	dArr := [n][n]int{}
	t.Errorf("dArr is %v", dArr)
}

func TestLCSLength(t *testing.T) {

	A := []string{"X", "A", "B", "C", "B", "D", "A", "B"}
	B := []string{"X", "B", "D", "C", "A", "B", "A"}
	//  A := []string{"X","A","B"}
	//  B := []string{"X","A","B"}
	l, p := LCSLength(A, B)
	t.Errorf("l is %v,p is %v", l, p)
}

func TestMatrixChainOrder(t *testing.T) {
	matrixes := [][2]int{
		{30, 35},
		{35, 15},
		{15, 5},
		{5, 10},
		{10, 20},
		{20, 25},
	}
	m, s := matrixChainOrder_v2(matrixes)
	t.Errorf("m is %v,s is %v", m, s)
}

func TestFatestWay(test *testing.T) {
	e1, e2, n1, n2 := 2, 4, 3, 2
	a := [2][]int{
		{7, 9, 3, 4, 8, 4},
		{8, 5, 6, 4, 5, 7},
	}
	t := [2][]int{
		{2, 3, 1, 3, 4},
		{2, 1, 2, 2, 1},
	}
	minTime, path := fatestWay(e1, e2, n1, n2, 6, a, t)
	test.Errorf("minTime is %v,path is %v", minTime, path)
}

func TestCuttingRope(t *testing.T) {
	n := 8
	t.Errorf("res is %v", cuttingRope(n))
}

func TestRob(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	t.Errorf("res is %v", rob(nums))
}

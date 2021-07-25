package algorithm

import (
	"math"
	"strconv"
)

//装配线调度问题非递归解
func fatestWay(e1, e2, n1, n2, n int, a, t [2][]int) (int, []int) {
	// 二维数组，第一维表示装配线，分为装配线1和装配线2
	// 第二维表示装配站的个数
	f := [2][]int{
		make([]int, n),
		make([]int, n),
	}

	l := [2][]int{
		make([]int, n),
		make([]int, n),
	}

	f[0][0] = e1 + a[0][0]
	f[1][0] = e2 + a[1][0]

	for i := 1; i < n; i++ {
		if f[1][i-1]+a[0][i]+t[1][i-1] < f[0][i-1]+a[0][i] {
			f[0][i] = f[1][i-1] + a[0][i] + t[1][i-1]
			l[0][i] = 1
		} else {
			f[0][i] = f[0][i-1] + a[0][i]
			l[0][i] = 0
		}
		if f[0][i-1]+a[1][i]+t[0][i-1] < f[1][i-1]+a[1][i] {
			f[1][i] = f[0][i-1] + a[1][i] + t[0][i-1]
			l[1][i] = 0
		} else {
			f[1][i] = f[1][i-1] + a[1][i]
			l[1][i] = 1
		}
	}

	var minTime int = f[0][n-1] + n1
	path := make([]int, n)
	path[n-1] = 0
	if minTime > f[1][n-1]+n2 {
		minTime = f[1][n-1] + n2
		path[n-1] = 1
	}

	for i := n - 1; i >= 1; i-- {
		path[i-1] = l[path[i]][i]
	}
	return minTime, path
}

func findSplit(s [][]int, leftBound, rightBound int) []int {
	splits := make([]int, 0)
	if leftBound == rightBound {
		return splits
	}
	split := s[leftBound][rightBound]
	splits = append(splits, split)
	splits = append(splits, findSplit(s, leftBound, split)...)
	splits = append(splits, findSplit(s, split+1, rightBound)...)
	return splits
}

// 矩阵链乘法问题非递归解
func matrixChainOrder_v2(matrixes [][2]int) (int, []int) {

	var l int
	if l = len(matrixes); l <= 1 {
		return 0, nil
	}

	// 二维数组m表示一个子链的计算代价
	// m的第一维表示子链的左边界
	// m的第二维表示子链的右边界
	m, s := make([][]int, 0, l), make([][]int, 0, l)
	for i := 0; i < l; i++ {
		m = append(m, make([]int, l))
		s = append(s, make([]int, l))
	}

	for span := 1; span <= l; span++ {
		for i := 0; i+span < l; i++ {
			l, r := i, i+span
			m[l][r] = math.MaxInt32
			for j := l; j < r; j++ {
				if m[l][j]+m[j+1][r]+(matrixes[l][0]*matrixes[j][1]*matrixes[r][1]) < m[l][r] {
					m[l][r] = m[l][j] + m[j+1][r] + (matrixes[l][0] * matrixes[j][1] * matrixes[r][1])
					s[l][r] = j
				}
			}
		}
	}

	lbound, rbound := 0, l-1
	return m[0][l-1], findSplit(s, lbound, rbound)
}

func LCSLength(A, B []string) (int, []string) {

	if len(A) == 0 || len(B) == 0 {
		return 0, nil
	}

	// 二维数组length的第一维表示数组A的下标
	// 第二维表示数组B的下标
	// 三维数组lcs的第一维表示数组A的下标，
	// 第二维表示数组B的下标
	// 第三维表示lcs序列
	al, bl := len(A), len(B)
	length, lcs := make([][]int, 0, al), make([][][]string, 0, al)

	for i := 0; i < al; i++ {
		length = append(length, make([]int, bl))
	}

	for i := 1; i < al; i++ {
		length[i][0] = 0
	}
	for j := 1; j < bl; j++ {
		length[0][j] = 0
	}

	for i := 0; i < al; i++ {
		arr := make([][]string, 0, bl)
		for j := 0; j < bl; j++ {
			arr = append(arr, make([]string, 0))
		}
		lcs = append(lcs, arr)
	}

	for i := 1; i < al; i++ {
		for j := 1; j < bl; j++ {
			if A[i] == B[j] {
				length[i][j] = length[i-1][j-1] + 1
				lcs[i][j] = append(lcs[i-1][j-1], A[i])
			} else {
				lcs[i][j] = append(lcs[i][j], lcs[i][j-1]...)
				if length[i][j] = length[i][j-1]; length[i][j] < length[i-1][j] {
					length[i][j] = length[i-1][j]
					lcs[i][j] = append(lcs[i][j], lcs[i-1][j]...)
				}
			}
		}
	}

	return length[al-1][bl-1], lcs[al-1][bl-1]
}

func weight(p, q []float64, l, r int) float64 {
	var w float64 = 0
	for i := l; i <= r; i++ {
		w += p[i]
	}
	for i := l - 1; i <= r; i++ {
		w += q[i]
	}
	return w
}

func findRoot(l, r int, roots [][]int) []string {

	res := make([]string, 0)
	if l > r {
		res = append(res, "nil")
		return res
	}
	res = append(res, strconv.Itoa(roots[l][r]))
	res = append(res, findRoot(l, roots[l][r]-1, roots)...)
	res = append(res, findRoot(roots[l][r]+1, r, roots)...)
	return res
}

func optimalBST(p, q []float64, n int) (float64, []string) {
	if (len(p) != n && n != len(q)-1) || n == 0 {
		return 0, nil
	}
	// 二维数组cost表示子数组所代表的BST的代价
	// 第一维表示子数组的左边界
	// 第二维表示子数组的右边界
	costs, roots, span := make([][]float64, 0, n), make([][]int, 0, n), 0
	//初始化cost
	for i := 0; i < n+2; i++ {
		cost := make([]float64, n+2)
		for j := 0; j < n+2; j++ {
			if j == i-1 {
				cost[j] = q[j]
			}
		}
		costs = append(costs, cost)

		root := make([]int, n+2)
		for j := 1; j < n+2; j++ {
			if j == i {
				root[j] = j
			}
		}
		roots = append(roots, root)
	}

	// 计算cost
	for span = 0; span < n; span++ {
		for i := 1; i <= n-span; i++ {
			l, r := i, i+span
			costs[l][r] = math.MaxFloat64
			roots[l][r] = l
			for j := l; j <= r; j++ {
				if costs[l][j-1]+costs[j+1][r]+weight(p, q, l, r) < costs[l][r] {
					costs[l][r] = costs[l][j-1] + costs[j+1][r] + weight(p, q, l, r)
					roots[l][r] = j
				}
			}
		}
	}
	//输出结果
	return costs[1][n], findRoot(1, n, roots)
}

// 打家劫舍问题
func rob(nums []int) int {

	if len(nums) <= 0{
		return 0
	}

	l, max := len(nums), math.MinInt32
	res := make([]int, l)
	for i := 0; i < l; i++ {
		lmax := 0
		for j := 0; j <= i-2; j++ {
			if res[j] > lmax {
				lmax = res[j]
			}
		}
		res[i] = lmax + nums[i]
		if res[i] > max {
			max = res[i]
		}
	}
	return max
}

func cuttingRope(n int) int {
	if n <= 3 {
		if n <= 0 {
			return 0
		}
		return n - 1
	}

	// 二维数组res表示对绳子的某一段做“剪绳子”操作得到的最优结果
	// 第一维表示子段的左边界
	// 第二维表示子段的右边界
	res := make([][]int, 0, n)
	for i := 0; i < n; i++ {
		res = append(res, make([]int, n))
	}

	for i := 0; i < n; i++ {
		res[i][i] = 1
	}

	for span := 1; span < n; span++ {
		for i := 0; i+span < n; i++ {
			l, r := i, i+span
			res[l][r] = math.MinInt32
			for j := l; j < r; j++ {
				if res[l][j]*res[j+1][r] > res[l][r] {
					res[l][r] = res[l][j] * res[j+1][r]
				}
			}
			if r-l+1 > res[l][r] {
				res[l][r] = r - l + 1
			}
		}
	}
	return res[0][n-1]
}

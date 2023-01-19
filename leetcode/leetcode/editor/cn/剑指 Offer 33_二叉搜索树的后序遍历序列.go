package leetcode

/*
class Solution:
    def verifyPostorder(self, postorder: [int]) -> bool:
        def recur(i, j):
            if i >= j: return True
            p = i
            while postorder[p] < postorder[j]: p += 1
            m = p
            while postorder[p] > postorder[j]: p += 1
            return p == j and recur(i, m - 1) and recur(m, j - 1)

        return recur(0, len(postorder) - 1)

*/

//leetcode submit region begin(Prohibit modification and deletion)
func verifyPostorderRe(postorder []int, i, j int) bool {
	if i >= j {
		return true
	}
	p := i
	for ; postorder[p] < postorder[j]; p++ {
	}
	// 左半区间
	m := p
	for ; postorder[p] > postorder[j]; p++ {
	}
	// 右半区间
	return p == j && verifyPostorderRe(postorder, i, m-1) && verifyPostorderRe(postorder, m, j-1)
}
func verifyPostorder(postorder []int) bool {
	return verifyPostorderRe(postorder, 0, len(postorder)-1)
}

//leetcode submit region end(Prohibit modification and deletion)

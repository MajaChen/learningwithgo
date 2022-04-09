package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {

	arr := make([]int, 0)
	for p := head; p != nil; p = p.Next {
		arr = append(arr, p.Val)
	}

	for i, l := 0, len(arr); i < l/2; i++ {
		t := arr[i]
		arr[i] = arr[l-i-1]
		arr[l-i-1] = t
	}

	return arr
}

//leetcode submit region end(Prohibit modification and deletion)

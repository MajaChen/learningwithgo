package leetcode

import "math"

var res []string

//1)修改pre，使其next指针指向正确的结点
//2)记录head的下一个结点是谁，并“互换”head与head.next
//3)重开，参数分别是(head,head.next)
func generateParenthesisRe(l, r int, str string) {
	if l > r || l < 0 || r < 0 {
		return
	}
	if l == 0 && r == 0 {
		res = append(res, str)
		return
	}
	// 做出选择"("或者")"
	generateParenthesisRe(l-1, r, str+"(")
	generateParenthesisRe(l, r-1, str+")")
	return
}

func generateParenthesis(n int) []string {
	l, r, str := n, n, ""
	res = make([]string, 0)
	generateParenthesisRe(l, r, str)
	return res
}

func swapPairsRe(pre, head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	headNext := head.Next
	pre.Next = head.Next
	head.Next = head.Next.Next
	headNext.Next = head
	pre = pre.Next.Next
	swapPairsRe(pre, head)
}

func swapPairs(head *ListNode) *ListNode {
	pre := &ListNode{Next: head}
	swapPairsRe(pre, head)
	return pre.Next
}

func letterOrNum(s string) bool {
	if len(s) != 1 {
		return false
	}
	r := []rune(s)[0]
	return number(r) || letter(r)
}

func number(r rune) bool {
	if r >= 48 && r <= 57 {
		return true
	} else {
		return false
	}
}

func letter(r rune) bool {
	if (r >= 65 && r <= 90) || (r >= 97 && r <= 122) {
		return true
	} else {
		return false
	}
}

func upperOrLower(a, b string) bool {
	if len(a) != len(b) || len(a) != 1 {
		return false
	}
	ai, bi := []rune(a)[0], []rune(b)[0]
	if ai == bi || (letter(ai) && letter(bi) && math.Abs(float64(ai-bi)) == 32) {
		return true
	} else {
		return false
	}

}

func isPalindrome(s string) bool {
	flag := true
	for i, j := 0, len(s)-1; i < j; {
		for ; i < j && !letterOrNum(s[i:i+1]); i++ {
		}
		for ; i < j && !letterOrNum(s[j:j+1]); j-- {
		}
		if (i < j) && !upperOrLower(s[i:i+1], s[j:j+1]) {
			flag = false
			break
		}
		i++
		j--
	}
	return flag
}

func singleNumber(nums []int) int {
	mapping := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := mapping[nums[i]]; ok {
			count := mapping[nums[i]]
			count++
			mapping[nums[i]] = count
		} else {
			mapping[nums[i]] = 1
		}
	}

	for num, count := range mapping {
		if count == 1 {
			return num
		}
	}
	return 0
}

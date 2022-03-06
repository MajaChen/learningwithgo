package leetcode

func canMeasureWater(jug1Capacity int, jug2Capacity int, targetCapacity int) bool {

	if jug1Capacity+jug2Capacity < targetCapacity {
		return false
	}

	if jug1Capacity <= 0 || jug2Capacity <= 0 {
		return false
	}

	var min int
	if min = jug1Capacity; min > jug2Capacity {
		min = jug2Capacity
	}
	for i := min; i > 0; i-- {
		if jug1Capacity%i == 0 && jug2Capacity%i == 0 {
			if targetCapacity%i == 0 {
				return true
			} else {
				return false
			}
		}
	}

	return true
}

func hasCycle(head *ListNode) bool {

	mapping := make(map[*ListNode]*ListNode)
	for p := head; p != nil; p = p.Next {
		if _, ok := mapping[p.Next]; ok {
			return true
		}

		mapping[p] = p.Next
	}

	return false
}

func detectCycle(head *ListNode) *ListNode {

	mapping := make(map[*ListNode]*ListNode)
	for p := head; p != nil; p = p.Next {
		if _, ok := mapping[p.Next]; ok {
			return p.Next
		}

		mapping[p] = p.Next
	}

	return nil

}

func detectCycle2(head *ListNode) *ListNode {

	var f, s *ListNode

	for f, s = head, head; f != nil && f.Next != nil; {
		f = f.Next.Next
		s = s.Next
		if f == s { // 这个地方不要放在循环的开头，弱智问题
			break
		}
	}
	if f == nil || f.Next == nil {
		return nil
	}

	for f = head; f != s; {
		f = f.Next
		s = s.Next
	}

	return s
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	l1, _ = reverseList(l1)
	l2, _ = reverseList(l2)
	l := addTwoNumbers2(l1, l2)
	l, _ = reverseList(l)
	return l
}

func findLeft(height []int, index int) int {

	for i := index - 1; i >= 0; i-- {
		if height[i] <= height[index] {
			continue
		}

		if i == 0 || height[i] > height[i-1] {
			return i
		}
	}

	return -1
}

func findRight(height []int, index int) int {

	for i := index + 1; i < len(height); i++ {
		if height[i] <= height[index] {
			continue
		}

		if i == len(height)-1 || height[i] > height[i+1] {
			return i
		}
	}

	return -1
}

func trap2(height []int) int {

	sum := 0
	for i := 1; i < len(height)-1; i++ {
		var leftIndex, rightIndex, min int
		if leftIndex = findLeft(height, i); leftIndex < 0 {
			continue
		}
		if rightIndex = findRight(height, i); rightIndex < 0 {
			continue
		}

		if min = height[leftIndex]; min > height[rightIndex] {
			min = height[rightIndex]
		}

		sum += min - height[i]
	}

	return sum
}

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {

	mapping, n := make(map[int]int), len(nums1)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum := nums1[i] + nums2[j]
			if _, ok := mapping[sum]; ok {
				count := mapping[sum]
				mapping[sum] = count + 1
			} else {
				mapping[sum] = 1
			}
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum := -(nums3[i] + nums4[j])
			if _, ok := mapping[sum]; ok {
				res += 1
			}
		}
	}

	return res
}

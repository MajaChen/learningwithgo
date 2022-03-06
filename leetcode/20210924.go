package leetcode

import "strconv"

//https://leetcode-solution-leetcode-pp.gitbook.io/leetcode-solution/medium/229.majority-element-ii
func majorityElementV2(nums []int) []int {
	res := make([]int, 0)
	majority, secondMajority, counterA, counterB := -1, -1, 0, 0
	for i := 0; i < len(nums); i++ {
		if counterA == 0 {
			counterA = 1
			majority = nums[i]
			continue
		} else if counterB == 0 && nums[i] != majority {
			counterB = 1
			secondMajority = nums[i]
			continue
		}
		if nums[i] != majority && nums[i] != secondMajority {
			counterA--
			counterB--
		} else if nums[i] == majority {
			counterA++
		} else {
			counterB++
		}
	}

	counterA, counterB = 0, 0
	for _, num := range nums {
		if num == majority {
			counterA++
		} else if num == secondMajority {
			counterB++
		}
	}
	if counterA > len(nums)/3 {
		res = append(res, majority)
	}
	if counterB > len(nums)/3 {
		res = append(res, secondMajority)
	}
	return res
}

func isDigit(r rune) bool {
	if r >= 48 && r <= 57 {
		return true
	}
	return false
}

//https://leetcode-solution-leetcode-pp.gitbook.io/leetcode-solution/medium/227.basic-calculator-ii
func calculate2(s string) int {

	runes := []rune(s)
	nums, ops := Stack{}, Stack{}
	for i := 0; i < len(s); {
		if runes[i] >= 48 && runes[i] <= 57 { // 数字
			var j int
			for j = i; j < len(runes) && isDigit(runes[j]); j++ {
			}
			digit, _ := strconv.Atoi(s[i:j])
			i = j

			nums.Push(digit)
			if ops.IsEmpty() {
				continue
			}
			op := ops.Pop().(int32)
			num := nums.Pop().(int)
			if op == 43 {
				num *= 1
			} else if op == 45 {
				num *= -1
			} else if op == 42 {
				num *= nums.Pop().(int)
			} else if op == 47 {
				num = nums.Pop().(int) / num
			}
			nums.Push(num)
		} else { //运算符
			ops.Push(runes[i])
			i++
		}
	}

	var sum int
	for !nums.IsEmpty() {
		sum += nums.Pop().(int)
	}
	return sum
}

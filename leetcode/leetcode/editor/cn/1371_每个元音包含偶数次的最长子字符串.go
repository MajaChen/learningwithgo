package leetcode

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
func findTheLongestSubstring(s string) int {

	//// 各元音首次出现的位置
	//// 哪个元音奇数个出现
	//mappingA,mappingB := make(map[string]int),make(map[string]bool)
	//mappingA["a"]=-1
	//mappingA["e"]=-1
	//mappingA["i"]=-1
	//mappingA["o"]=-1
	//mappingA["u"]=-1
	//
	//longestLength := math.MinInt32
	//
	//for i := 0;i < len(s);i++{
	//	c := s[i:i+1]
	//	if _,ok := mappingA[c];ok{// 元音
	//		if mappingA[c] < 0{
	//			mappingA[c] = i
	//		}
	//		if mappingB[c]{// 奇数次出现
	//			delete(mappingB,c)
	//		} else {
	//			mappingB[c] = true
	//		}
	//	}
	//
	//	startPos := 0
	//	for k,_ := range mappingB{
	//		if mappingA[k]+1 > startPos{
	//			startPos = mappingA[k]+1
	//		}
	//	}
	//
	//	if i-startPos+1 > longestLength{
	//		longestLength = i-startPos+1
	//	}
	//}
	//
	//return longestLength

	arr, ans, cur := make([]int, 32), 0, 0
	for i := 0; i < 32; i++ {
		arr[i] = math.MaxInt32
	}
	arr[0] = -1

	for i := 0; i < len(s); i++ {
		if c := s[i : i+1]; c == "a" {
			cur ^= 1
		} else if c == "e" {
			cur ^= 2
		} else if c == "i" {
			cur ^= 4
		} else if c == "o" {
			cur ^= 8
		} else if c == "u" {
			cur ^= 16
		}
		if arr[cur] == math.MaxInt32 {
			arr[cur] = i
		} else {
			if i-arr[cur] > ans {
				ans = i - arr[cur]
			}
		}
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

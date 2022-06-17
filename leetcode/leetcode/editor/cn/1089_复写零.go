package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func duplicateZeros(arr []int)  {
	copiedArr := make([]int, len(arr))
	copy(copiedArr, arr)
	for i, k := 0, 0; i < len(arr)&&k < len(arr); i++ {
		if copiedArr[i] != 0 {
			arr[k] = copiedArr[i]
			k++
		} else {
			arr[k] = 0
			if k++;k < len(arr) {
				arr[k] = 0
				k++
			}
		}
	}
	return
}
//leetcode submit region end(Prohibit modification and deletion)

package leetcode

func match(s string) bool {

	for i, j := 0, len(s); i < j; {
		if s[i:i+1] != s[j:j+1] {
			return false
		}
		i++
		j--
	}
	return true
}

func partitionRe(s string, startIndex int) [][]string {

	if startIndex >= len(s) {
		return [][]string{[]string{""}}
	}

	partitions := make([][]string, 0)
	for i := startIndex + 1; i <= len(s); i++ {
		if match(s[startIndex:i]) {
			subStrs := partitionRe(s, i)
			for index, strs := range subStrs {
				strs = append([]string{s[startIndex:i]}, strs...)
				subStrs[index] = strs
			}
			partitions = append(partitions, subStrs...)
		}
	}

	return partitions
}

func partition(s string) [][]string {

	return partitionRe(s, 0)
}

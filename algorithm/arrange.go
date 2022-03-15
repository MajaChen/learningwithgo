package algorithm

import (
	"strconv"
	"strings"
)

//去除0前缀
func removeZ(s string) string {

	runes := []rune(s)
	i := 0
	for ; i < len(runes) && runes[i] == 48; i++ {
	}
	if i == len(runes) {
		return "0"
	}
	return s[i:]
}

var computed map[int][]strings.Builder

// 数字全排列--递归
func arrangeNumsRe(n int) []strings.Builder {
	if n == 0 {
		computed[0] = make([]strings.Builder, 1)
		return make([]strings.Builder, 1)
	}
	res := make([]strings.Builder, 0)
	for i := 0; i < 10; i++ {
		var tmpRes []strings.Builder
		var ok bool
		if tmpRes, ok = computed[n-1]; !ok {
			tmpRes = arrangeNumsRe(n - 1)
		}

		for _, re := range tmpRes {
			pre := strings.Builder{}
			pre.WriteString(strconv.Itoa(i))
			pre.WriteString(re.String())
			res = append(res, pre)
		}
	}

	delete(computed, n-1)
	computed[n] = res
	return res
}

func arrangeNums(n int) []string {
	computed = make(map[int][]strings.Builder)
	res := arrangeNumsRe(n)
	strs := make([]string, 0, len(res))
	for _, re := range res {
		strs = append(strs, removeZ(re.String()))
	}
	return strs
}

// 数字全排列，非递归
func arrangeNumsNoRe(n int) []strings.Builder {
	res := make([]strings.Builder, 1)
	for i := 0; i < n; i++ {

		var sres []strings.Builder
		for j := 0; j < 10; j++ {
			for _, re := range res {
				var pre strings.Builder
				pre.WriteString(strconv.Itoa(j))
				pre.WriteString(re.String())
				sres = append(sres, pre)
			}
		}
		res = sres

	}
	return res
}

// 字母全排列
func arrangeAlphabetRe(n int) []strings.Builder {
	if n == 0 {
		return make([]strings.Builder, 1)
	}

	res := make([]strings.Builder, 0)
	for i := 0; i < 26; i++ {
		tmpRes := arrangeAlphabetRe(n - 1)
		for index, re := range tmpRes {
			pre := strings.Builder{}
			pre.WriteString(string(rune(i + 65)))
			pre.WriteString(re.String())
			tmpRes[index] = pre
		}
		res = append(res, tmpRes...)
	}
	return res
}

func arrangeAlphabet(n int) []string {
	res := arrangeAlphabetRe(2)
	strs := make([]string, len(res))
	for _, re := range res {
		strs = append(strs, re.String())
	}
	return strs
}

// 排列问题
// 8个选手，3座奖牌，按照金牌，银牌，铜牌的顺序列出所有可能的颁奖方式(n >= 3)
func arrange(candidates []string, n int, visited map[string]bool) []strings.Builder {
	if n == 0 {
		return make([]strings.Builder, 1)
	}

	res := make([]strings.Builder, 0)
	for i := 0; i < len(candidates); i++ {
		if visited[candidates[i]] {
			continue
		}
		visited[candidates[i]] = true
		tmpRes := arrange(candidates, n-1, visited)
		visited[candidates[i]] = false
		for index, re := range tmpRes {
			re.WriteString(candidates[i])
			tmpRes[index] = re
		}
		res = append(res, tmpRes...)
	}
	return res
}

// 组合问题
// 8个选手，3座奖牌，这里的奖牌是可乐瓶，列出所有可能的获奖方式
func combine(strs []string, n int) []strings.Builder {

	if n == 1 {
		res := make([]strings.Builder, 0, len(strs))
		for _, str := range strs {
			builder := strings.Builder{}
			builder.WriteString(str)
			res = append(res, builder)
		}
		return res
	}

	res := make([]strings.Builder, 0)
	for i := 0; i <= len(strs)-n; i++ {
		tmpRes := combine(strs[i+1:], n-1)
		for index, re := range tmpRes {
			re.WriteString(strs[i])
			tmpRes[index] = re
		}
		res = append(res, tmpRes...)
	}
	return res
}

package leetcode

import (
	"math"
)

func minSubArrayLen(target int, nums []int) int {

	var count int
	minimalLength := math.MaxInt32
	for i, l, r := 0, 0, 0; i < len(nums); i++ {
		r = i
		count += nums[i]
		if count >= target {
			// 尝试缩减左区间
			for l < r && count-nums[l] >= target {
				count -= nums[l]
				l++
			}
			if r-l+1 < minimalLength {
				minimalLength = r - l + 1
				if minimalLength == 1 {
					return minimalLength
				}
			}
		}
	}

	if minimalLength == math.MinInt32 {
		return 0
	}
	return minimalLength
}

func minSubArrayLen2(target int, nums []int) int {

	var count int
	minimalLength := math.MaxInt32
	for i, l, r := 0, 0, 0; i < len(nums); i++ {
		r = i
		count += nums[i]
		if count > target {
			// 尝试缩减左区间
			for l < r && count-nums[l] >= target {
				count -= nums[l]
				l++
			}
		}
		if count == target {
			if r-l+1 < minimalLength {
				minimalLength = r - l + 1
				if minimalLength == 1 {
					return minimalLength
				}
			}
		}
	}

	if minimalLength == math.MaxInt32 {
		return 0
	}
	return minimalLength
}

type WordDictionary struct {
	root *TrieNode
}

func WordDictionaryConstructor() WordDictionary {
	return WordDictionary{root: &TrieNode{val: "", children: make(map[string]*TrieNode)}}
}

func (this *WordDictionary) AddWord(word string) {

	node := this.root
	for index := 0; index < len(word); index++ {
		val := word[index : index+1]
		if _, ok := node.children[val]; !ok {
			childNode := &TrieNode{
				val:      val,
				children: make(map[string]*TrieNode),
			}
			node.children[val] = childNode
			node = childNode
		} else {
			node = node.children[val]
		}
	}
	node.isEnd = true

}

func searchWithIndex(word string, index int, parent *TrieNode) bool {

	if index >= len(word) {
		return parent.isEnd
	}

	if value := word[index : index+1]; value == "." {
		var flag bool
		for _, child := range parent.children {
			flag = flag || searchWithIndex(word, index+1, child)
		}
		return flag
	} else {
		if _, ok := parent.children[value]; !ok {
			return false
		}
		return searchWithIndex(word, index+1, parent.children[value])
	}

}

func (this *WordDictionary) Search(word string) bool {
	return searchWithIndex(word, 0, this.root)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := ConstructorTrie();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

func nextGreaterElement(nums []int) []int {

	indexes, stack := make([]int, len(nums)), Stack{elems: make([]interface{}, 0)}
	for i := len(nums) - 1; i >= 0; i-- {
		for !stack.IsEmpty() && nums[stack.GetTop().(int)] <= nums[i] {
			stack.Pop()
		}
		if stack.IsEmpty() {
			indexes[i] = -1
		} else {
			indexes[i] = stack.GetTop().(int)
		}
		stack.Push(i)
	}
	return indexes
}

func nextSmallerElement(nums []int) []int {

	indexes, stack := make([]int, len(nums)), Stack{elems: make([]interface{}, 0)}
	for i := len(nums) - 1; i >= 0; i-- {
		for !stack.IsEmpty() && nums[stack.GetTop().(int)] >= nums[i] {
			stack.Pop()
		}
		if stack.IsEmpty() {
			indexes[i] = -1
		} else {
			indexes[i] = stack.GetTop().(int)
		}
		stack.Push(i)
	}
	return indexes
}

func pos(nums []int, target int) int {

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}
	return -1
}

func nextGreaterElement2(nums1 []int, nums2 []int) []int {

	indexes := nextGreaterElement(nums2)
	res := make([]int, 0, len(nums1))
	for i := 0; i < len(nums1); i++ {
		if index := indexes[pos(nums2, nums1[i])]; index < 0 {
			res[i] = -1
		} else {
			res[i] = nums2[index]
		}
	}
	return res
}

func previousGreaterElement(nums []int) []int {

	indexes, stack := make([]int, len(nums)), Stack{elems: make([]interface{}, 0)}
	for i := 0; i < len(nums); i++ {
		for !stack.IsEmpty() && nums[stack.GetTop().(int)] <= nums[i] {
			stack.Pop()
		}
		if stack.IsEmpty() {
			indexes[i] = -1
			stack.Push(i)
		} else {
			indexes[i] = stack.GetTop().(int)
		}
	}

	return indexes
}

func greaterElement(nums []int) []int {

	nextIndexes := nextGreaterElement(nums)
	previousIndexes := previousGreaterElement(nums)

	for i := 0; i < len(nums); i++ {
		if nextIndexes[i] == -1 {
			if previousIndexes[i] != -1 {
				nextIndexes[i] = previousIndexes[i]
			}
		}

		if nextIndexes[i] != -1 {
			nextIndexes[i] = nums[nextIndexes[i]]
		}
	}

	return nextIndexes
}

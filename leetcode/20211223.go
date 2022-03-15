package leetcode

import "math"

// 1 2/3
func rangeBitwiseAnd(left int, right int) int {

	var res int
	for left > 0 {
		var i int
		if i = int(math.Exp2(math.Floor(math.Log2(float64(left))))); 2*i <= right {
			return res
		} else {
			res += i
		}

		left -= i
		right -= i
	}
	return res
}

type Trie struct {
	root *TrieNode
}

func ConstructorTrie() Trie {
	return Trie{root: &TrieNode{val: "", children: make(map[string]*TrieNode)}}
}

func (this *Trie) Insert(word string) {

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

func (this *Trie) Search(word string) bool {

	node := this.root
	for index := 0; index < len(word); index++ {
		if _, ok := node.children[word[index:index+1]]; !ok {
			return false
		}

		node = node.children[word[index:index+1]]
	}
	return node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {

	for index, node := 0, this.root; index <= len(prefix); index++ {
		if index == len(prefix) {
			return true
		}

		if _, ok := node.children[prefix[index:index+1]]; !ok {
			return false
		}

		node = node.children[prefix[index:index+1]]
	}
	return false
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := QueueConstructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

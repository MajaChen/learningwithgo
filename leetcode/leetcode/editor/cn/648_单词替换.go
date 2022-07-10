package leetcode

import "strings"

//leetcode submit region begin(Prohibit modification and deletion)

type trieNode struct {
	nodes [26]*trieNode
	index int
}

var root *trieNode

func insert(i int, dictionary []string) {
	word := dictionary[i]
	cur := root
	for _, r := range word {
		if cur.nodes[r-97] == nil {
			cur.nodes[r-97] = &trieNode{index: -1}
		}
		cur = cur.nodes[r-97]
	}
	cur.index = i
}

func replace(word string, dictionary []string) string {
	cur := root
	for _, r := range word {
		if cur.nodes[r-97] == nil {
			return word
		}
		if cur = cur.nodes[r-97]; cur.index >= 0 {
			return dictionary[cur.index]
		}
	}
	return word
}

func replaceWords(dictionary []string, sentence string) string {
	root = &trieNode{index: -1}
	for i := range dictionary {
		insert(i, dictionary)
	}
	reps := make([]string, 0)
	for _, word := range strings.Split(sentence, " ") {
		reps = append(reps, replace(word, dictionary))
	}
	return strings.Join(reps, " ")
}

//leetcode submit region end(Prohibit modification and deletion)

package leetcode

//leetcode submit region begin(Prohibit modification and deletion)

type MagicDictionary struct {
	entries map[rune]*MagicDictionary
	isEnd   bool
}

func MagicDictionaryConstructor() MagicDictionary {
	return MagicDictionary{entries: make(map[rune]*MagicDictionary)}
}

func insertDictionary(word string, root *MagicDictionary) {
	cur := root
	for _, r := range word {
		if cur.entries[r-97] == nil {
			cur.entries[r-97] = &MagicDictionary{entries: make(map[rune]*MagicDictionary)}
		}
		cur = cur.entries[r-97]
	}
	cur.isEnd = true
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, entry := range dictionary {
		insertDictionary(entry, this)
	}
}

func search(word string, root *MagicDictionary, i int, hasDiff bool) bool {
	if i == len(word) && root.isEnd && hasDiff {
		return true
	}

	if i >= len(word) {
		return false
	}

	for k, v := range root.entries {
		var res bool
		if k != rune(word[i]-97) {
			if !hasDiff {
				res = search(word, v, i+1, true)
			}
		} else {
			res = search(word, v, i+1, hasDiff)
		}
		if res {
			return true
		}
	}

	return false
}

func (this *MagicDictionary) Search(searchWord string) bool {
	return search(searchWord, this, 0, false)
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := MyCalendarConstructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
//leetcode submit region end(Prohibit modification and deletion)

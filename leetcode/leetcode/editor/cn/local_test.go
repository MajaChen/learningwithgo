package leetcode

import "testing"

func TestXXX(t *testing.T) {
	finder := Constructor()
	finder.AddNum(6)
	finder.AddNum(10)
	finder.AddNum(2)
	finder.AddNum(6)
	finder.AddNum(5)
	finder.AddNum(0)
	finder.AddNum(6)
	finder.AddNum(3)
	finder.AddNum(1)
	finder.AddNum(0)
	finder.AddNum(0)
	t.Error(finder.FindMedian())
}

func TestMap(t *testing.T) {
	mapping := make(map[int]bool)
	mapping[0] = true
	t.Error(mapping[1])
}

func TestSlice(t *testing.T) {
	s := make([]int, 0, 2)
	s = append(s, 1)
	s = append(s, 2)
	t.Error(s[1:1])
}

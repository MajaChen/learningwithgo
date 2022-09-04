package common

type DisjointSet struct {
	parents map[interface{}]interface{}
	size    map[interface{}]int
}

func InitDisjointSet() *DisjointSet {
	return &DisjointSet{parents: make(map[interface{}]interface{}), size: make(map[interface{}]int)}
}

func (set *DisjointSet) Find(x interface{}) interface{} {
	if set.parents[x] == nil {
		set.parents[x] = x
		set.size[x] = 1
		return x
	}
	if set.parents[x] == x {
		return x
	}
	set.parents[x] = set.Find(set.parents[x])
	return set.parents[x]
}

func (set *DisjointSet) Union(x, y interface{}) {
	rx, ry := set.Find(x), set.Find(y)
	if rx == ry {
		return
	}

	if set.size[rx] > set.size[ry] {
		set.parents[ry] = set.parents[rx]
		set.size[rx] = set.size[rx] + set.size[ry]
	} else {
		set.parents[rx] = ry
		set.size[ry] = set.size[ry] + set.size[rx]
	}
	return
}

package leetcode

type Stack struct {
	elems []interface{}
}

func (s *Stack) IsEmpty() bool {
	return len(s.elems) == 0
}

func (s *Stack) len() int {
	return len(s.elems)
}

func (s *Stack) Pop() interface{} {
	elem := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]
	return elem
}

func (s *Stack) Push(elem interface{}) {
	s.elems = append(s.elems, elem)
}

func (s *Stack) GetTop() interface{} {
	return s.elems[len(s.elems)-1]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swap(i, j int, nums []int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}

type Queue struct {
	elems []interface{}
}

func (q *Queue) Size() int {
	return len(q.elems)
}

func (q *Queue) Pop() interface{} {
	elem := q.elems[0]
	q.elems = q.elems[1:]
	return elem
}

func (q *Queue) Push(elem interface{}) {
	q.elems = append(q.elems, elem)
}

func (q Queue) IsEmpty() bool {
	return len(q.elems) == 0
}

type Item struct {
	color int //0 白色，1 灰色
	Node  *TreeNode
}

type TrieNode struct {
	val      string
	children map[string]*TrieNode
	isEnd    bool
}

type DisjointSet struct {
	parents map[interface{}]interface{}
	size    map[interface{}]int
}

func InitDisjointSet() *DisjointSet {
	return &DisjointSet{parents: make(map[interface{}]interface{}), size: make(map[interface{}]int)}
}

func (set *DisjointSet) Find(elem interface{}) interface{} {

	var parent interface{}
	if parent = set.parents[elem]; parent == elem {
		return elem
	}
	root := set.Find(parent)
	set.parents[parent] = root
	return root
}

func (set *DisjointSet) Add(elem, parent interface{}) {

	if _, ok := set.parents[parent]; !ok {
		set.parents[parent] = parent
	}
	set.parents[elem] = parent
	set.size[parent] = 1
}

func (set *DisjointSet) Union(elemA, elemB interface{}) {

	if set.Connected(elemA, elemB) {
		return
	}

	pa, pb := set.Find(elemA), set.Find(elemB)
	if set.size[pa] > set.size[pb] {
		set.parents[pb] = pa
		delete(set.size, pb)
	} else {
		set.parents[pa] = pb
		delete(set.size, pa)
	}
	return
}

func (set DisjointSet) Connected(elemA, elemB interface{}) bool {

	return set.Find(elemA) == set.Find(elemB)
}

func (set DisjointSet) RootCount() int {

	return len(set.size)
}

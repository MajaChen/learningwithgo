package leetcode

import (
	"math"
	"sort"
)

type RLEIterator struct {
	elements []int
	index    int
}

func RLEIteratorConstructor(encoding []int) RLEIterator {
	return RLEIterator{elements: encoding, index: 0}
}

func (this *RLEIterator) Next(n int) int {

	for ; this.index < len(this.elements)-1 && n > this.elements[this.index]; this.index += 2 {
		n -= this.elements[this.index]
	}
	if this.index >= len(this.elements)-1 {
		return -1
	} else {
		this.elements[this.index] -= n
		elem := this.elements[this.index+1]
		if this.elements[this.index] <= 0 {
			this.index += 2
		}

		return elem
	}
}

type Candidate struct {
	name  int
	count int
}

type Leader struct {
	Candidate
	time int
}

type TopVotedCandidate struct {
	leaders []*Leader
}

type Vote struct {
	time  int
	count int
}

func elect(persons []int, times []int) []*Leader {

	leaders := make([]*Leader, 0)
	mapping, leadingCandidate := make(map[int]*Candidate), &Candidate{name: -1, count: 0}
	for i := 0; i < len(persons); i++ {
		var c *Candidate
		if _, ok := mapping[persons[i]]; ok {
			c = mapping[persons[i]]
			c.count += 1
		} else {
			c = &Candidate{name: persons[i], count: 1}
		}
		mapping[persons[i]] = c

		if c.count >= leadingCandidate.count {
			leadingCandidate = c
		}

		var leader Leader
		leader.time, leader.name, leader.count = times[i], leadingCandidate.name, leadingCandidate.count
		leaders = append(leaders, &leader)
	}

	return leaders
}

func Constructor(persons []int, times []int) TopVotedCandidate {

	return TopVotedCandidate{leaders: elect(persons, times)}
}

func (this *TopVotedCandidate) Q(t int) int {
	i := sort.Search(len(this.leaders), func(i int) bool {
		return this.leaders[i].time >= t
	})

	if i >= len(this.leaders) || this.leaders[i].time != t {
		i -= 1
	}

	return this.leaders[i].name
}

func BinarySearch(n int, f func(int) bool) int {
	// Define f(-1) == false and f(n) == true.
	// Invariant: f(i-1) == false, f(j) == true.
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		// i ≤ h < j
		if !f(h) {
			i = h + 1 // preserves f(i-1) == false
		} else {
			j = h // preserves f(j) == true
		}
	}
	// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
	return i
}

func findPivot(nums []int) int {

	pivot, pivotIndex := nums[0], 0
	for low, high := 0, len(nums)-1; low < high; {
		for ; high > low && nums[high] >= pivot; high-- {
		}
		swap(pivotIndex, high, nums)
		pivotIndex = high
		for ; high > low && nums[low] <= pivot; low++ {
		}
		swap(pivotIndex, low, nums)
		pivotIndex = low
	}
	nums[pivotIndex] = pivot
	return pivotIndex
}

func sortArrayRe(nums []int) {

	if len(nums) <= 1 {
		return
	}

	pivot := findPivot(nums)
	sortArrayRe(nums[:pivot])
	sortArrayRe(nums[pivot+1:])
}

func sortArray(nums []int) []int {

	sortArrayRe(nums)
	return nums
}

var dialerTraverseMapping map[int][]int
var dialerVisitMapping [][]int

func knightDialerRe(n, i int) int {

	if n == 0 {
		return 1
	}

	if dialerVisitMapping[i][n] > 0 {
		return dialerVisitMapping[i][n]
	}

	count := 0
	for i, nums := 0, dialerTraverseMapping[i]; i < len(nums); i++ {
		count += knightDialerRe(n-1, nums[i])
	}

	dialerVisitMapping[i][n] = count
	return count
}

func knightDialer(n int) int {

	dialerTraverseMapping = map[int][]int{
		1: []int{6, 8},
		2: []int{7, 9},
		3: []int{4, 8},
		4: []int{0, 3, 9},
		5: []int{},
		6: []int{0, 1, 7},
		7: []int{2, 6},
		8: []int{1, 3},
		9: []int{2, 4},
		0: []int{4, 6},
	}

	dialerVisitMapping = make([][]int, 0)
	for i := 0; i < 10; i++ {
		dialerVisitMapping = append(dialerVisitMapping, make([]int, n))
	}

	count := 0
	for i := 0; i < 10; i++ {
		count += knightDialerRe(n-1, i)
	}

	return count
}

func knightDialer2(n int) int {

	var a0, a1, a2, a3, a4, a5, a6, a7, a8, a9 int64 = 1, 1, 1, 1, 1, 1, 1, 1, 1, 1
	k := int64(math.Pow(float64(10), float64(9)) + 7)
	for i := 1; i <= n-1; i++ {
		a0, a1, a2, a3, a4, a5, a6, a7, a8, a9 = (a4+a6)%k, (a6+a8)%k, (a7+a9)%k, (a4+a8)%k, (a0+a3+a9)%k, 0, (a0+a1+a7)%k, (a2+a6)%k, (a1+a3)%k, (a2+a4)%k
	}

	return int((a0 + a1 + a2 + a3 + a4 + a5 + a6 + a7 + a8 + a9) % k)
}

func maxTurbulenceSize(arr []int) int {

	if len(arr) <= 2 {
		if len(arr) == 2 && arr[0] == arr[1] {
			return 1
		}
		return len(arr)
	}

	maximalLength := 1
	for i, j := 0, 2; j < len(arr); j++ {
		if arr[j]-arr[j-1] == 0 {
			i = j
		} else if (arr[j]-arr[j-1])^(arr[j-1]-arr[j-2]) >= 0 {
			i = j - 1
		}
		if (j - i + 1) > maximalLength {
			maximalLength = j - i + 1
		}
	}

	return maximalLength

}

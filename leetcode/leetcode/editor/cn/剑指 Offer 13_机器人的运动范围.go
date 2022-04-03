package leetcode

//leetcode submit region begin(Prohibit modification and deletion)

var visitedNodes [][]int

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

func isOutOfBound(i, j, k int) bool {

	sumi, sumj := 0, 0
	for i > 0 {
		sumi += i % 10
		i /= 10
	}
	for j > 0 {
		sumj += j % 10
		j /= 10
	}

	return sumi+sumj > k
}

func addNode(m, n, i, j, k int, queue *Queue) {

	if i < 0 || i >= m || j < 0 || j >= n || visitedNodes[i][j] == 1 || isOutOfBound(i, j, k) {
		return
	}
	visitedNodes[i][j] = 1
	queue.Push([]int{i, j})
}

func movingCount(m int, n int, k int) int {

	visitedNodes = make([][]int, 0, m)
	for i := 0; i < m; i++ {
		visitedNodes = append(visitedNodes, make([]int, n))
	}

	queue, counter := &Queue{make([]interface{}, 0)}, 0
	visitedNodes[0][0] = 1
	queue.Push([]int{0, 0})
	for !queue.IsEmpty() {
		node := queue.Pop().([]int)
		counter++
		addNode(m, n, node[0]+1, node[1], k, queue)
		addNode(m, n, node[0]-1, node[1], k, queue)
		addNode(m, n, node[0], node[1]+1, k, queue)
		addNode(m, n, node[0], node[1]-1, k, queue)
	}

	return counter
}

//leetcode submit region end(Prohibit modification and deletion)

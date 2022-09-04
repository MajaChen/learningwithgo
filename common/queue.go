package common

type Queue struct {
	Elems []interface{}
}

func (q *Queue) Size() int {
	return len(q.Elems)
}

func (q *Queue) Pop() interface{} {
	elem := q.Elems[0]
	q.Elems = q.Elems[1:]
	return elem
}

func (q *Queue) Push(elem interface{}) {
	q.Elems = append(q.Elems, elem)
}

func (q Queue) IsEmpty() bool {
	return len(q.Elems) == 0
}

func (q Queue) GetTop() interface{} {
	return q.Elems[0]
}

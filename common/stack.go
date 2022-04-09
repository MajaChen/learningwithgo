package common

type Stack struct {
	Elems []interface{}
}

func (s *Stack) IsEmpty() bool {
	return len(s.Elems) == 0
}

func (s *Stack) Len() int {
	return len(s.Elems)
}

func (s *Stack) Pop() interface{} {
	elem := s.Elems[len(s.Elems)-1]
	s.Elems = s.Elems[:len(s.Elems)-1]
	return elem
}

func (s *Stack) Push(elem interface{}) {
	s.Elems = append(s.Elems, elem)
}

func (s *Stack) GetTop() interface{} {
	return s.Elems[len(s.Elems)-1]
}

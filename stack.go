package generic

type simpleStack struct {
	top    *stackNode
	length int
}

type stackNode struct {
	value interface{}
	prev  *stackNode
}

func NewStack() Stack {
	return &simpleStack{nil, 0}
}

func (s *simpleStack) Len() int {
	return s.length
}

func (s *simpleStack) Peek() interface{} {
	if s.length == 0 {
		return nil
	}
	return s.top.value
}

func (s *simpleStack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}

	n := s.top
	s.top = n.prev
	s.length--
	return n.value
}

func (s *simpleStack) Push(value interface{}) {
	n := &stackNode{value, s.top}
	s.top = n
	s.length++
}

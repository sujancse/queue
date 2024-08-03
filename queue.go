package queue

type node[T any] struct {
	next *node[T]
	data T
}

type Queue[T any] struct {
	first *node[T]
	last  *node[T]
}

func (s *Queue[T]) Add(data T) {
	t := &node[T]{
		data: data,
		next: nil,
	}

	if s.last == nil {
		s.last = t
		s.first = t
	} else {
		s.last.next = t
		s.last = s.last.next
	}
}

func (s *Queue[T]) Remove() T {
	data := s.first.data
	s.first = s.first.next

	if s.first == nil {
		s.last = nil
	}

	return data
}

func (s *Queue[T]) Peek() T {
	return s.first.data
}

func (s *Queue[T]) IsEmpty() bool {
	return s.first == nil
}

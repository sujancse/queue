package queue

type node[T any] struct {
	next *node[T]
	data T
}

type Queue[T any] struct {
	first *node[T]
	last  *node[T]
}

func (q *Queue[T]) Add(data T) {
	t := &node[T]{
		data: data,
		next: nil,
	}

	if q.last == nil {
		q.last = t
		q.first = t
	} else {
		q.last.next = t
		q.last = q.last.next
	}
}

func (q *Queue[T]) Remove() T {
	data := q.first.data
	q.first = q.first.next

	if q.first == nil {
		q.last = nil
	}

	return data
}

func (q *Queue[T]) Peek() T {
	return q.first.data
}

func (q *Queue[T]) IsEmpty() bool {
	return q.first == nil
}

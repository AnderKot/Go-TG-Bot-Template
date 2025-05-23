package Queue

type MyQueue[T any] struct {
	start  *node[T]
	end    *node[T]
	length int
}

type node[T any] struct {
	value T
	next  *node[T]
}

func New[T any]() *MyQueue[T] {
	return &MyQueue[T]{nil, nil, 0}
}

func (q *MyQueue[T]) Dequeue() (bool, *T) {
	if q.length == 0 {
		return false, nil
	}
	n := q.start
	if q.length == 1 {
		q.start = nil
		q.end = nil
	} else {
		q.start = q.start.next
	}
	q.length--
	return true, &(n.value)
}

func (q *MyQueue[T]) Enqueue(value *T) {
	n := &(node[T]{value: *(value)})
	if q.length == 0 {
		q.start = n
		q.end = n
	} else {
		q.end.next = n
		q.end = n
	}
	q.length++
}

func (q *MyQueue[T]) Len() int {
	return q.length
}

func (q *MyQueue[T]) Peek() (bool, *T) {
	if q.length == 0 {
		return false, nil
	}
	return true, &(q.start.value)
}

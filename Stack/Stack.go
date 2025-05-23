package Stack

type MyStack[T any] struct {
	start  *node[T]
	length int
}

type node[T any] struct {
	value T
	pre   *node[T]
}

func New[T any]() *MyStack[T] {
	return &MyStack[T]{nil, 0}
}

func (q *MyStack[T]) DeStack() (bool, *T) {
	if q.length == 0 {
		return false, nil
	}
	n := q.start
	if q.length == 1 {
		q.start = nil
	} else {
		q.start = q.start.pre
	}
	q.length--
	return true, &(n.value)
}

func (q *MyStack[T]) AddStack(value *T) {
	n := &(node[T]{value: *(value)})
	if q.length == 0 {
		q.start = n
	} else {
		n.pre = q.start
		q.start = n
	}
	q.length++
}

func (q *MyStack[T]) Len() int {
	return q.length
}

func (q *MyStack[T]) Peek() (bool, *T) {
	if q.length == 0 {
		return false, nil
	}
	return true, &(q.start.value)
}

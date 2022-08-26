package gollect

type Queue[T any] struct {
	data Vector[T]
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{data: NewVector[T]()}
}

func NewQueueFromData[T any](values ...T) Queue[T] {
	return Queue[T]{data: NewVectorFromData(values...)}
}

func NewQueueFromDataRef[T any](values ...*T) Queue[T] {
	return Queue[T]{data: NewVectorFromDataRef(values...)}
}

func NewQueueFromQueue[T any](other Queue[T]) Queue[T] {
	return Queue[T]{data: NewVectorFromVector(other.data)}
}

func MakeQueue[T any]() *Queue[T] {
	return &Queue[T]{data: NewVector[T]()}
}

func MakeQueueFromData[T any](values ...T) *Queue[T] {
	return &Queue[T]{data: NewVectorFromData(values...)}
}

func MakeQueueFromDataRef[T any](values ...*T) *Queue[T] {
	return &Queue[T]{data: NewVectorFromDataRef(values...)}
}

func MakeQueueFromQueue[T any](other Queue[T]) *Queue[T] {
	return &Queue[T]{data: NewVectorFromVector(other.data)}
}

func (q *Queue[T]) Front() T {
	return q.data.Front()
}

func (q *Queue[T]) FrontRef() *T {
	return q.data.FrontRef()
}

func (q *Queue[T]) Data() []T {
	return q.data.Data()
}

func (q *Queue[T]) IsEmpty() bool {
	return q.data.IsEmpty()
}

func (q *Queue[T]) Size() int {
	return q.data.Size()
}

func (q *Queue[T]) Clear() {
	q.data.Clear()
}

func (q *Queue[T]) PushBack(value T) {
	q.data.PushBack(value)
}

func (q *Queue[T]) PushBackRef(value *T) {
	q.data.PushBackRef(value)
}

func (q *Queue[T]) PopFront() {
	q.data.PopFront()
}

func (q *Queue[T]) Swap(other *Queue[T]) {
	q.data.Swap(&other.data)
}

func (q *Queue[T]) String() string {
	return q.data.String()
}

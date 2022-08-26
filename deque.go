package gollect

type Deque[T any] struct {
	data Vector[T]
}

func NewDeque[T any]() Deque[T] {
	return Deque[T]{data: NewVector[T]()}
}

func NewDequeFromData[T any](values ...T) Deque[T] {
	return Deque[T]{data: NewVectorFromData(values...)}
}

func NewDequeFromDataRef[T any](values ...*T) Deque[T] {
	return Deque[T]{data: NewVectorFromDataRef(values...)}
}

func NewDequeFromDeque[T any](other Deque[T]) Deque[T] {
	return Deque[T]{data: NewVectorFromVector(other.data)}
}

func MakeDeque[T any]() *Deque[T] {
	return &Deque[T]{data: NewVector[T]()}
}

func MakeDequeFromData[T any](values ...T) *Deque[T] {
	return &Deque[T]{data: NewVectorFromData(values...)}
}

func MakeDequeFromDataRef[T any](values ...*T) *Deque[T] {
	return &Deque[T]{data: NewVectorFromDataRef(values...)}
}

func MakeDequeFromDeque[T any](other Deque[T]) *Deque[T] {
	return &Deque[T]{data: NewVectorFromVector(other.data)}
}

func (d *Deque[T]) Front() T {
	return d.data.Front()
}

func (d *Deque[T]) FrontRef() *T {
	return d.data.FrontRef()
}

func (d *Deque[T]) Back() T {
	return d.data.Back()
}

func (d *Deque[T]) BackRef() *T {
	return d.data.BackRef()
}

func (d *Deque[T]) Data() []T {
	return d.data.Data()
}

func (d *Deque[T]) IsEmpty() bool {
	return d.data.IsEmpty()
}

func (d *Deque[T]) Size() int {
	return d.data.Size()
}

func (d *Deque[T]) Clear() {
	d.data.Clear()
}

func (d *Deque[T]) PushBack(value T) {
	d.data.PushBack(value)
}

func (d *Deque[T]) PushBackRef(value *T) {
	d.data.PushBackRef(value)
}

func (d *Deque[T]) PushFront(value T) {
	d.data.PushFront(value)
}

func (d *Deque[T]) PushFrontRef(value *T) {
	d.data.PushFrontRef(value)
}

func (d *Deque[T]) PopBack() {
	d.data.PopBack()
}

func (d *Deque[T]) PopFront() {
	d.data.PopFront()
}

func (d *Deque[T]) Swap(other *Deque[T]) {
	d.data.Swap(&other.data)
}

func (d *Deque[T]) String() string {
	return d.data.String()
}

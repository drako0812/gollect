package gollect

type Stack[T any] struct {
	data Vector[T]
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{data: NewVector[T]()}
}

func NewStackFromData[T any](values ...T) Stack[T] {
	return Stack[T]{data: NewVectorFromData(values...)}
}

func NewStackFromDataRef[T any](values ...*T) Stack[T] {
	return Stack[T]{data: NewVectorFromDataRef(values...)}
}

func NewStackFromStack[T any](other Stack[T]) Stack[T] {
	return Stack[T]{data: NewVectorFromVector(other.data)}
}

func MakeStack[T any]() *Stack[T] {
	return &Stack[T]{data: NewVector[T]()}
}

func MakeStackFromData[T any](values ...T) *Stack[T] {
	return &Stack[T]{data: NewVectorFromData(values...)}
}

func MakeStackFromDataRef[T any](values ...*T) *Stack[T] {
	return &Stack[T]{data: NewVectorFromDataRef(values...)}
}

func MakeStackFromStack[T any](other Stack[T]) *Stack[T] {
	return &Stack[T]{data: NewVectorFromVector(other.data)}
}

func (d *Stack[T]) Top() T {
	return d.data.Back()
}

func (d *Stack[T]) TopRef() *T {
	return d.data.BackRef()
}

func (d *Stack[T]) Data() []T {
	return d.data.Data()
}

func (d *Stack[T]) IsEmpty() bool {
	return d.data.IsEmpty()
}

func (d *Stack[T]) Size() int {
	return d.data.Size()
}

func (d *Stack[T]) Clear() {
	d.data.Clear()
}

func (d *Stack[T]) Push(value T) {
	d.data.PushBack(value)
}

func (d *Stack[T]) PushRef(value *T) {
	d.data.PushBackRef(value)
}

func (d *Stack[T]) Pop() {
	d.data.PopBack()
}

func (d *Stack[T]) Swap(other *Stack[T]) {
	d.data.Swap(&other.data)
}

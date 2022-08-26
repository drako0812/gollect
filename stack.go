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

func (s *Stack[T]) Top() T {
	return s.data.Back()
}

func (s *Stack[T]) TopRef() *T {
	return s.data.BackRef()
}

func (s *Stack[T]) Data() []T {
	return s.data.Data()
}

func (s *Stack[T]) IsEmpty() bool {
	return s.data.IsEmpty()
}

func (s *Stack[T]) Size() int {
	return s.data.Size()
}

func (s *Stack[T]) Clear() {
	s.data.Clear()
}

func (s *Stack[T]) Push(value T) {
	s.data.PushBack(value)
}

func (s *Stack[T]) PushRef(value *T) {
	s.data.PushBackRef(value)
}

func (s *Stack[T]) Pop() {
	s.data.PopBack()
}

func (s *Stack[T]) Swap(other *Stack[T]) {
	s.data.Swap(&other.data)
}

func (s *Stack[T]) String() string {
	return s.data.String()
}

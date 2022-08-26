package gollect

import "fmt"

// Destructible identifies a type that should have some kind of tear-down code executed when they are no longer being used.
type Destructible interface {
	Destruct()
}

// CollectionVisitor is a function that will be called on every element of a collection.
type CollectionVisitor[T any] func(*T, *bool)

// GeneralCollector identifies a general-purpose collection type.
type GeneralCollector[T any] interface {
	Front() T
	FrontRef() *T
	Back() T
	BackRef() *T
	IsEmpty() bool
	Size() int
	Clear()
	Insert(index int, value T)
	InsertRef(index int, value *T)
	Erase(index int)
	PushBack(value T)
	PushBackRef(value *T)
	PushFront(value T)
	PushFrontRef(value *T)
	PopBack()
	PopFront()
	Visit(visitor CollectionVisitor[T])
	VisitReverse(visitor CollectionVisitor[T])
	ContainsValue(value T) bool
	ContainsRef(value *T) bool
	OrderedSearch(value T) (found bool, index int)
	OrderedRefSearch(value *T) (found bool, index int)
	OrderedSearchRef(value T) *T
	OrderedRefSearchRef(value *T) *T
	Search(value T) (found bool, index int)
	RefSearch(value *T) (found bool, index int)
	SearchRef(value T) *T
	RefSearchRef(value *T) *T
	fmt.Stringer
}

// NativeEquatable identifies the set of types that can have the `==` and `!=` operators used on them.
type NativeEquatable interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64 | ~complex64 | ~complex128 | ~string | ~bool
}

// NativeComparable identifies the set of types that can have the `<`, `>`, `<=`, and `>=` operators used on them.
type NativeComparable interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64 | ~string
}

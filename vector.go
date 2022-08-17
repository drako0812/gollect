package gollect

import (
	"sort"

	"golang.org/x/exp/constraints"
)

type Destructible interface {
	Destruct()
}

type Vector[T any] struct {
	data []T
}

func NewVector[T any]() Vector[T] {
	return Vector[T]{data: []T{}}
}

func NewVectorFromData[T any](values ...T) Vector[T] {
	return Vector[T]{data: values}
}

func NewVectorFromDataRef[T any](values ...*T) Vector[T] {
	v := NewVector[T]()
	for _, val := range values {
		v.PushBackRef(val)
	}
	return v
}

func NewVectorFromVector[T any](other Vector[T]) Vector[T] {
	v := NewVectorFromData(other.data...)
	//copy(v.data, other.data)
	return v
}

func MakeVector[T any]() *Vector[T] {
	return &Vector[T]{data: []T{}}
}

func MakeVectorFromData[T any](values ...T) *Vector[T] {
	return &Vector[T]{data: values}
}

func MakeVectorFromDataRef[T any](values ...*T) *Vector[T] {
	v := MakeVector[T]()
	for _, val := range values {
		v.PushBackRef(val)
	}
	return v
}

func MakeVectorFromVector[T any](other Vector[T]) *Vector[T] {
	v := MakeVectorFromData(other.data...)
	return v
}

func (v *Vector[T]) At(index int) T {
	return v.data[index]
}

func (v *Vector[T]) SafeAt(index int) T {
	if len(v.data) > 0 {
		if (index >= 0) && (index < len(v.data)) {
			return v.data[index]
		}
		panic("ERROR: Vector.SafeAt - index out of range")
	}
	panic("ERROR: Vector.SafeAt - empty vector")
}

func (v *Vector[T]) AtRef(index int) *T {
	return &v.data[index]
}

func (v *Vector[T]) SafeAtRef(index int) *T {
	if len(v.data) > 0 {
		if (index >= 0) && (index < len(v.data)) {
			return &v.data[index]
		}
		panic("ERROR: Vector.SafeAtRef - index out of range")
	}
	panic("ERROR: Vector.SafeAtRef - empty vector")
}

func (v *Vector[T]) Front() T {
	if len(v.data) > 0 {
		return v.data[0]
	}
	panic("ERROR: Vector.Front - empty vector")
}

func (v *Vector[T]) FrontRef() *T {
	if len(v.data) > 0 {
		return &v.data[0]
	}
	panic("ERROR: Vector.FrontRef - empty vector")
}

func (v *Vector[T]) Back() T {
	if len(v.data) > 0 {
		return v.data[len(v.data)-1]
	}
	panic("ERROR: Vector.Back - empty vector")
}

func (v *Vector[T]) BackRef() *T {
	if len(v.data) > 0 {
		return &v.data[len(v.data)-1]
	}
	panic("ERROR: Vector.BackRef - empty vector")
}

func (v *Vector[T]) Data() []T {
	return v.data
}

func (v *Vector[T]) IsEmpty() bool {
	return len(v.data) == 0
}

func (v *Vector[T]) Size() int {
	return len(v.data)
}

func (v *Vector[T]) Clear() {
	for !v.IsEmpty() {
		v.PopBack()
	}
}

func (v *Vector[T]) Insert(index int, value T) {
	if (index < 0) || (index > v.Size()) {
		panic("ERROR: Vector.Insert - index out of bounds")
	}
	if v.Size() != 0 {
		if index == v.Size() {
			v.data = append(v.data, value)
		} else {
			v.data = append(v.data[:index+1], v.data[index:]...)
			v.data[index] = value
		}
	} else {
		v.data = []T{value}
	}
}

func (v *Vector[T]) InsertRef(index int, value *T) {
	if (index < 0) || (index > v.Size()) {
		panic("ERROR: Vector.Insert - index out of bounds")
	}
	if v.Size() != 0 {
		if index == v.Size() {
			v.data = append(v.data, *value)
		} else {
			v.data = append(v.data[:index+1], v.data[index:]...)
			v.data[index] = *value
		}
	} else {
		v.data = []T{*value}
	}
}

func (v *Vector[T]) Erase(index int) {
	if !v.IsEmpty() {
		if (index >= 0) && (index < v.Size()) {
			if e, isDestructible := interface{}(v.AtRef(index)).(Destructible); isDestructible {
				e.Destruct()
			}
			v.data = append(v.data[:index], v.data[index+1:]...)
		} else {
			panic("ERROR: Vector.Erase - index out of bounds")
		}
	} else {
		panic("ERROR: Vector.Erase - empty vector")
	}
}

func (v *Vector[T]) PushBack(value T) {
	v.data = append(v.data, value)
}

func (v *Vector[T]) PushBackRef(value *T) {
	v.data = append(v.data, *value)
}

func (v *Vector[T]) PushFront(value T) {
	v.data = append([]T{value}, v.data...)
}

func (v *Vector[T]) PushFrontRef(value *T) {
	v.data = append([]T{*value}, v.data...)
}

func (v *Vector[T]) PopBack() {
	if !v.IsEmpty() {
		if e, isDestructible := interface{}(v.BackRef()).(Destructible); isDestructible {
			e.Destruct()
		}
		v.data = v.data[:len(v.data)-1]
	} else {
		panic("ERROR: Vector.PopBack - empty vector")
	}
}

func (v *Vector[T]) PopFront() {
	if !v.IsEmpty() {
		if e, isDestructible := interface{}(v.FrontRef()).(Destructible); isDestructible {
			e.Destruct()
		}
		v.data = v.data[1:]
	} else {
		panic("ERROR: Vector.PopFront - empty vector")
	}
}

func (v *Vector[T]) Resize(new_size int) {
	if new_size < 0 {
		panic("ERROR: Vector.Resize - negative new size")
	}
	if new_size == v.Size() {
		return
	} else if new_size > v.Size() {
		tmp := make([]T, new_size)
		copy(tmp, v.data)
		v.data = tmp
	} else {
		for v.Size() > new_size {
			v.PopBack()
		}
	}
}

func (v *Vector[T]) Swap(other *Vector[T]) {
	v.data, other.data = other.data, v.data
}

type SortableVector[T constraints.Ordered] struct {
	Vector[T]
}

func NewSortableVector[T constraints.Ordered]() SortableVector[T] {
	return SortableVector[T]{Vector: NewVector[T]()}
}

func NewSortableVectorFromData[T constraints.Ordered](values ...T) SortableVector[T] {
	return SortableVector[T]{Vector: NewVectorFromData(values...)}
}

func NewSortableVectorFromDataRef[T constraints.Ordered](values ...*T) SortableVector[T] {
	v := NewSortableVector[T]()
	for _, val := range values {
		v.PushBackRef(val)
	}
	return v
}

func NewSortableVectorFromVector[T constraints.Ordered](other SortableVector[T]) SortableVector[T] {
	v := NewSortableVector[T]()
	copy(v.data, other.data)
	return v
}

func MakeSortableVector[T constraints.Ordered]() *SortableVector[T] {
	return &SortableVector[T]{Vector: NewVector[T]()}
}

func MakeSortableVectorFromData[T constraints.Ordered](values ...T) *SortableVector[T] {
	return &SortableVector[T]{Vector: NewVectorFromData(values...)}
}

func MakeSortableVectorFromDataRef[T constraints.Ordered](values ...*T) *SortableVector[T] {
	v := MakeSortableVector[T]()
	for _, val := range values {
		v.PushBackRef(val)
	}
	return v
}

func MakeSortableVectorFromVector[T constraints.Ordered](other SortableVector[T]) *SortableVector[T] {
	v := MakeSortableVector[T]()
	copy(v.data, other.data)
	return v
}

func (v *SortableVector[T]) Sort() {
	sort.Slice(v.data, func(i, j int) bool { return v.data[i] < v.data[j] })
}

func (v *SortableVector[T]) StableSort() {
	sort.SliceStable(v.data, func(i, j int) bool { return v.data[i] < v.data[j] })
}

func (v *SortableVector[T]) SortFunc(f func(left *T, right *T) bool) {
	sort.Slice(v.data, func(i, j int) bool { return f(&v.data[i], &v.data[j]) })
}

func (v *SortableVector[T]) StableSortFunc(f func(left *T, right *T) bool) {
	sort.SliceStable(v.data, func(i, j int) bool { return f(&v.data[i], &v.data[j]) })
}

func (v *SortableVector[T]) IsSorted() bool {
	return sort.SliceIsSorted(v.data, func(i, j int) bool { return v.data[i] < v.data[j] })
}

func (v *SortableVector[T]) IsSortedFunc(f func(left *T, right *T) bool) bool {
	return sort.SliceIsSorted(v.data, func(i, j int) bool { return f(&v.data[i], &v.data[j]) })
}

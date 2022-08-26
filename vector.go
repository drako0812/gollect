package gollect

import (
	"fmt"
	"runtime"
	"sort"
	"strings"
	"sync"

	"golang.org/x/exp/constraints"
)

// Vector is a general-purpose collection that can be used for any type.
//
// It does have special functionality for element types that implement the Destructible
// interface.
//
// Note, if you want a Vector of a native type, you will most likely get significantly
// better performance out of an NVector.
type Vector[T any] struct {
	data []T
}

// NewVector creates a new empty Vector, by value
func NewVector[T any]() Vector[T] {
	return Vector[T]{data: []T{}}
}

// NewVectorFromData creates a new Vector using the elements in values, by value.
func NewVectorFromData[T any](values ...T) Vector[T] {
	return Vector[T]{data: values}
}

// NewVectorFromDataRef creates a new Vector using pointers to the elements in values, by value.
func NewVectorFromDataRef[T any](values ...*T) Vector[T] {
	v := NewVector[T]()
	for _, val := range values {
		v.PushBackRef(val)
	}
	return v
}

// NewVectorFromVector creates a new Vector using the values of another, by value.
func NewVectorFromVector[T any](other Vector[T]) Vector[T] {
	v := NewVectorFromData(other.data...)
	//copy(v.data, other.data)
	return v
}

// MakeVector creates a new empty Vector instance.
func MakeVector[T any]() *Vector[T] {
	return &Vector[T]{data: []T{}}
}

// MakeVectorFromData creates a new Vector instance using the elements in values.
func MakeVectorFromData[T any](values ...T) *Vector[T] {
	return &Vector[T]{data: values}
}

// MakeVectorFromDataRef creates a new Vector instance using pointers to the elements in values.
func MakeVectorFromDataRef[T any](values ...*T) *Vector[T] {
	v := MakeVector[T]()
	for _, val := range values {
		v.PushBackRef(val)
	}
	return v
}

// MakeVectorFromVector creates a new Vector instance using the values of another.
func MakeVectorFromVector[T any](other Vector[T]) *Vector[T] {
	v := MakeVectorFromData(other.data...)
	return v
}

// At gets the element at index by value.
//
// Note, this function does no bounds checking besides what the Go runtime does already.
func (v *Vector[T]) At(index int) T {
	return v.data[index]
}

// SafeAt gets the element at index by value.
//
// Note, this function does do bounds checking.
func (v *Vector[T]) SafeAt(index int) T {
	if len(v.data) > 0 {
		if (index >= 0) && (index < len(v.data)) {
			return v.data[index]
		}
		panic("ERROR: Vector.SafeAt - index out of range")
	}
	panic("ERROR: Vector.SafeAt - empty vector")
}

// AtRef gets a pointer to the element at index.
//
// Note, this function does no bounds checking besides what the Go runtime does already.
func (v *Vector[T]) AtRef(index int) *T {
	return &v.data[index]
}

// SafeAtRef gets a pointer to the element at index.
//
// Note, this function does do bounds checking.
func (v *Vector[T]) SafeAtRef(index int) *T {
	if len(v.data) > 0 {
		if (index >= 0) && (index < len(v.data)) {
			return &v.data[index]
		}
		panic("ERROR: Vector.SafeAtRef - index out of range")
	}
	panic("ERROR: Vector.SafeAtRef - empty vector")
}

// Front gets the element at the front of the Vector by value.
func (v *Vector[T]) Front() T {
	if len(v.data) > 0 {
		return v.data[0]
	}
	panic("ERROR: Vector.Front - empty vector")
}

// FrontRef gets a pointer to the element at the front of the Vector.
func (v *Vector[T]) FrontRef() *T {
	if len(v.data) > 0 {
		return &v.data[0]
	}
	panic("ERROR: Vector.FrontRef - empty vector")
}

// Back gets the element at the back of the Vector by value.
func (v *Vector[T]) Back() T {
	if len(v.data) > 0 {
		return v.data[len(v.data)-1]
	}
	panic("ERROR: Vector.Back - empty vector")
}

// BackRef gets a pointer to the element at the back of the Vector.
func (v *Vector[T]) BackRef() *T {
	if len(v.data) > 0 {
		return &v.data[len(v.data)-1]
	}
	panic("ERROR: Vector.BackRef - empty vector")
}

// Data gets the underlying slice of the elements.
func (v *Vector[T]) Data() []T {
	return v.data
}

// IsEmpty returns true if the Vector is empty.
func (v *Vector[T]) IsEmpty() bool {
	return len(v.data) == 0
}

// Size returns the number of elements in the Vector.
func (v *Vector[T]) Size() int {
	return len(v.data)
}

// Clear removes all the elements from the Vector.
//
// If the elements implement the Destructible interface, then they will have the Destruct method called on them.
func (v *Vector[T]) Clear() {
	for !v.IsEmpty() {
		v.PopBack()
	}
}

// Insert adds an element at the specified index, moving all later elements one further index back.
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

// Insert adds an element at the specified index, moving all later elements one further index back.
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

// Erase removes an element at the specified index, moving all later elements one index forward.
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

// PushBack adds an element to the back of the Vector.
func (v *Vector[T]) PushBack(value T) {
	v.data = append(v.data, value)
}

// PushBackRef adds an element to the back of the Vector.
func (v *Vector[T]) PushBackRef(value *T) {
	v.data = append(v.data, *value)
}

// PushFront adds an element to the front of the Vector, moving all later elements one index backward.
func (v *Vector[T]) PushFront(value T) {
	v.data = append([]T{value}, v.data...)
}

// PushFrontRef adds an element to the front of the Vector, moving all later elements one index backward.
func (v *Vector[T]) PushFrontRef(value *T) {
	v.data = append([]T{*value}, v.data...)
}

// PopBack removes an element from the back of the Vector.
//
// If the element implements the Destructible interface, it will have the Destruct method called on it.
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

// PopFront removes an element from the front of the Vector, moving all later elements one index forward.
//
// If the element implements the Destructible interface, it will have the Destruct method called on it.
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

// Resize resizes the Vector.
//
// If the size increases, the new elements are zero-valued.
//
// If the size decreases and the removed elements implement the Destructible interface, they
// will have the Destruct method called on them.
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

// Swap swaps the data of two Vectors.
func (v *Vector[T]) Swap(other *Vector[T]) {
	v.data, other.data = other.data, v.data
}

// String returns a string representation of the Vector and it's contents.
func (v *Vector[T]) String() string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "{")
	for idx, val := range v.data {
		if idx == v.Size()-1 {
			fmt.Fprintf(&builder, "%v", val)
		} else {
			fmt.Fprintf(&builder, "%v, ", val)
		}
	}
	fmt.Fprintf(&builder, "}")
	return builder.String()
}

// Visit calls a function for every element in the Vector.
func (v *Vector[T]) Visit(visitor CollectionVisitor[T]) {
	if v.IsEmpty() {
		return
	}
	break_out := false
	idx := 0
	for (!break_out) && (idx < v.Size()) {
		visitor(v.AtRef(idx), &break_out)
		idx++
	}
}

// Visit calls a function for every element in the Vector in reverse order.
func (v *Vector[T]) VisitReverse(visitor CollectionVisitor[T]) {
	if v.IsEmpty() {
		return
	}
	break_out := false
	idx := v.Size() - 1
	for (!break_out) && (idx >= 0) {
		visitor(v.AtRef(idx), &break_out)
		idx--
	}
}

// ContainsValue returns true if the Vector contains value.
func (v *Vector[T]) ContainsValue(value T) bool {
	if v.IsEmpty() {
		return false
	}
	if _, isEqComparable := interface{}(v.FrontRef()).(EqualityComparable[T]); !isEqComparable {
		return false
	}
	ret := false
	idx := 0
	v.Visit(func(vv *T, break_out *bool) {
		vvv, _ := interface{}(vv).(EqualityComparable[T])
		if vvv.Equal(value) {
			ret = true
			*break_out = true
		}
		idx++
	})
	return ret
}

// ContainsRef returns true if the Vector contains the exact reference of value.
func (v *Vector[T]) ContainsRef(value *T) bool {
	if v.IsEmpty() {
		return false
	}
	ret := false
	idx := 0
	v.Visit(func(vv *T, break_out *bool) {
		if vv == value {
			ret = true
			*break_out = true
		}
		idx++
	})
	return ret
}

// OrderedSearch searches for a value, and returns the index to the first match.
func (v *Vector[T]) OrderedSearch(value T) (found bool, index int) {
	if v.IsEmpty() {
		return false, -1
	}
	if _, isEqComparable := interface{}(v.FrontRef()).(EqualityComparable[T]); !isEqComparable {
		return false, -1
	}
	found = false
	index = 0
	v.Visit(func(vv *T, break_out *bool) {
		vvv, _ := interface{}(vv).(EqualityComparable[T])
		if vvv.Equal(value) {
			found = true
			*break_out = true
		} else {
			index++
		}
	})
	return
}

// OrderedRefSearch searches for an instance, and returns the index to the first match.
func (v *Vector[T]) OrderedRefSearch(value *T) (found bool, index int) {
	if v.IsEmpty() {
		return false, -1
	}
	found = false
	index = 0
	v.Visit(func(vv *T, break_out *bool) {
		if vv == value {
			found = true
			*break_out = true
		} else {
			index++
		}
	})
	return
}

// OrderedSearchRef searches for a value, and returns a pointer to the first match.
func (v *Vector[T]) OrderedSearchRef(value T) *T {
	if v.IsEmpty() {
		return nil
	}
	if _, isEqComparable := interface{}(v.FrontRef()).(EqualityComparable[T]); !isEqComparable {
		return nil
	}
	var ret *T = nil
	index := 0
	v.Visit(func(vv *T, break_out *bool) {
		vvv, _ := interface{}(vv).(EqualityComparable[T])
		if vvv.Equal(value) {
			ret = vv
			*break_out = true
		}
		index++
	})
	return ret
}

// OrderedRefSearchRef searches for an instance, and returns a pointer to the first match.
func (v *Vector[T]) OrderedRefSearchRef(value *T) *T {
	if v.IsEmpty() {
		return nil
	}
	var ret *T = nil
	index := 0
	v.Visit(func(vv *T, break_out *bool) {
		if vv == value {
			ret = vv
			*break_out = true
		}
		index++
	})
	return ret
}

// ChunkMultiplier influences how large a Vector must be before Searches are performed in parallel.
var ChunkMultiplier = 4 // TODO: try to fine-tune this magic number a little bit

// Search searches for a value, and returns an index to a match.
//
// Note, this method may perform a search in parallel, so the match might not be the first
// in the Vector.
func (v *Vector[T]) Search(value T) (found bool, index int) {
	found = false
	index = -1

	if v.IsEmpty() {
		return
	}

	if _, isEqComparable := interface{}(v.FrontRef()).(EqualityComparable[T]); !isEqComparable {
		return
	}

	chunks := runtime.GOMAXPROCS(0)
	if v.Size() < chunks*ChunkMultiplier {
		chunks = 1
	}
	chunk_size := v.Size() / chunks
	chunk_rem := v.Size() % chunks
	last_chunk_num := chunks - 1

	var waitGrp sync.WaitGroup
	var retMtx sync.Mutex

	for i := 0; i < chunks; i++ {
		waitGrp.Add(1)
		go func(ii int) {
			defer waitGrp.Done()
			var my_chunk_size int = chunk_size
			if ii == last_chunk_num {
				my_chunk_size += chunk_rem
			}

			for index2 := chunk_size * ii; index2 < (chunk_size*ii)+my_chunk_size; index2++ {
				vv, _ := interface{}(v.AtRef(index2)).(EqualityComparable[T])
				if vv.Equal(value) {
					retMtx.Lock()
					found = true
					index = index2
					retMtx.Unlock()
					return
				} else {
					//retMtx.Lock()
					if found {
						//	retMtx.Unlock()
						return
					}
					//retMtx.Unlock()
				}
			}
		}(i)
	}
	waitGrp.Wait()

	return
}

// RefSearch searches for an instance, and returns an index to a match.
//
// Note, this method may perform a search in parallel, so the match might not be the first
// in the Vector.
func (v *Vector[T]) RefSearch(value *T) (found bool, index int) {
	found = false
	index = -1

	if v.IsEmpty() {
		return
	}

	chunks := runtime.GOMAXPROCS(0)
	if v.Size() < chunks*ChunkMultiplier {
		chunks = 1
	}
	chunk_size := v.Size() / chunks
	chunk_rem := v.Size() % chunks
	last_chunk_num := chunks - 1

	var waitGrp sync.WaitGroup
	var retMtx sync.Mutex

	for i := 0; i < chunks; i++ {
		waitGrp.Add(1)
		go func(ii int) {
			defer waitGrp.Done()
			var my_chunk_size int = chunk_size
			if ii == last_chunk_num {
				my_chunk_size += chunk_rem
			}

			for index2 := chunk_size * ii; index2 < (chunk_size*ii)+my_chunk_size; index2++ {
				if v.AtRef(index2) == value {
					retMtx.Lock()
					found = true
					index = index2
					retMtx.Unlock()
					return
				} else {
					//retMtx.Lock()
					if found {
						//	retMtx.Unlock()
						return
					}
					//retMtx.Unlock()
				}
			}
		}(i)
	}
	waitGrp.Wait()

	return
}

// SearchRef searches for a value, and returns a pointer to a match.
//
// Note, this method may perform a search in parallel, so the match might not be the first
// in the Vector.
func (v *Vector[T]) SearchRef(value T) *T {
	var ret *T = nil

	if v.IsEmpty() {
		return ret
	}

	if _, isEqComparable := interface{}(v.FrontRef()).(EqualityComparable[T]); !isEqComparable {
		return ret
	}

	chunks := runtime.GOMAXPROCS(0)
	if v.Size() < chunks*ChunkMultiplier {
		chunks = 1
	}
	chunk_size := v.Size() / chunks
	chunk_rem := v.Size() % chunks
	last_chunk_num := chunks - 1

	var waitGrp sync.WaitGroup
	var retMtx sync.Mutex

	for i := 0; i < chunks; i++ {
		waitGrp.Add(1)
		go func(ii int) {
			defer waitGrp.Done()
			var my_chunk_size int = chunk_size
			if ii == last_chunk_num {
				my_chunk_size += chunk_rem
			}

			for index2 := chunk_size * ii; index2 < (chunk_size*ii)+my_chunk_size; index2++ {
				vv, _ := interface{}(v.AtRef(index2)).(EqualityComparable[T])
				if vv.Equal(value) {
					retMtx.Lock()
					ret = v.AtRef(index2)
					retMtx.Unlock()
					return
				} else {
					//retMtx.Lock()
					if ret != nil {
						//	retMtx.Unlock()
						return
					}
					//retMtx.Unlock()
				}
			}
		}(i)
	}
	waitGrp.Wait()

	return ret
}

// SearchRef searches for an instance, and returns a pointer to a match.
//
// Note, this method may perform a search in parallel, so the match might not be the first
// in the Vector.
func (v *Vector[T]) RefSearchRef(value *T) *T {
	var ret *T = nil

	if v.IsEmpty() {
		return ret
	}

	chunks := runtime.GOMAXPROCS(0)
	if v.Size() < chunks*ChunkMultiplier {
		chunks = 1
	}
	chunk_size := v.Size() / chunks
	chunk_rem := v.Size() % chunks
	last_chunk_num := chunks - 1

	var waitGrp sync.WaitGroup
	var retMtx sync.Mutex

	for i := 0; i < chunks; i++ {
		waitGrp.Add(1)
		go func(ii int) {
			defer waitGrp.Done()
			var my_chunk_size int = chunk_size
			if ii == last_chunk_num {
				my_chunk_size += chunk_rem
			}

			for index2 := chunk_size * ii; index2 < (chunk_size*ii)+my_chunk_size; index2++ {
				if v.AtRef(index2) == value {
					retMtx.Lock()
					ret = v.AtRef(index2)
					retMtx.Unlock()
					return
				} else {
					//retMtx.Lock()
					if ret != nil {
						//	retMtx.Unlock()
						return
					}
					//retMtx.Unlock()
				}
			}
		}(i)
	}
	waitGrp.Wait()

	return ret
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

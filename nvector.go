package gollect

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
)

type NVector[T NativeEquatable] struct {
	data []T
}

func NewNVector[T NativeEquatable]() NVector[T] {
	return NVector[T]{data: []T{}}
}

func NewNVectorFromData[T NativeEquatable](values ...T) NVector[T] {
	return NVector[T]{data: values}
}

func NewNVectorFromNVector[T NativeEquatable](other NVector[T]) NVector[T] {
	return NewNVectorFromData(other.data...)
}

func NewNVectorFromVector[T NativeEquatable](other Vector[T]) NVector[T] {
	return NewNVectorFromData(other.data...)
}

func NewVectorFromNVector[T NativeEquatable](other NVector[T]) Vector[T] {
	return NewVectorFromData(other.data...)
}

func (v *NVector[T]) At(index int) T {
	return v.data[index]
}

func (v *NVector[T]) SafeAt(index int) T {
	if !v.IsEmpty() {
		if (index >= 0) && (index < v.Size()) {
			return v.At(index)
		}
		panic("ERROR: NVector.SafeAt - index out of range")
	}
	panic("ERROR: NVector.SafeAt - empty vector")
}

func (v *NVector[T]) IsEmpty() bool {
	return v.Size() == 0
}

func (v *NVector[T]) Size() int {
	return len(v.data)
}

func (v *NVector[T]) AtRef(index int) *T {
	return &v.data[index]
}

func (v *NVector[T]) SafeAtRef(index int) *T {
	if !v.IsEmpty() {
		if (index >= 0) && (index < v.Size()) {
			return v.AtRef(index)
		}
		panic("ERROR: NVector.SafeAtRef - index out of range")
	}
	panic("ERROR: NVector.SafeAtRef - empty vector")
}

func (v *NVector[T]) Front() T {
	if !v.IsEmpty() {
		return v.At(0)
	}
	panic("ERROR: NVector.Front - empty vector")
}

func (v *NVector[T]) FrontRef() *T {
	if !v.IsEmpty() {
		return v.AtRef(0)
	}
	panic("ERROR: NVector.FrontRef - empty vector")
}

func (v *NVector[T]) Back() T {
	if !v.IsEmpty() {
		return v.At(v.Size() - 1)
	}
	panic("ERROR: NVector.Back - empty vector")
}

func (v *NVector[T]) BackRef() *T {
	if !v.IsEmpty() {
		return v.AtRef(v.Size() - 1)
	}
	panic("ERROR: NVector.BackRef - empty vector")
}

func (v *NVector[T]) Data() []T {
	return v.data
}

func (v *NVector[T]) Clear() {
	v.data = []T{}
}

func (v *NVector[T]) Insert(index int, value T) {
	if (index < 0) || (index > v.Size()) {
		panic("ERROR: NVector.Insert - index out of bounds")
	}
	if !v.IsEmpty() {
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

func (v *NVector[T]) InsertRef(index int, value *T) {
	v.Insert(index, *value)
}

func (v *NVector[T]) Erase(index int) {
	if !v.IsEmpty() {
		if (index >= 0) && (index < v.Size()) {
			v.data = append(v.data[:index], v.data[index+1:]...)
		} else {
			panic("ERROR: NVector.Erase - index out of bounds")
		}
	} else {
		panic("ERROR: NVector.Erase - empty vector")
	}
}

func (v *NVector[T]) PushBack(value T) {
	v.data = append(v.data, value)
}

func (v *NVector[T]) PushBackRef(value *T) {
	v.PushBack(*value)
}

func (v *NVector[T]) PushFront(value T) {
	v.data = append([]T{value}, v.data...)
}

func (v *NVector[T]) PushFrontRef(value *T) {
	v.PushFront(*value)
}

func (v *NVector[T]) PopBack() {
	if !v.IsEmpty() {
		v.data = v.data[:v.Size()-1]
	} else {
		panic("ERROR: NVector.PopBack - empty vector")
	}
}

func (v *NVector[T]) PopFront() {
	if !v.IsEmpty() {
		v.data = v.data[1:]
	} else {
		panic("ERROR: NVector.PopFront - empty vector")
	}
}

func (v *NVector[T]) Resize(new_size int) {
	if new_size < 0 {
		panic("ERROR: NVector.Resize - negative new size")
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

func (v *NVector[T]) Swap(other *NVector[T]) {
	v.data, other.data = other.data, v.data
}

func (v *NVector[T]) String() string {
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

func (v *NVector[T]) Visit(visitor CollectionVisitor[T]) {
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

func (v *NVector[T]) VisitReverse(visitor CollectionVisitor[T]) {
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

func (v *NVector[T]) ContainsValue(value T) bool {
	if v.IsEmpty() {
		return false
	}
	for _, val := range v.data {
		if val == value {
			return true
		}
	}
	return false
}

func (v *NVector[T]) ContainsRef(value *T) bool {
	if v.IsEmpty() {
		return false
	}
	for i := 0; i < v.Size(); i++ {
		if value == v.AtRef(i) {
			return true
		}
	}
	return false
}

func (v *NVector[T]) OrderedSearch(value T) (found bool, index int) {
	if v.IsEmpty() {
		return false, -1
	}
	for k, val := range v.data {
		if val == value {
			return true, k
		}
	}
	return false, -1
}

func (v *NVector[T]) OrderedRefSearch(value *T) (found bool, index int) {
	if v.IsEmpty() {
		return false, -1
	}
	for i := 0; i < v.Size(); i++ {
		if value == v.AtRef(i) {
			return true, i
		}
	}
	return false, -1
}

func (v *NVector[T]) OrderedSearchRef(value T) *T {
	if v.IsEmpty() {
		return nil
	}
	for i := 0; i < v.Size(); i++ {
		if v.At(i) == value {
			return v.AtRef(i)
		}
	}
	return nil
}

func (v *NVector[T]) OrderedRefSearchRef(value *T) *T {
	if v.IsEmpty() {
		return nil
	}
	for i := 0; i < v.Size(); i++ {
		if v.AtRef(i) == value {
			return v.AtRef(i)
		}
	}
	return nil
}

func (v *NVector[T]) Search(value T) (found bool, index int) {
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
				if v.At(index2) == value {
					retMtx.Lock()
					found = true
					index = index2
					retMtx.Unlock()
					return
				} else {
					if found {
						return
					}
				}
			}
		}(i)
	}
	waitGrp.Wait()

	return
}

func (v *NVector[T]) RefSearch(value *T) (found bool, index int) {
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
					if found {
						return
					}
				}
			}
		}(i)
	}
	waitGrp.Wait()

	return
}

func (v *NVector[T]) SearchRef(value T) (ret *T) {
	ret = nil
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
				if v.At(index2) == value {
					retMtx.Lock()
					ret = v.AtRef(index2)
					retMtx.Unlock()
					return
				} else {
					if ret != nil {
						return
					}
				}
			}
		}(i)
	}
	waitGrp.Wait()

	return
}

func (v *NVector[T]) RefSearchRef(value *T) (ret *T) {
	ret = nil
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
					ret = v.AtRef(index2)
					retMtx.Unlock()
					return
				} else {
					if ret != nil {
						return
					}
				}
			}
		}(i)
	}
	waitGrp.Wait()

	return
}

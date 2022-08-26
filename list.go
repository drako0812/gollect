package gollect

import (
	"fmt"
	"strings"
)

type listNode[T any] struct {
	data T
	prev *listNode[T]
	next *listNode[T]
}

func newListNode[T any]() *listNode[T] {
	return &listNode[T]{prev: nil, next: nil}
}

func newListNodeValue[T any](value T) *listNode[T] {
	return &listNode[T]{data: value, prev: nil, next: nil}
}

type List[T any] struct {
	front *listNode[T]
	back  *listNode[T]
}

func NewList[T any]() List[T] {
	return List[T]{front: nil, back: nil}
}

func NewListFromData[T any](values ...T) List[T] {
	ret := NewList[T]()
	for _, v := range values {
		ret.PushBack(v)
	}
	return ret
}

func NewListFromDataRef[T any](values ...*T) List[T] {
	ret := NewList[T]()
	for _, v := range values {
		ret.PushBackRef(v)
	}
	return ret
}

func NewListFromList[T any](other List[T]) List[T] {
	ret := NewList[T]()
	other.Visit(func(item *T, break_out *bool) {
		ret.PushBack(*item)
	})
	return ret
}

func MakeList[T any]() *List[T] {
	return &List[T]{front: nil, back: nil}
}

func MakeListFromData[T any](values ...T) *List[T] {
	ret := MakeList[T]()
	for _, v := range values {
		ret.PushBack(v)
	}
	return ret
}

func MakeListFromDataRef[T any](values ...*T) *List[T] {
	ret := MakeList[T]()
	for _, v := range values {
		ret.PushBackRef(v)
	}
	return ret
}

func MakeListFromList[T any](other List[T]) *List[T] {
	ret := MakeList[T]()
	other.Visit(func(item *T, break_out *bool) {
		ret.PushBack(*item)
	})
	return ret
}

func (l *List[T]) Front() T {
	if l.IsEmpty() {
		panic("ERROR: List.Front - empty vector")
	}
	return l.front.data
}

func (l *List[T]) FrontRef() *T {
	if l.IsEmpty() {
		panic("ERROR: List.FrontRef - empty vector")
	}
	return &l.front.data
}

func (l *List[T]) Back() T {
	if l.IsEmpty() {
		panic("ERROR: List.Back - empty vector")
	}
	return l.back.data
}

func (l *List[T]) BackRef() *T {
	if l.IsEmpty() {
		panic("ERROR: List.BackRef - empty vector")
	}
	return &l.back.data
}

func (l *List[T]) IsEmpty() bool {
	return l.front == nil
}

func (l *List[T]) Size() int {
	if l.IsEmpty() {
		return 0
	}
	total := 0
	l.Visit(func(item *T, break_out *bool) {
		total++
	})
	return total
}

func (l *List[T]) Clear() {
	for !l.IsEmpty() {
		l.PopBack()
	}
}

func (l *List[T]) Insert(index int, value T) {
	if l.IsEmpty() && index != 0 {
		panic("ERROR: List.Insert - index out of bounds for empty List")
	} else if l.IsEmpty() {
		l.front = newListNodeValue(value)
		l.back = l.front
	} else {
		new_node := newListNodeValue(value)
		idx := 0
		l.visitNode(func(node *listNode[T], break_out *bool) {
			if idx < index {
				idx++
			} else if idx == index {
				prev := node.prev
				next := node.next
				if next == nil {
					prev.next = new_node
					new_node.prev = prev
					new_node.next = node
					node.prev = new_node
					node.next = nil
					l.back = node
				} else {
					prev.next = new_node
					new_node.prev = prev
					new_node.next = node
					node.prev = new_node
					node.next = next
				}
				*break_out = true
			}
		})
		if idx < index {
			panic("ERROR: List.Insert - index out of bounds")
		}
	}
}

func (l *List[T]) InsertRef(index int, value *T) {
	if l.IsEmpty() && index != 0 {
		panic("ERROR: List.Insert - index out of bounds for empty List")
	} else if l.IsEmpty() {
		l.front = newListNodeValue(*value)
		l.back = l.front
	} else {
		new_node := newListNodeValue(*value)
		idx := 0
		l.visitNode(func(node *listNode[T], break_out *bool) {
			if idx < index {
				idx++
			} else if idx == index {
				prev := node.prev
				next := node.next
				if next == nil {
					prev.next = new_node
					new_node.prev = prev
					new_node.next = node
					node.prev = new_node
					node.next = nil
					l.back = node
				} else {
					prev.next = new_node
					new_node.prev = prev
					new_node.next = node
					node.prev = new_node
					node.next = next
				}
				*break_out = true
			}
		})
		if idx < index {
			panic("ERROR: List.Insert - index out of bounds")
		}
	}
}

func (l *List[T]) Erase(index int) {
	if l.IsEmpty() {
		panic("ERROR: List.Erase - empty list")
	} else {
		idx := 0
		l.visitNode(func(node *listNode[T], break_out *bool) {
			if idx < index {
				idx++
			} else if idx == index {
				prev := node.prev
				next := node.next

				if e, isDestructible := interface{}(&node.data).(Destructible); isDestructible {
					e.Destruct()
				}

				prev.next = next
				if next != nil {
					next.prev = prev
				} else {
					l.back = prev
				}
			}
			*break_out = true
		})
		if idx < index {
			panic("ERROR: List.Erase - index out of bounds")
		}
	}
}

func (l *List[T]) PushBack(value T) {
	node := newListNodeValue(value)
	if !l.IsEmpty() {
		l.back.next = node
	} else {
		l.front = node
	}
	node.prev = l.back
	l.back = node
}

func (l *List[T]) PushBackRef(value *T) {
	node := newListNodeValue(*value)
	if !l.IsEmpty() {
		l.back.next = node
	} else {
		l.front = node
	}
	node.prev = l.back
	l.back = node
}

func (l *List[T]) PushFront(value T) {
	node := newListNodeValue(value)
	if !l.IsEmpty() {
		l.front.prev = node
	} else {
		l.back = node
	}
	node.next = l.front
	l.front = node
}

func (l *List[T]) PushFrontRef(value *T) {
	node := newListNodeValue(*value)
	if !l.IsEmpty() {
		l.front.prev = node
	} else {
		l.back = node
	}
	node.next = l.front
	l.front = node
}

func (l *List[T]) PopBack() {
	if l.IsEmpty() {
		panic("ERROR: List.PopBack - empty list")
	}
	if e, isDestructible := interface{}(&l.back.data).(Destructible); isDestructible {
		e.Destruct()
	}
	if l.front == l.back {
		l.front, l.back = nil, nil
	} else {
		l.back.prev.next = nil
		l.back = l.back.prev
	}
}

func (l *List[T]) PopFront() {
	if l.IsEmpty() {
		panic("ERROR: List.PopFront - empty list")
	}
	if e, isDestructible := interface{}(&l.front.data).(Destructible); isDestructible {
		e.Destruct()
	}
	if l.front == l.back {
		l.front, l.back = nil, nil
	} else {
		l.front.next.prev = nil
		l.front = l.front.next
	}
}

func (l *List[T]) Swap(other *List[T]) {
	l.front, l.back, other.front, other.back = other.front, other.back, l.front, l.back
}

type nodeVisitor[T any] func(*listNode[T], *bool)

func (l *List[T]) visitNode(visitor nodeVisitor[T]) {
	if l.IsEmpty() {
		return
	}
	break_out := false
	node := l.front
	for (!break_out) && (node != nil) {
		visitor(node, &break_out)
		node = node.next
	}
}

func (l *List[T]) visitNodeReverse(visitor nodeVisitor[T]) {
	if l.IsEmpty() {
		return
	}
	break_out := false
	node := l.back
	for (!break_out) && (node != nil) {
		visitor(node, &break_out)
		node = node.prev
	}
}

func (l *List[T]) Visit(visitor CollectionVisitor[T]) {
	if l.IsEmpty() {
		return
	}
	break_out := false
	node := l.front
	for (!break_out) && (node != nil) {
		visitor(&node.data, &break_out)
		node = node.next
	}
}

func (l *List[T]) VisitReverse(visitor CollectionVisitor[T]) {
	if l.IsEmpty() {
		return
	}
	break_out := false
	node := l.back
	for (!break_out) && (node != nil) {
		visitor(&node.data, &break_out)
		node = node.prev
	}
}

func (l *List[T]) ContainsValue(value T) bool {
	if l.IsEmpty() {
		return false
	}
	if _, isEqComparable := interface{}(l.FrontRef()).(EqualityComparable[T]); !isEqComparable {
		return false
	}
	ret := false
	idx := 0
	l.Visit(func(vv *T, break_out *bool) {
		vvv, _ := interface{}(vv).(EqualityComparable[T])
		if vvv.Equal(value) {
			ret = true
			*break_out = true
		}
		idx++
	})
	return ret
}

func (l *List[T]) ContainsRef(value *T) bool {
	if l.IsEmpty() {
		return false
	}
	ret := false
	idx := 0
	l.Visit(func(vv *T, break_out *bool) {
		if vv == value {
			ret = true
			*break_out = true
		}
		idx++
	})
	return ret
}

func (l *List[T]) OrderedSearch(value T) (found bool, index int) {
	if l.IsEmpty() {
		return false, -1
	}
	if _, isEqComparable := interface{}(l.FrontRef()).(EqualityComparable[T]); !isEqComparable {
		return false, -1
	}
	found = false
	index = 0
	l.Visit(func(vv *T, break_out *bool) {
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

func (l *List[T]) OrderedRefSearch(value *T) (found bool, index int) {
	if l.IsEmpty() {
		return false, -1
	}
	found = false
	index = 0
	l.Visit(func(vv *T, break_out *bool) {
		if vv == value {
			found = true
			*break_out = true
		} else {
			index++
		}
	})
	return
}

func (l *List[T]) OrderedSearchRef(value T) *T {
	if l.IsEmpty() {
		return nil
	}
	if _, isEqComparable := interface{}(l.FrontRef()).(EqualityComparable[T]); !isEqComparable {
		return nil
	}
	var ret *T = nil
	l.Visit(func(vv *T, break_out *bool) {
		vvv, _ := interface{}(vv).(EqualityComparable[T])
		if vvv.Equal(value) {
			ret = vv
			*break_out = true
		}
	})
	return ret
}

func (l *List[T]) OrderedRefSearchRef(value *T) *T {
	if l.IsEmpty() {
		return nil
	}
	var ret *T = nil
	l.Visit(func(vv *T, break_out *bool) {
		if vv == value {
			ret = vv
			*break_out = true
		}
	})
	return ret
}

func (l *List[T]) Search(value T) (found bool, index int) {
	return l.OrderedSearch(value)
}

func (l *List[T]) RefSearch(value *T) (found bool, index int) {
	return l.OrderedRefSearch(value)
}

func (l *List[T]) SearchRef(value T) *T {
	return l.OrderedSearchRef(value)
}

func (l *List[T]) RefSearchRef(value *T) *T {
	return l.OrderedRefSearchRef(value)
}

func (l *List[T]) String() string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "{")
	l.Visit(func(value *T, break_out *bool) {
		if value == &l.back.data {
			fmt.Fprintf(&builder, "%v", *value)
		} else {
			fmt.Fprintf(&builder, "%v, ", *value)
		}
	})
	fmt.Fprintf(&builder, "}")
	return builder.String()
}

package gollect

import "testing"

func TestNewVectorInt64(t *testing.T) {
	v := NewVector[int64]()
	if sz := v.Size(); sz != 0 {
		t.Fatalf("Vector should be size 0, got %v", sz)
	}
}

func TestNewVectorFloat64(t *testing.T) {
	v := NewVector[float64]()
	if sz := v.Size(); sz != 0 {
		t.Fatalf("Vector should be size 0, got %v", sz)
	}
}

func TestNewVectorString(t *testing.T) {
	v := NewVector[string]()
	if sz := v.Size(); sz != 0 {
		t.Fatalf("Vector should be size 0, got %v", sz)
	}
}

func TestNewVectorFromData(t *testing.T) {
	v := NewVectorFromData[int64](1)
	if sz := v.Size(); sz != 1 {
		t.Fatalf("Vector should be size 1, got %v", sz)
	}
	if i := v.At(0); i != 1 {
		t.Fatalf("Item at index 0 should be 1, got %v", i)
	}
}

func TestNewVectorFromDataRef(t *testing.T) {
	var i int64 = 1
	v := NewVectorFromDataRef(&i)
	if sz := v.Size(); sz != 1 {
		t.Fatalf("Vector should be size 1, got %v", sz)
	}
	if ii := v.At(0); ii != 1 {
		t.Fatalf("Item at index 0 should be %v, got %v", i, ii)
	}
}

func TestNewVectorFromVector(t *testing.T) {
	v1 := NewVectorFromData[int64](1, 2, 3)
	v2 := NewVectorFromVector(v1)
	if v1.Size() != v2.Size() {
		t.Fatalf("Vector 2 should be size %v, got %v", v1.Size(), v2.Size())
	}
	if v1.At(0) != v2.At(0) {
		t.Fatalf("Vector 2[0] should be %v, got %v", v1.At(0), v2.At(0))
	}
}

func TestAt(t *testing.T) {
	v := NewVectorFromData(1, 2, 3)
	if v.At(0) != 1 {
		t.Fatalf("Vector[0] should be 1, got %v", v.At(0))
	}
	if v.At(1) != 2 {
		t.Fatalf("Vector[1] should be 2, got %v", v.At(1))
	}
	if v.At(2) != 3 {
		t.Fatalf("Vector[2] should be 3, got %v", v.At(2))
	}
}

func TestSafeAt(t *testing.T) {
	v := NewVectorFromData(1, 2, 3)
	if v.SafeAt(0) != 1 {
		t.Fatalf("Vector[0] should be 1, got %v", v.At(0))
	}
	if v.SafeAt(1) != 2 {
		t.Fatalf("Vector[1] should be 2, got %v", v.At(1))
	}
	if v.SafeAt(2) != 3 {
		t.Fatalf("Vector[2] should be 3, got %v", v.At(2))
	}
	f1 := func(ret *bool) {
		*ret = false
		defer func() {
			result := recover().(string)
			if result == "ERROR: Vector.SafeAt - index out of range" {
				*ret = true
			}
		}()
		v.SafeAt(-1)
	}
	f2 := func(ret *bool) {
		*ret = false
		defer func() {
			result := recover().(string)
			if result == "ERROR: Vector.SafeAt - index out of range" {
				*ret = true
			}
		}()
		v.SafeAt(3)
	}
	v2 := NewVector[int64]()
	f3 := func(ret *bool) {
		*ret = false
		defer func() {
			result := recover().(string)
			if result == "ERROR: Vector.SafeAt - empty vector" {
				*ret = true
			}
		}()
		v2.SafeAt(0)
	}
	var result2 bool
	f1(&result2)
	if !result2 {
		t.Fatalf("Should have underflowed")
	}
	f2(&result2)
	if !result2 {
		t.Fatalf("Should have overflowed")
	}
	f3(&result2)
	if !result2 {
		t.Fatalf("Should have failed because empty")
	}
}

func TestFront(t *testing.T) {
	v1 := NewVectorFromData(1, 2, 3)
	front := v1.Front()
	if front != 1 {
		t.Fatalf("Vector1.Front should be 1, got %v", front)
	}

	v2 := NewVector[int64]()
	f1 := func(ret *bool) {
		*ret = false
		defer func() {
			result := recover().(string)
			if result == "ERROR: Vector.Front - empty vector" {
				*ret = true
			}
		}()
		v2.Front()
	}
	var result bool
	f1(&result)
	if !result {
		t.Fatalf("Should have failed because empty")
	}
}

func TestBack(t *testing.T) {
	v1 := NewVectorFromData(1, 2, 3)
	back := v1.Back()
	if back != 3 {
		t.Fatalf("Vector1.Back should be 3, got %v", back)
	}

	v2 := NewVector[int64]()
	f1 := func(ret *bool) {
		*ret = false
		defer func() {
			result := recover().(string)
			if result == "ERROR: Vector.Back - empty vector" {
				*ret = true
			}
		}()
		v2.Back()
	}
	var result bool
	f1(&result)
	if !result {
		t.Fatalf("Should have failed because empty")
	}
}

func TestIsEmpty(t *testing.T) {
	v1 := NewVector[int64]()
	if !v1.IsEmpty() {
		t.Fatalf("Should be empty")
	}

	v2 := NewVectorFromData[int64](1, 2, 3)
	if v2.IsEmpty() {
		t.Fatalf("Should be not empty")
	}
}

func TestSize(t *testing.T) {
	v1 := NewVector[int64]()
	if v1.Size() != 0 {
		t.Fatalf("Size should be 0, got %v", v1.Size())
	}

	v2 := NewVectorFromData(1)
	if v2.Size() != 1 {
		t.Fatalf("Size should be 1, got %v", v2.Size())
	}

	v3 := NewVectorFromData(1, 2)
	if v3.Size() != 2 {
		t.Fatalf("Size should be 2, got %v", v3.Size())
	}
}

func TestClear(t *testing.T) {
	v1 := NewVectorFromData(1, 2, 3)
	v1.Clear()
	if v1.Size() != 0 {
		t.Fatalf("Size should be 0, got %v", v1.Size())
	}
}

var Msgs []string = []string{}

type DBool bool

func (db *DBool) Destruct() {
	*db = false
	Msgs = append(Msgs, "DBool.Destruct")
}

func TestClearAdvanced(t *testing.T) {
	Msgs = []string{}
	v1 := NewVectorFromData[DBool](true, true, true)
	v1.Clear()
	if len(Msgs) != 3 {
		t.Fatalf("Destruct method should have been called 3 times, got %v", len(Msgs))
	}
	for _, v := range Msgs {
		if v != "DBool.Destruct" {
			t.Fatalf("Only DBool.Destruct messages should be here!")
		}
	}
	if v1.Size() != 0 {
		t.Fatalf("Size should be 0, got %v", v1.Size())
	}
}

func TestInsertToEmptyNegative(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("The code did not panic, got \"%v\"", r)
		}
	}()

	v1 := NewVector[int64]()
	v1.Insert(-1, 1)
}

func TestInsertToEmptyZero(t *testing.T) {
	v1 := NewVector[int64]()
	v1.Insert(0, 1)
	if v1.Size() != 1 {
		t.Fatalf("Size should be 1, got %v", v1.Size())
	}
	if v1.At(0) != 1 {
		t.Fatalf("Vector[0] should be 1, got %v", v1.At(0))
	}
}

func TestInsertToEmptyPositive(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("The code did not panic, got \"%v\"", r)
		}
	}()

	v1 := NewVector[int64]()
	v1.Insert(1, 1)
}

func TestInsertBeginning(t *testing.T) {
	v1 := NewVectorFromData(1, 2, 3, 4)
	v1.Insert(0, 5)
	if v1.Size() != 5 {
		t.Fatalf("Size should be 5, got %v", v1.Size())
	}
	if v1.At(0) != 5 {
		t.Fatalf("Vector[0] should be 5, got %v", v1.At(0))
	}
	if v1.At(1) != 1 {
		t.Fatalf("Vector[1] should be 1, got %v", v1.At(1))
	}
}

func TestInsertMiddle(t *testing.T) {
	v1 := NewVectorFromData(1, 2, 3, 4)
	v1.Insert(1, 5)
	if v1.Size() != 5 {
		t.Fatalf("Size should be 5, got %v", v1.Size())
	}
	if v1.At(0) != 1 {
		t.Fatalf("Vector[0] should be 1, got %v", v1.At(0))
	}
	if v1.At(1) != 5 {
		t.Fatalf("Vector[1] should be 5, got %v", v1.At(1))
	}
	if v1.At(2) != 2 {
		t.Fatalf("Vector[2] should be 2, got %v", v1.At(2))
	}
}

func TestInsertEnding(t *testing.T) {
	v1 := NewVectorFromData(1, 2, 3, 4)
	v1.Insert(4, 5)
	if v1.Size() != 5 {
		t.Fatalf("Size should be 5, got %v", v1.Size())
	}
	if v1.At(3) != 4 {
		t.Fatalf("Vector[3] should be 4, got %v", v1.At(3))
	}
	if v1.At(4) != 5 {
		t.Fatalf("Vector[4] should be 5, got %v", v1.At(4))
	}
}

func TestEraseEmpty(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("The code did not panic, got \"%v\"", r)
		}
	}()

	v1 := NewVector[int64]()
	v1.Erase(0)
}

func TestEraseNegative(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("The code did not panic, got \"%v\"", r)
		}
	}()

	v1 := NewVectorFromData(1, 2, 3)
	v1.Erase(-1)
}

func TestEraseBeyond(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("The code did not panic, got \"%v\"", r)
		}
	}()

	v1 := NewVectorFromData(1, 2, 3)
	v1.Erase(3)
}

func TestEraseBeginning(t *testing.T) {
	v1 := NewVectorFromData(1, 2, 3)
	v1.Erase(0)

	if v1.Size() != 2 {
		t.Fatalf("Size should be 2, got %v", v1.Size())
	}
	if v1.At(0) != 2 {
		t.Fatalf("Vector[0] should be 2, got %v", v1.At(0))
	}
}

func TestEraseMiddle(t *testing.T) {
	v1 := NewVectorFromData(1, 2, 3)
	v1.Erase(1)

	if v1.Size() != 2 {
		t.Fatalf("Size should be 2, got %v", v1.Size())
	}
	if v1.At(1) != 3 {
		t.Fatalf("Vector[1] should be 3, got %v", v1.At(1))
	}
}

func TestEraseEnd(t *testing.T) {
	v1 := NewVectorFromData(1, 2, 3)
	v1.Erase(2)

	if v1.Size() != 2 {
		t.Fatalf("Size should be 2, got %v", v1.Size())
	}
	if v1.At(1) != 2 {
		t.Fatalf("Vector[1] should be 2, got %v", v1.At(1))
	}
}

func TestPushBack(t *testing.T) {
	v1 := NewVector[int64]()
	v1.PushBack(1)
	if v1.Size() != 1 {
		t.Fatalf("Size should be 1, got %v", v1.Size())
	}
	if v1.At(0) != 1 {
		t.Fatalf("Vector[0] should be 1, got %v", v1.At(0))
	}

	v1.PushBack(2)
	if v1.Size() != 2 {
		t.Fatalf("Size should be 2, got %v", v1.Size())
	}
	if v1.At(1) != 2 {
		t.Fatalf("Vector[1] should be 1, got %v", v1.At(1))
	}
}

func TestPushFront(t *testing.T) {
	v1 := NewVector[int64]()
	v1.PushFront(1)
	if v1.Size() != 1 {
		t.Fatalf("Size should be 1, got %v", v1.Size())
	}
	if v1.At(0) != 1 {
		t.Fatalf("Vector[0] should be 1, got %v", v1.At(0))
	}

	v1.PushFront(2)
	if v1.Size() != 2 {
		t.Fatalf("Size should be 2, got %v", v1.Size())
	}
	if v1.At(0) != 2 {
		t.Fatalf("Vector[0] should be 2, got %v", v1.At(0))
	}
}

func TestPopBack(t *testing.T) {
	v1 := NewVectorFromData(1, 2, 3)
	v1.PopBack()
	if v1.Size() != 2 {
		t.Fatalf("Size should be 2, got %v", v1.Size())
	}
	if (v1.At(0) != 1) || (v1.At(1) != 2) {
		t.Fatalf("Vector should be {1, 2}, got %v", v1.Data())
	}
}

func TestPopBackEmpty(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("The code did not panic, got \"%v\"", r)
		}
	}()

	v1 := NewVector[int64]()
	v1.PopBack()
}

func TestPopFront(t *testing.T) {
	v1 := NewVectorFromData(1, 2, 3)
	v1.PopFront()
	if v1.Size() != 2 {
		t.Fatalf("Size should be 2, got %v", v1.Size())
	}
	if (v1.At(0) != 2) || (v1.At(1) != 3) {
		t.Fatalf("Vector should be {2, 3}, got %v", v1.Data())
	}
}

func TestPopFrontEmpty(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("The code did not panic, got \"%v\"", r)
		}
	}()

	v1 := NewVector[int64]()
	v1.PopFront()
}

func TestResizeNegative(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("The code did not panic, got \"%v\"", r)
		}
	}()

	v1 := NewVectorFromData(1, 2, 3)
	v1.Resize(-1)
}

func TestResizeSameSize(t *testing.T) {
	v1 := NewVectorFromData(1, 2, 3)
	v1.Resize(3)
	if v1.Size() != 3 {
		t.Fatalf("Size should be 3, got %v", v1.Size())
	}
}

func TestResizeLarger(t *testing.T) {
	v1 := NewVectorFromData(1, 2, 3)
	v1.Resize(4)
	if v1.Size() != 4 {
		t.Fatalf("Size should be 4, got %v", v1.Size())
	}
	if v1.At(3) != 0 {
		t.Fatalf("Vector[3] should be 0, got %v", v1.At(3))
	}
}

func TestResizeSmaller(t *testing.T) {
	v1 := NewVectorFromData(1, 2, 3)
	v1.Resize(2)
	if v1.Size() != 2 {
		t.Fatalf("Size should be 2, got %v", v1.Size())
	}
	if v1.At(1) != 2 {
		t.Fatalf("Vector[1] should be 2, got %v", v1.At(1))
	}
}

func TestResizeFromZero(t *testing.T) {
	v1 := NewVector[int64]()
	v1.Resize(3)
	if v1.Size() != 3 {
		t.Fatalf("Size should be 3, got %v", v1.Size())
	}
	if v1.At(0) != 0 {
		t.Fatalf("Vector[0] should be 0, got %v", v1.At(0))
	}
	if v1.At(2) != 0 {
		t.Fatalf("Vector[2] should be 0, got %v", v1.At(2))
	}
}

func TestResizeToZero(t *testing.T) {
	v1 := NewVectorFromData(1, 2, 3)
	v1.Resize(0)
	if v1.Size() != 0 {
		t.Fatalf("Size should be 0, got %v", v1.Size())
	}
}

func TestSwap(t *testing.T) {
	v1 := NewVectorFromData(1, 2, 3)
	v2 := NewVectorFromData(4, 5, 6)
	v1.Swap(&v2)
	if !((v1.At(0) == 4) && (v1.At(1) == 5) && (v1.At(2) == 6) && (v2.At(0) == 1) && (v2.At(1) == 2) && (v2.At(2) == 3)) {
		t.Fatalf("v1 should be {4, 5, 6} and v2 should be {1, 2, 3}, got %v and %v", v1, v2)
	}
}

func TestSort(t *testing.T) {
	v1 := NewSortableVectorFromData[int64](3, 2, 1)
	v1.Sort()

	if !((v1.At(0) == 1) && (v1.At(1) == 2) && (v1.At(2) == 3)) {
		t.Fatalf("v1 should be {1, 2, 3}, got %v", v1.Data())
	}
}

func TestStableSort(t *testing.T) {
	v1 := NewSortableVectorFromData[int64](3, 2, 1)
	v1.StableSort()

	if !((v1.At(0) == 1) && (v1.At(1) == 2) && (v1.At(2) == 3)) {
		t.Fatalf("v1 should be {1, 2, 3}, got %v", v1.Data())
	}
}

func TestSortFunc(t *testing.T) {
	v1 := NewSortableVectorFromData[int64](1, 2, 3)
	v1.SortFunc(func(left *int64, right *int64) bool { return *left > *right })

	if !((v1.At(0) == 3) && (v1.At(1) == 2) && (v1.At(2) == 1)) {
		t.Fatalf("v1 should be {3, 2, 1}, got %v", v1.Data())
	}
}

func TestStableSortFunc(t *testing.T) {
	v1 := NewSortableVectorFromData[int64](1, 2, 3)
	v1.StableSortFunc(func(left *int64, right *int64) bool { return *left > *right })

	if !((v1.At(0) == 3) && (v1.At(1) == 2) && (v1.At(2) == 1)) {
		t.Fatalf("v1 should be {3, 2, 1}, got %v", v1.Data())
	}
}

func TestIsSorted(t *testing.T) {
	v1 := NewSortableVectorFromData[int64](1, 2, 3)
	v2 := NewSortableVectorFromData[int64](3, 2, 1)

	if !v1.IsSorted() {
		t.Fatalf("v1 should be sorted")
	}

	if v2.IsSorted() {
		t.Fatalf("v2 should not be sorted")
	}
}

func TestIsSortedFunc(t *testing.T) {
	f := func(left *int64, right *int64) bool {
		return *left > *right
	}
	v1 := NewSortableVectorFromData[int64](1, 2, 3)
	v2 := NewSortableVectorFromData[int64](3, 2, 1)

	if v1.IsSortedFunc(f) {
		t.Fatalf("v1 should not be sorted")
	}

	if !v2.IsSortedFunc(f) {
		t.Fatalf("v2 should be sorted")
	}
}

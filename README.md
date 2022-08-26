# Gollect

*Gollect* is a simple collections/data structure library with functionality similar to the STL collections from C++. It utilizes *Go's* new generics functionality.

Documentation can be found [HERE](https://pkg.go.dev/github.com/drako0812/gollect).

## Contents

This package currently implements the following collections.

### Vector

A resizable array type, functionally a wrapper around a *Go* slice.

```go
my_vector := NewVectorFromData[int64](1, 2, 3, 4, 5)
```

### SortableVector

A `Vector` with the ability to sort it's elements. Requires elements satisfy the `constraints.Ordered` interface (from [golang.org/x/exp/constraints](golang.org/x/exp/constraints))

### Deque

A `Deque` is a double-ended queue.

### Queue

A `Queue` is a single-ended queue.

### Stack

A `Stack` is a FIFO stack.

### Destructible

Elements of the collections included in this package can implement the `Destructible` interface which allows the collections to call `Destruct()` on the elements when they are removed from the collections.

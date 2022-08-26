package gollect

type EqualityComparable[T any] interface {
	Equal(other T) bool
	NotEqual(other T) bool
}

type Comparable[T any] interface {
	EqualityComparable[T]
	LesserThan(other T) bool
	GreaterThan(other T) bool
	LesserThanOrEqual(other T) bool
	GreaterThanOrEqual(other T) bool
}

type Int int
type Int8 int8
type Int16 int16
type Int32 int32
type Int64 int64
type Uint uint
type Uint8 uint8
type Uint16 uint16
type Uint32 uint32
type Uint64 uint64
type UintPtr uintptr
type Float32 float32
type Float64 float64
type Complex64 complex64
type Complex128 complex128
type Byte byte
type Rune rune
type String string

func (i *Int) Equal(other Int) bool              { return *i == other }
func (i *Int) NotEqual(other Int) bool           { return *i != other }
func (i *Int) LesserThan(other Int) bool         { return *i < other }
func (i *Int) GreaterThan(other Int) bool        { return *i > other }
func (i *Int) LesserThanOrEqual(other Int) bool  { return *i <= other }
func (i *Int) GreaterThanOrEqual(other Int) bool { return *i >= other }

func (i *Int8) Equal(other Int8) bool              { return *i == other }
func (i *Int8) NotEqual(other Int8) bool           { return *i != other }
func (i *Int8) LesserThan(other Int8) bool         { return *i < other }
func (i *Int8) GreaterThan(other Int8) bool        { return *i > other }
func (i *Int8) LesserThanOrEqual(other Int8) bool  { return *i <= other }
func (i *Int8) GreaterThanOrEqual(other Int8) bool { return *i >= other }

func (i *Int16) Equal(other Int16) bool              { return *i == other }
func (i *Int16) NotEqual(other Int16) bool           { return *i != other }
func (i *Int16) LesserThan(other Int16) bool         { return *i < other }
func (i *Int16) GreaterThan(other Int16) bool        { return *i > other }
func (i *Int16) LesserThanOrEqual(other Int16) bool  { return *i <= other }
func (i *Int16) GreaterThanOrEqual(other Int16) bool { return *i >= other }

func (i *Int32) Equal(other Int32) bool              { return *i == other }
func (i *Int32) NotEqual(other Int32) bool           { return *i != other }
func (i *Int32) LesserThan(other Int32) bool         { return *i < other }
func (i *Int32) GreaterThan(other Int32) bool        { return *i > other }
func (i *Int32) LesserThanOrEqual(other Int32) bool  { return *i <= other }
func (i *Int32) GreaterThanOrEqual(other Int32) bool { return *i >= other }

func (i *Int64) Equal(other Int64) bool              { return *i == other }
func (i *Int64) NotEqual(other Int64) bool           { return *i != other }
func (i *Int64) LesserThan(other Int64) bool         { return *i < other }
func (i *Int64) GreaterThan(other Int64) bool        { return *i > other }
func (i *Int64) LesserThanOrEqual(other Int64) bool  { return *i <= other }
func (i *Int64) GreaterThanOrEqual(other Int64) bool { return *i >= other }

func (i *Uint) Equal(other Uint) bool              { return *i == other }
func (i *Uint) NotEqual(other Uint) bool           { return *i != other }
func (i *Uint) LesserThan(other Uint) bool         { return *i < other }
func (i *Uint) GreaterThan(other Uint) bool        { return *i > other }
func (i *Uint) LesserThanOrEqual(other Uint) bool  { return *i <= other }
func (i *Uint) GreaterThanOrEqual(other Uint) bool { return *i >= other }

func (i *Uint8) Equal(other Uint8) bool              { return *i == other }
func (i *Uint8) NotEqual(other Uint8) bool           { return *i != other }
func (i *Uint8) LesserThan(other Uint8) bool         { return *i < other }
func (i *Uint8) GreaterThan(other Uint8) bool        { return *i > other }
func (i *Uint8) LesserThanOrEqual(other Uint8) bool  { return *i <= other }
func (i *Uint8) GreaterThanOrEqual(other Uint8) bool { return *i >= other }

func (i *Uint16) Equal(other Uint16) bool              { return *i == other }
func (i *Uint16) NotEqual(other Uint16) bool           { return *i != other }
func (i *Uint16) LesserThan(other Uint16) bool         { return *i < other }
func (i *Uint16) GreaterThan(other Uint16) bool        { return *i > other }
func (i *Uint16) LesserThanOrEqual(other Uint16) bool  { return *i <= other }
func (i *Uint16) GreaterThanOrEqual(other Uint16) bool { return *i >= other }

func (i *Uint32) Equal(other Uint32) bool              { return *i == other }
func (i *Uint32) NotEqual(other Uint32) bool           { return *i != other }
func (i *Uint32) LesserThan(other Uint32) bool         { return *i < other }
func (i *Uint32) GreaterThan(other Uint32) bool        { return *i > other }
func (i *Uint32) LesserThanOrEqual(other Uint32) bool  { return *i <= other }
func (i *Uint32) GreaterThanOrEqual(other Uint32) bool { return *i >= other }

func (i *Uint64) Equal(other Uint64) bool              { return *i == other }
func (i *Uint64) NotEqual(other Uint64) bool           { return *i != other }
func (i *Uint64) LesserThan(other Uint64) bool         { return *i < other }
func (i *Uint64) GreaterThan(other Uint64) bool        { return *i > other }
func (i *Uint64) LesserThanOrEqual(other Uint64) bool  { return *i <= other }
func (i *Uint64) GreaterThanOrEqual(other Uint64) bool { return *i >= other }

func (i *UintPtr) Equal(other UintPtr) bool              { return *i == other }
func (i *UintPtr) NotEqual(other UintPtr) bool           { return *i != other }
func (i *UintPtr) LesserThan(other UintPtr) bool         { return *i < other }
func (i *UintPtr) GreaterThan(other UintPtr) bool        { return *i > other }
func (i *UintPtr) LesserThanOrEqual(other UintPtr) bool  { return *i <= other }
func (i *UintPtr) GreaterThanOrEqual(other UintPtr) bool { return *i >= other }

func (i *Float32) Equal(other Float32) bool              { return *i == other }
func (i *Float32) NotEqual(other Float32) bool           { return *i != other }
func (i *Float32) LesserThan(other Float32) bool         { return *i < other }
func (i *Float32) GreaterThan(other Float32) bool        { return *i > other }
func (i *Float32) LesserThanOrEqual(other Float32) bool  { return *i <= other }
func (i *Float32) GreaterThanOrEqual(other Float32) bool { return *i >= other }

func (i *Float64) Equal(other Float64) bool              { return *i == other }
func (i *Float64) NotEqual(other Float64) bool           { return *i != other }
func (i *Float64) LesserThan(other Float64) bool         { return *i < other }
func (i *Float64) GreaterThan(other Float64) bool        { return *i > other }
func (i *Float64) LesserThanOrEqual(other Float64) bool  { return *i <= other }
func (i *Float64) GreaterThanOrEqual(other Float64) bool { return *i >= other }

func (i *Complex64) Equal(other Complex64) bool    { return *i == other }
func (i *Complex64) NotEqual(other Complex64) bool { return *i != other }

func (i *Complex128) Equal(other Complex128) bool    { return *i == other }
func (i *Complex128) NotEqual(other Complex128) bool { return *i != other }

func (i *Byte) Equal(other Byte) bool              { return *i == other }
func (i *Byte) NotEqual(other Byte) bool           { return *i != other }
func (i *Byte) LesserThan(other Byte) bool         { return *i < other }
func (i *Byte) GreaterThan(other Byte) bool        { return *i > other }
func (i *Byte) LesserThanOrEqual(other Byte) bool  { return *i <= other }
func (i *Byte) GreaterThanOrEqual(other Byte) bool { return *i >= other }

func (i *Rune) Equal(other Rune) bool              { return *i == other }
func (i *Rune) NotEqual(other Rune) bool           { return *i != other }
func (i *Rune) LesserThan(other Rune) bool         { return *i < other }
func (i *Rune) GreaterThan(other Rune) bool        { return *i > other }
func (i *Rune) LesserThanOrEqual(other Rune) bool  { return *i <= other }
func (i *Rune) GreaterThanOrEqual(other Rune) bool { return *i >= other }

func (i *String) Equal(other String) bool              { return *i == other }
func (i *String) NotEqual(other String) bool           { return *i != other }
func (i *String) LesserThan(other String) bool         { return *i < other }
func (i *String) GreaterThan(other String) bool        { return *i > other }
func (i *String) LesserThanOrEqual(other String) bool  { return *i <= other }
func (i *String) GreaterThanOrEqual(other String) bool { return *i >= other }

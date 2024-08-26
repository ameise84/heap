package heap

import "github.com/ameise84/heap/compare"

type ID = uint64

type Heap[T compare.Ordered] interface {
	Push(T) ID
	Update(ID, T) bool
	Peek() (ID, T, error)
	Pop() (T, error)
	Find(ID) (T, bool)
	Remove(ID) (T, bool)
	CleanToSlice() []T
	CleanToMap() map[ID]T
	Range(func(T) bool)
	Len() int
}

type IDHeap[T compare.Ordered] interface {
	Push(ID, T) error
	PushOrUpdate(ID, T)
	Update(ID, T) bool
	Peek() (ID, T, error)
	Pop() (T, error)
	Find(ID) (T, bool)
	Remove(ID) (T, bool)
	CleanToSlice() []T
	CleanToMap() map[ID]T
	Range(func(T) bool)
	Len() int
}

// NewHeapMin 小顶堆
func NewHeapMin[T compare.Ordered](size ...int) Heap[T] {
	return newHeapImpl[T](compare.Smaller, size...)
}

// NewHeapMax 大顶堆
func NewHeapMax[T compare.Ordered](size ...int) Heap[T] {
	return newHeapImpl[T](compare.Larger, size...)
}

// NewIDHeapMin 自定义ID的小顶堆
func NewIDHeapMin[T compare.Ordered](size ...int) IDHeap[T] {
	return newIDHeapImpl[T](compare.Smaller, size...)
}

// NewIDHeapMax 自定义ID的大顶堆
func NewIDHeapMax[T compare.Ordered](size ...int) IDHeap[T] {
	return newIDHeapImpl[T](compare.Larger, size...)
}

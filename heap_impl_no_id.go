package heap

import (
	stdHeap "container/heap"
	"github.com/ameise84/heap/compare"
)

func newHeapImpl[T compare.Ordered](result compare.Result, size ...int) Heap[T] {
	n := defaultSize
	if size != nil && size[0] > 0 {
		n = size[0]
	}
	return &heapImpl[T]{
		base:   newHeap(n, result),
		events: map[ID]*entity{},
	}
}

type heapImpl[T compare.Ordered] struct {
	base   *heap
	events map[ID]*entity
	zero   T
}

func (h *heapImpl[T]) Push(ctx T) ID {
	evt := gPool.Get().(*entity)
	evt.ctx = ctx
	stdHeap.Push(h.base, evt)
	h.events[evt.id] = evt
	return evt.id
}

func (h *heapImpl[T]) Update(id ID, ctx T) bool {
	if evt, ok := h.events[id]; ok {
		evt.ctx = ctx
		stdHeap.Fix(h.base, evt.index)
		return true
	}
	return false
}

func (h *heapImpl[T]) Peek() (ID, T, error) {
	if h.base.Len() == 0 {
		return 0, h.zero, ErrHeapEmpty
	}
	evt := h.base.items[0]
	return evt.id, evt.ctx.(T), nil
}

func (h *heapImpl[T]) Pop() (T, error) {
	if h.base.Len() == 0 {
		return h.zero, ErrHeapEmpty
	}
	evt := stdHeap.Pop(h.base).(*entity)
	delete(h.events, evt.id)
	ctx := evt.ctx
	evt.ctx = nil
	gPool.Put(evt)
	return ctx.(T), nil
}

func (h *heapImpl[T]) Find(id ID) (T, bool) {
	if v, ok := h.events[id]; ok {
		return v.ctx.(T), true
	}
	return h.zero, false
}

func (h *heapImpl[T]) Remove(id ID) (ctx T, ok bool) {
	var evt *entity
	if evt, ok = h.events[id]; ok {
		ctx = evt.ctx.(T)
		stdHeap.Remove(h.base, evt.index)
		evt.ctx = nil
		delete(h.events, id)
		gPool.Put(evt)
	}
	return
}

func (h *heapImpl[T]) CleanToSlice() []T {
	objs := make([]T, 0, h.base.Len())
	for _, evt := range h.base.items {
		objs = append(objs, evt.ctx.(T))
		evt.ctx = nil
		gPool.Put(evt)
	}
	h.base.items = h.base.items[:0]
	h.events = map[ID]*entity{}
	return objs
}

func (h *heapImpl[T]) CleanToMap() map[ID]T {
	objs := make(map[ID]T, h.base.Len())
	for _, evt := range h.base.items {
		objs[evt.id] = evt.ctx.(T)
		evt.ctx = nil
		gPool.Put(evt)
	}
	h.base.items = h.base.items[:0]
	h.events = map[ID]*entity{}
	return objs
}

func (h *heapImpl[T]) Range(f func(T) bool) {
	for _, v := range h.events {
		if !f(v.ctx.(T)) {
			break
		}
	}
}

func (h *heapImpl[T]) Len() int {
	return h.base.Len()
}

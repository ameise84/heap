package heap

import "github.com/ameise84/heap/compare"

func newHeap(size int, r compare.Result) *heap {
	q := &heap{
		items:         make([]*entity, 0, size),
		compareResult: r,
	}
	return q
}

type heap struct {
	items         []*entity
	compareResult compare.Result
}

func (h *heap) Len() int { return len(h.items) }

func (h *heap) Less(i, j int) bool {
	return h.items[i].ctx.Compare(h.items[j].ctx) == h.compareResult
}

func (h *heap) Swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
	h.items[i].index, h.items[j].index = i, j
}

func (h *heap) Push(x any) {
	item := x.(*entity)
	item.index = len(h.items)
	h.items = append(h.items, item)
}

func (h *heap) Pop() any {
	old := h.items
	n := len(old)
	x := old[n-1]
	h.items = old[0 : n-1]
	return x
}

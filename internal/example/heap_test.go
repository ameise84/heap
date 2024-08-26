package example

import (
	"github.com/ameise84/heap"
	"github.com/ameise84/heap/compare"
	"log"
	"math/rand/v2"
	"testing"
)

type item struct {
	val int64
	id  uint64
}

func (i *item) Compare(j compare.Ordered) compare.Result {
	switch k := j.(type) {
	case *item:
		if i.val < k.val {
			return compare.Smaller
		}
		if i.val == k.val {
			return compare.Equal
		}
		return compare.Larger
	}
	panic("compare err type")
}

func TestHeap(t *testing.T) {
	minHeap := heap.NewHeapMin[*item]()
	maxHeap := heap.NewHeapMax[*item]()
	var is []*item
	n := uint64(10)
	for i := uint64(0); i < n; i++ {
		is = append(is, &item{
			val: rand.Int64N(1000),
			id:  i,
		})
	}

	for i := uint64(0); i < n; i++ {
		minHeap.Push(is[i])
		maxHeap.Push(is[i])
	}
	for i := uint64(0); i < n; i++ {
		m1, _ := minHeap.Pop()
		m2, _ := maxHeap.Pop()
		log.Println(m1.val, m2.val)
	}

	log.Println("===========================")
	minIDHeap := heap.NewIDHeapMin[*item]()
	maxIDHeap := heap.NewIDHeapMax[*item]()

	for i := uint64(0); i < n; i++ {
		_ = minIDHeap.Push(is[i].id, is[i])
		_ = maxIDHeap.Push(is[i].id, is[i])
	}
	for i := uint64(0); i < n; i++ {
		m1, _ := minIDHeap.Pop()
		m2, _ := maxIDHeap.Pop()
		log.Println(m1.val, m2.val)
	}
}

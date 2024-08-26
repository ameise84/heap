package heap

import (
	"github.com/ameise84/heap/compare"
	"sync"
	"sync/atomic"
)

var (
	gPool sync.Pool
	_gID  atomic.Uint64
)

func init() {
	gPool = sync.Pool{New: func() any {
		return &entity{
			id: _gID.Add(1),
		}
	}}
}

type entity struct {
	id    ID
	index int
	ctx   compare.Ordered
}

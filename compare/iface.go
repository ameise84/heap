package compare

type Result int

const (
	Smaller Result = -1
	Equal   Result = 0
	Larger  Result = 1
)

type Ordered interface {
	Compare(Ordered) Result
}

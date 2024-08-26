package heap

import "errors"

var (
	ErrHeapPush  = errors.New("base can't push ctx with id")
	ErrHeapEmpty = errors.New("base is empty")
)

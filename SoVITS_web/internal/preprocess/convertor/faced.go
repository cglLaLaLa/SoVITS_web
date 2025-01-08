package convertor

import "sync"

var (
	sliceStartConverter *SliceStartConverter
	sliceStartOnce      sync.Once
)

func NewSliceStartConverter() *SliceStartConverter {
	sliceStartOnce.Do(func() {
		sliceStartConverter = new(SliceStartConverter)
	})
	return sliceStartConverter
}

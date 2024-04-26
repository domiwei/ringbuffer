package ringbuffer

import (
	"sync"
)

var _ RingBuffer[int] = (*SimpleRingBuffer[int])(nil)

type SimpleRingBuffer[T any] struct {
	data       []*T
	size       int
	curSize    int
	frontIndex int
	backIndex  int
	mutex      sync.RWMutex
}

func NewSimpleRingBuffer[T any](size int) *SimpleRingBuffer[T] {
	return &SimpleRingBuffer[T]{
		data:       make([]*T, size),
		size:       size,
		curSize:    0,
		frontIndex: 0,
		backIndex:  0,
		mutex:      sync.RWMutex{},
	}
}

func (r *SimpleRingBuffer[T]) Push(value *T) *T {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	// assign nil to oldValue
	var oldValue *T = nil
	if r.curSize == r.size {
		oldValue = r.popNoLock()
	}
	r.data[r.backIndex] = value
	r.backIndex = (r.backIndex + 1) % r.size
	r.curSize++
	return oldValue
}

func (r *SimpleRingBuffer[T]) popNoLock() *T {
	if r.curSize == 0 {
		return nil
	}
	value := r.data[r.frontIndex]
	r.frontIndex = (r.frontIndex + 1) % r.size
	r.curSize--
	return value
}

func (r *SimpleRingBuffer[T]) Pop() *T {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.popNoLock()
}

func (r *SimpleRingBuffer[T]) Front() *T {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	if r.curSize == 0 {
		return nil
	}
	return r.data[r.frontIndex]
}

func (r *SimpleRingBuffer[T]) Back() *T {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	if r.curSize == 0 {
		return nil
	}
	return r.data[(r.backIndex-1+r.size)%r.size]
}

func (r *SimpleRingBuffer[T]) Size() int {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.curSize
}

package ringbuffer

type RingBuffer[T any] interface {
	// Push a new element to the ring buffer, if the buffer is full, the oldest element will be removed and returned
	Push(T) T
	// Pop the oldest element from the ring buffer
	Pop() T
	// Get the oldest element from the ring buffer
	Front() T
	// Get the newest element from the ring buffer
	Back() T
	// Get the size of the ring buffer
	Size() int
	// Iterate the ring buffer
	Iterate(func(T))
}

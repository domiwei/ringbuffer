package ringbuffer

import (
	"log"
	"testing"
)

func TestRingBuffer(t *testing.T) {
	var ring RingBuffer[*int] = NewSimpleRingBuffer[*int](3)
	log.Printf("testcase 1: no element in the ring buffer")
	if ring.Size() != 0 {
		t.Errorf("Size of the ring buffer should be 0, got %d", ring.Size())
	}
	if ring.Front() != nil {
		t.Errorf("Front of the ring buffer should be nil, got %v", *ring.Front())
	}
	if ring.Back() != nil {
		t.Errorf("Back of the ring buffer should be nil, got %v", *ring.Back())
	}
	if ring.Pop() != nil {
		t.Errorf("Pop should return nil, got %v", *ring.Pop())
	}

	log.Printf("testcase 2: push element 1 to the ring buffer")
	if ring.Push(refInt(1)) != nil {
		t.Errorf("Push should return nil, got %v", ring.Push(refInt(1)))
	}
	if ring.Size() != 1 {
		t.Errorf("Size of the ring buffer should be 1, got %d", ring.Size())
	}
	if *ring.Front() != 1 {
		t.Errorf("Front of the ring buffer should be 1, got %v", *ring.Front())
	}
	if *ring.Back() != 1 {
		t.Errorf("Back of the ring buffer should be 1, got %v", *ring.Back())
	}

	log.Printf("testcase 3: push element 2 to the ring buffer")
	if ring.Push(refInt(2)) != nil {
		t.Errorf("Push should return nil, got %v", ring.Push(refInt(2)))
	}
	if ring.Size() != 2 {
		t.Errorf("Size of the ring buffer should be 2, got %d", ring.Size())
	}
	if *ring.Front() != 1 {
		t.Errorf("Front of the ring buffer should be 1, got %v", *ring.Front())
	}
	if *ring.Back() != 2 {
		t.Errorf("Back of the ring buffer should be 2, got %v", *ring.Back())
	}

	log.Printf("testcase 4: push element 3 to the ring buffer")
	if k := ring.Push(refInt(3)); k != nil {
		t.Errorf("Push should return nil, got %v", *k)
	}
	if ring.Size() != 3 {
		t.Errorf("Size of the ring buffer should be 3, got %d", ring.Size())
	}
	if *ring.Front() != 1 {
		t.Errorf("Front of the ring buffer should be 1, got %v", *ring.Front())
	}
	if *ring.Back() != 3 {
		t.Errorf("Back of the ring buffer should be 3, got %v", *ring.Back())
	}

	log.Printf("testcase 5: push element 4 to the ring buffer, element 1 should be removed and returned")
	if k := ring.Push(refInt(4)); *k != 1 {
		t.Errorf("Push should return 1, got %v", *k)
	}
	if ring.Size() != 3 {
		t.Errorf("Size of the ring buffer should be 3, got %d", ring.Size())
	}
	if *ring.Front() != 2 {
		t.Errorf("Front of the ring buffer should be 2, got %v", *ring.Front())
	}
	if *ring.Back() != 4 {
		t.Errorf("Back of the ring buffer should be 4, got %v", *ring.Back())
	}

	log.Printf("testcase 6: pop element 2 from the ring buffer")
	if k := ring.Pop(); *k != 2 {
		t.Errorf("Pop should return 2, got %v", *k)
	}
	if ring.Size() != 2 {
		t.Errorf("Size of the ring buffer should be 2, got %d", ring.Size())
	}
	if *ring.Front() != 3 {
		t.Errorf("Front of the ring buffer should be 3, got %v", *ring.Front())
	}
	if *ring.Back() != 4 {
		t.Errorf("Back of the ring buffer should be 4, got %v", *ring.Back())
	}

	log.Printf("testcase 7: pop element 3 from the ring buffer")
	if k := ring.Pop(); *k != 3 {
		t.Errorf("Pop should return 3, got %v", *k)
	}
	if ring.Size() != 1 {
		t.Errorf("Size of the ring buffer should be 1, got %d", ring.Size())
	}
	if *ring.Front() != 4 {
		t.Errorf("Front of the ring buffer should be 4, got %v", *ring.Front())
	}
	if *ring.Back() != 4 {
		t.Errorf("Back of the ring buffer should be 4, got %v", *ring.Back())
	}

	log.Printf("testcase 8: pop element 4 from the ring buffer, the ring buffer should be empty")
	if k := ring.Pop(); *k != 4 {
		t.Errorf("Pop should return 4, got %v", *k)
	}
	if ring.Size() != 0 {
		t.Errorf("Size of the ring buffer should be 0, got %d", ring.Size())
	}
	if ring.Front() != nil {
		t.Errorf("Front of the ring buffer should be nil, got %v", *ring.Front())
	}
	if ring.Back() != nil {
		t.Errorf("Back of the ring buffer should be nil, got %v", *ring.Back())
	}
	if ring.Pop() != nil {
		t.Errorf("Pop should return nil, got %v", ring.Pop())
	}

	log.Printf("testcase 9: push element 5 to the ring buffer")
	if k := ring.Push(refInt(5)); k != nil {
		t.Errorf("Push should return nil, got %v", *k)
	}
	if ring.Size() != 1 {
		t.Errorf("Size of the ring buffer should be 1, got %d", ring.Size())
	}
	if *ring.Front() != 5 {
		t.Errorf("Front of the ring buffer should be 5, got %v", *ring.Front())
	}
	if *ring.Back() != 5 {
		t.Errorf("Back of the ring buffer should be 5, got %v", *ring.Back())
	}

	log.Printf("testcase 10: push element 6,7 to the ring buffer, ring buffer should be full")
	if k := ring.Push(refInt(6)); k != nil {
		t.Errorf("Push should return nil, got %v", *k)
	}
	if k := ring.Push(refInt(7)); k != nil {
		t.Errorf("Push should return nil, got %v", *k)
	}
	if ring.Size() != 3 {
		t.Errorf("Size of the ring buffer should be 3, got %d", ring.Size())
	}

	log.Printf("testcase 11: push element 8,9,10 to the ring buffer, element 5,6,7 should be removed and returned")
	if k := ring.Push(refInt(8)); *k != 5 {
		t.Errorf("Push should return 5, got %v", *k)
	}
	if k := ring.Push(refInt(9)); *k != 6 {
		t.Errorf("Push should return 6, got %v", *k)
	}
	if k := ring.Push(refInt(10)); *k != 7 {
		t.Errorf("Push should return 7, got %v", *k)
	}
	if ring.Size() != 3 {
		t.Errorf("Size of the ring buffer should be 3, got %d", ring.Size())
	}
	if *ring.Front() != 8 {
		t.Errorf("Front of the ring buffer should be 8, got %v", *ring.Front())
	}
	if *ring.Back() != 10 {
		t.Errorf("Back of the ring buffer should be 10, got %v", *ring.Back())
	}
}

func refInt(i int) *int {
	return &i
}

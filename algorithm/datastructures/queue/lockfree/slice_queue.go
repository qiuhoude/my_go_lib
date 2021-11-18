package lockfree

import "sync"

/*
传统的，我们可以实现一个mutex + slice组成的queue, 在不过分追求性能(时间+空间)的情况下实现一个简单的queue。

*/
// SliceQueue is an unbounded queue which uses a slice as underlying.
type SliceQueue struct {
	data []interface{}
	mu   sync.Mutex
}

// NewSliceQueue returns an empty queue.
// You can give a
func NewSliceQueue(n int) (q *SliceQueue) {
	return &SliceQueue{data: make([]interface{}, n)}
}

// Enqueue puts the given value v at the tail of the queue.
func (q *SliceQueue) Enqueue(v interface{}) {
	q.mu.Lock()
	q.data = append(q.data, v)
	q.mu.Unlock()
}

// Dequeue removes and returns the value at the head of the queue.
// It returns nil if the queue is empty.
func (q *SliceQueue) Dequeue() interface{} {
	q.mu.Lock()
	if len(q.data) == 0 {
		q.mu.Unlock()
		return nil
	}
	v := q.data[0]
	q.data = q.data[1:]
	q.mu.Unlock()
	return v
}
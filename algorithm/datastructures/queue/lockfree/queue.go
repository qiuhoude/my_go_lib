package lockfree

// Queue is a FIFO data structure.
// Enqueue puts a value into its tail,
// Dequeue removes a value from its head.
type Queue interface {
	Enqueue(v interface{})
	Dequeue() interface{}
}

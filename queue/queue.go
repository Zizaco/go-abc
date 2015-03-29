package queue

// Simple queue data structure
// Usage:
//    q := queue.New()
//    q.Push(3)
//    q.Push(6)
//    fmt.Println(q.Length()) // 2
//    fmt.Println(q.Pop())    // 3
//    fmt.Println(q.Pop())    // 6
//    fmt.Println(q.Pop())    // nil
type Queue struct {
	queue []interface{}
}

// Queue constructor
func New() *Queue {
	return new(Queue)
}

// Pushs a new element to the end of the queue
func (q *Queue) Push(v interface{}) {
	q.queue = append(q.queue, v)
}

// Pops out the first element of the queue
func (q *Queue) Pop() interface{} {
	if len(q.queue) < 1 {
		return nil
	}

	first := q.queue[0]
	q.queue = q.queue[1:]

	return first
}

// Returns the length of the queue
func (q *Queue) Length() int {
	return len(q.queue)
}

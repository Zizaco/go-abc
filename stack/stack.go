package stack

// Simple stack data structure
// Usage:
//    s := stack.New()
//    s.Push(3)
//    s.Push(6)
//    fmt.Println(s.Length()) // 2
//    fmt.Println(s.Pop())    // 6
//    fmt.Println(s.Pop())    // 3
//    fmt.Println(s.Pop())    // nil
type Stack struct {
	stack []interface{}
}

// Stack constructor
func New() *Stack {
	return new(Stack)
}

// Pushs a new element as last element of the stack
func (s *Stack) Push(v interface{}) {
	s.stack = append(s.stack, v)
}

// Pops the last element of the stack
func (s *Stack) Pop() interface{} {
	if len(s.stack) < 1 {
		return nil
	}

	last := s.stack[len(s.stack)-1]
	s.stack = s.stack[0 : len(s.stack)-1]

	return last
}

// Returns the length of the stack
func (s *Stack) Length() int {
	return len(s.stack)
}

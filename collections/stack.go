package collections

type StackEmptyError struct {
    
}

func (s StackEmptyError) Error() string {
    return "Stack is empty!"
}

type Stack struct {
    top *Node
}

func (s *Stack) IsEmpty() bool {
    return s.top == nil
}

// Returns the top element of stack. If stack is empty then returns error of type StackEmptyError
func (s *Stack) Peek() (interface{}, error) {
    if !s.IsEmpty() {
        return s.top.Value, nil
    } else {
        return nil, &StackEmptyError{}
    }
}

func (s *Stack) Push(val interface{}) {
    s.top = &Node{Value: val, Prev: s.top}
}

func (s *Stack) Pop() (interface{}, error) {
    if s.IsEmpty() {
        return nil, &StackEmptyError{}
    }

    top := s.top
    s.top = s.top.Prev
    return top.Value, nil
}

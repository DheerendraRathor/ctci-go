package collections

type Stack struct {
    top *Node
    length int
}

func (s *Stack) IsEmpty() bool {
    return s.top == nil
}

// Returns the top element of stack. If stack is empty then returns error of type StackEmptyError
func (s *Stack) Peek() (interface{}, *DataStructureEmptyError) {
    if !s.IsEmpty() {
        return s.top.Value, nil
    } else {
        return nil, &DataStructureEmptyError{}
    }
}

func (s *Stack) Push(val interface{}) {
    s.top = &Node{Value: val, Prev: s.top}
    s.length++
}

func (s *Stack) Pop() (interface{}, *DataStructureEmptyError) {
    if s.IsEmpty() {
        return nil, &DataStructureEmptyError{}
    }

    top := s.top
    s.top = s.top.Prev
    s.length--
    return top.Value, nil
}

func (s *Stack) Length() int {
    return s.length
}

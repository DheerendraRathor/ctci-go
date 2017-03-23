package collections

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestStack_IsEmpty(t *testing.T) {
    stack := new(Stack)
    if !stack.IsEmpty() {
        t.Error("Stack is empty but my code is bad")
    }

    stack.Push(10)
    if stack.IsEmpty() {
        t.Error("Stack is not empty but I can't write code")
    }
    assert.Equal(t, 1, stack.Length())

    stack.Pop()
    if !stack.IsEmpty() {
        t.Error("Stack should be empty cause I popped it")
    }

    assert.Equal(t, 0, stack.Length())
}

func TestStack_Peek(t *testing.T) {
    stack := new(Stack)
    _, err := stack.Peek()
    if err == nil {
        t.Error("I tried to peek empty stack. Didn't get any error")
    }

    stack.Push(10)
    topVal, err := stack.Peek()
    if err != nil {
        t.Error("There is something on the top, but I can't feel it")
    }
    if topVal.(int) != 10 {
        t.Errorf("I'm quite sure I pushed 10, but got %v", topVal)
    }
    assert.Equal(t, 1, stack.Length())

    stack.Pop()
    _, err = stack.Peek()
    if err == nil {
        t.Error("I tried to peek empty stack when I popped something")
    }
    assert.Equal(t, 0, stack.Length())
}

func TestStack_Pop(t *testing.T) {
    stack := new(Stack)
    _, err := stack.Pop()
    if err == nil {
        t.Error("I tried to peek empty stack. Didn't get any error")
    }

    stack.Push(10)
    assert.Equal(t, 1, stack.Length())

    poppedVal, err := stack.Pop()
    if err != nil {
        t.Error("There is something on the top, but I can't feel it")
    }
    if poppedVal.(int) != 10 {
        t.Errorf("I'm quite sure I pushed 10, but got %v", poppedVal)
    }

    assert.Equal(t, 0, stack.Length())
}

func TestStack_Push(t *testing.T) {
    stack := new(Stack)
    assert.Equal(t, 0, stack.Length())

    stack.Push(10)
    topVal, _ := stack.Peek()
    if topVal.(int) != 10 {
        t.Errorf("Pushed 10, got shit %v", topVal)
    }
    assert.Equal(t, 1, stack.Length())

    stack.Push(20)
    topVal, _ = stack.Peek()
    if topVal.(int) != 20 {
        t.Error("Pushed 20, got crap %v", topVal)
    }
    assert.Equal(t, 2, stack.Length())
}

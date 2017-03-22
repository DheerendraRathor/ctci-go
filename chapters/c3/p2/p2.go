package p2

import "github.com/DheerendraRathor/ctci-go/collections"

type MinStack struct {
    mainStack collections.Stack
    minStack collections.Stack
}

func (ms *MinStack) Push(item int) {

    currentMin, err := ms.minStack.Peek()
    if err == nil {
        if item < currentMin.(int) {
            currentMin = item
        }
    } else {        // Error is not nil => Stack is empty
        currentMin = item
    }

    ms.mainStack.Push(item)
    ms.minStack.Push(currentMin)
}

func (ms *MinStack) Pop() (int, *collections.DataStructureEmptyError) {
    ms.minStack.Pop()
    item, err := ms.mainStack.Pop()
    intItem, _ := item.(int)
    return intItem, err
}

func (ms *MinStack) Min() (int, *collections.DataStructureEmptyError) {
    item, err := ms.minStack.Peek()
    intItem, _ := item.(int)
    return intItem, err
}

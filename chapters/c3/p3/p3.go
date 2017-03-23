package p3

import (
    "github.com/DheerendraRathor/ctci-go/collections"
)

type SetOfStacks struct {
    // Stacks of stack
    stacks []*collections.Stack
    capacity int
    currentCount int
}

func InitializeSetOfStack(capacity int) SetOfStacks {
    return SetOfStacks{capacity: capacity, currentCount: 0}
}

func (sos *SetOfStacks) Push(val interface{}) {

    currentStack := &collections.Stack{}
    if sos.currentCount != 0 {
        currentStack = sos.stacks[sos.currentCount - 1]
    }

    if sos.currentCount == 0 || currentStack.Length() >= sos.capacity {
        sos.currentCount++
        sos.stacks = append(sos.stacks, &collections.Stack{})
        currentStack = sos.stacks[sos.currentCount - 1]
    }

    currentStack.Push(val)
}


func (sos *SetOfStacks) Pop() (interface{}, *collections.DataStructureEmptyError) {

    if sos.currentCount == 0 {
        return nil, &collections.DataStructureEmptyError{}
    }

    lastStack := sos.stacks[sos.currentCount - 1]
    val, err := lastStack.Pop()

    if lastStack.Length() == 0 {
        sos.stacks[sos.currentCount - 1] = nil
        sos.stacks = sos.stacks[:sos.currentCount-1]
        sos.currentCount--
    }

    return val, err
}

// Pop at specific sub-stack. Remember counting starts from 1 😆
func (sos *SetOfStacks) PopAt(stack int) (interface{}, *collections.DataStructureEmptyError) {
    if stack > len(sos.stacks) {
        return nil, &collections.DataStructureEmptyError{}
    }

    subStack := sos.stacks[stack - 1]
    val, err := subStack.Pop()

    if subStack.Length() == 0 {
        sos.stacks[stack - 1] = nil
        copy(sos.stacks[stack - 1:], sos.stacks[stack:])
        sos.stacks = sos.stacks[:len(sos.stacks) - 1]
    }

    return val, err
}


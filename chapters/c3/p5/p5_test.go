package p5

import (
    "testing"
    "github.com/DheerendraRathor/ctci-go/collections"
    "github.com/stretchr/testify/assert"
)

func TestSortStack(t *testing.T) {
    var stack *collections.Stack
    SortStack(stack)

    assert.Nil(t, stack)

    stack = new(collections.Stack)

    valuesToPush := []int{4, 2, 5, 1, 10, 3, 11}
    sortedValues := []int{1, 2, 3, 4, 5, 10, 11}

    for _, val := range valuesToPush {
        stack.Push(val)
    }

    SortStack(stack)

    for _, val := range sortedValues {
        top, _ := stack.Pop()
        assert.Equal(t, val, top)
    }

}

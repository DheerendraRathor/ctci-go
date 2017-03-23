package p3

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestSetOfStacks_Push(t *testing.T) {
    sos := InitializeSetOfStack(2)
    sos.Push(1)
    sos.Push(2)
    assert.Equal(t, 1, len(sos.stacks))

    sos.Push(3)
    assert.Equal(t, 2, len(sos.stacks))

    sos.Push(4)
    sos.Push(5)
    sos.Push(6)
    assert.Equal(t, 3, len(sos.stacks))

    sos.Push(7)
    assert.Equal(t, 4, len(sos.stacks))
}

func TestSetOfStacks_Pop(t *testing.T) {
    sos := InitializeSetOfStack(2)
    for i := 1; i < 8; i++ {
        sos.Push(i)
    }

    assert.Equal(t, 4, len(sos.stacks))
    val, err := sos.Pop()
    assert.Equal(t, 3, len(sos.stacks))
    assert.Equal(t, 7, val)
    assert.Nil(t, err)

    sos.Pop()
    sos.Pop()
    val, err = sos.Pop()
    assert.Equal(t, 2, len(sos.stacks))
    assert.Equal(t, 4, val)
    assert.Nil(t, err)

    sos.Push(1)
    assert.Equal(t, 2, len(sos.stacks))

    sos.Push(3)
    assert.Equal(t, 3, len(sos.stacks))

    sos2 := InitializeSetOfStack(2)
    sos2.Push(1)
    sos2.Push(2)

    sos2.Pop()
    val, err = sos2.Pop()
    assert.Equal(t, 0, len(sos2.stacks))
    assert.Equal(t, 1, val)
    assert.Nil(t, err)

    _, err = sos2.Pop()
    assert.Equal(t, 0, len(sos2.stacks))
    assert.NotNil(t, err)
}

func TestSetOfStacks_PopAt(t *testing.T) {
    sos := InitializeSetOfStack(2)
    for i := 1; i < 8; i++ {
        sos.Push(i)
    }

    val, err := sos.PopAt(2)
    assert.Equal(t, 1, sos.stacks[1].Length())
    assert.Nil(t, err)
    assert.Equal(t, 4, val)

    val, err = sos.PopAt(5)
    assert.NotNil(t, err)
    assert.Nil(t, val)

    val, err = sos.PopAt(2)

    // This will be 2 since original stack had length 0 and that must be replaced by next stack
    assert.Equal(t, 2, sos.stacks[1].Length())
    assert.Nil(t, err)
    assert.Equal(t, 3, val)

    assert.Equal(t, 3, len(sos.stacks))
}

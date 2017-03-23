package p4

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestMyQueue_Push(t *testing.T) {
    mq := MyQueue{}

    mq.Push(1)
    mq.Push(2)
    mq.Push(3)
    mq.Push(4)

    assert.Equal(t, 4, mq.forward.Length())
    assert.Equal(t, 0, mq.backward.Length())
}

func TestMyQueue_Pop(t *testing.T) {
    mq := MyQueue{}

    mq.Push(1)
    mq.Push(2)

    assert.Equal(t, 2, mq.forward.Length())

    val, err := mq.Pop()
    assert.Nil(t, err)
    assert.Equal(t, 1, val)

    assert.Equal(t, 0, mq.forward.Length())
    assert.Equal(t, 1, mq.backward.Length())

    mq.Push(3)
    mq.Push(4)

    val, err = mq.Pop()
    assert.Nil(t, err)
    assert.Equal(t, 2, val)

    assert.Equal(t, 2, mq.forward.Length())
    assert.Equal(t, 0, mq.backward.Length())

    val, err = mq.Pop()
    assert.Nil(t, err)
    assert.Equal(t, 3, val)

    assert.Equal(t, 0, mq.forward.Length())
    assert.Equal(t, 1, mq.backward.Length())

    mq.Pop()
    val, err = mq.Pop()
    assert.NotNil(t, err)
}

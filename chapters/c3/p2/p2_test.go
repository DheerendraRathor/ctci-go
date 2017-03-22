package p2

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestMinStack(t *testing.T) {
    ms := new(MinStack)
    ms.Push(10)
    min, _ := ms.Min()
    assert.Equal(t, 10, min)

    ms.Push(20)
    min, _ = ms.Min()
    assert.Equal(t, 10, min)

    ms.Push(5)
    min, _ = ms.Min()
    assert.Equal(t, 5, min)

    ms.Pop()
    min, _ = ms.Min()
    assert.Equal(t, 10, min)

    ms.Pop()
    min, _ = ms.Min()
    assert.Equal(t, 10, min)
}

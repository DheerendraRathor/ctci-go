package collections

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestMinHeap_Push(t *testing.T) {
    mh := MinHeap{}
    mh.Push(5)
    mh.Push(4)
    mh.Push(10)
    mh.Push(2)

    assert.Equal(t, 4, mh.size)
    top, _ := mh.Top()
    assert.Equal(t, 2, top)
    validateSlices(t, []int{2, 4, 10, 5}, mh.array, mh.size)

    mh.Pop()
    mh.Push(3)
    assert.Equal(t, 4, mh.size)
    top, _ = mh.Top()
    assert.Equal(t, 3, top)
    validateSlices(t, []int{3, 4, 10, 5}, mh.array, mh.size)

    mh2 := MinHeap{}
    mh2.Push(5)
    mh2.Push(4)
    mh2.Push(2)
    mh2.Push(10)

    mh2.Pop()
    top, err := mh2.Top()
    assert.Nil(t, err)
    assert.Equal(t, 4, top)
    validateSlices(t, []int{4, 5, 10}, mh2.array, mh2.size)

}

func TestMinHeap_Top(t *testing.T) {
    mh := MinHeap{}
    top, err := mh.Top()
    assert.NotNil(t, err)

    mh.Push(5)
    top, err = mh.Top()
    assert.Nil(t, err)
    assert.Equal(t, 5, top)

    mh.Push(4)
    top, err = mh.Top()
    assert.Nil(t, err)
    assert.Equal(t, 4, top)

    mh.Push(10)
    top, err = mh.Top()
    assert.Nil(t, err)
    assert.Equal(t, 4, top)

    mh.Push(2)
    top, err = mh.Top()
    assert.Nil(t, err)
    assert.Equal(t, 2, top)
}

func TestMinHeap_Pop(t *testing.T) {
    mh := MinHeap{}
    _, err := mh.Pop()
    assert.NotNil(t, err)

    mh.Push(9)
    mh.Push(4)
    mh.Push(5)
    mh.Push(2)

    popped, err := mh.Pop()
    assert.Nil(t, err)
    assert.Equal(t, 2, popped)

    popped, err = mh.Pop()
    assert.Nil(t, err)
    assert.Equal(t, 4, popped)

    popped, err = mh.Pop()
    assert.Nil(t, err)
    assert.Equal(t, 5, popped)

    popped, err = mh.Pop()
    assert.Nil(t, err)
    assert.Equal(t, 9, popped)

    popped, err = mh.Pop()
    assert.NotNil(t, err)

}

func validateSlices(t *testing.T, expected []int, actual []int, size int) {
    for i := 0; i < size; i++ {
        assert.Equal(t, expected[i], actual[i], "Mismatch at index: %d", i)
    }
}

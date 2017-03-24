package collections

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestQueue_IsEmpty(t *testing.T) {
    queue := new(Queue)
    if !queue.IsEmpty() {
        t.Error("IsEmpty failed for emtpy list")
    }
    assert.Equal(t, 0, queue.Length())

    queue.Add(10)
    if queue.IsEmpty() {
        t.Error("IsEmpty failed for non empty list")
    }

    assert.Equal(t, 1, queue.Length())
}

func TestQueue_Add(t *testing.T) {
    queue := new(Queue)
    _, err := queue.Peek()
    if err == nil {
        t.Error("Peek didn't fail")
    }
    assert.Equal(t, 0, queue.Length())

    queue.Add(10)
    assert.Equal(t, 1, queue.Length())

    topVal, err := queue.Peek()
    if err != nil {
        t.Error("Peek failed. Error is not nil")
    }

    if topVal.(int) != 10 {
        t.Error("Peek failed. Top value is not 10")
    }

    queue.Add(20)
    assert.Equal(t, 2, queue.Length())
    topVal, err = queue.Peek()
    if err != nil {
        t.Error("Peek failed. Error is not nil after second Add")
    }

    if topVal.(int) != 10 {
        t.Error("Peek failed. Top value is not 10")
    }

    queue.Remove()
    assert.Equal(t, 1, queue.Length())

    topVal, err = queue.Peek()
    if err != nil {
        t.Error("Peek failed. Error is not nil after Remove")
    }

    if topVal.(int) != 20 {
        t.Error("Peek failed. Top value is not 20")
    }
}

func TestQueue_Remove(t *testing.T) {
    queue := new(Queue)
    _, err := queue.Remove()
    if err == nil {
        t.Error("Remove didn't fail")
    }
    assert.Equal(t, 0, queue.Length())

    queue.Add(10)
    queue.Add(20)
    assert.Equal(t, 2, queue.Length())

    removedVal, err := queue.Remove()
    if err != nil {
        t.Error("Remove failed. Error is not nil after Remove")
    }
    if removedVal.(int) != 10 {
        t.Error("Removed value is not 10")
    }

    removedVal, err = queue.Remove()
    assert.Equal(t, 0, queue.Length())
    if err != nil {
        t.Error("Remove failed. Error is not nil after Remove")
    }
    if removedVal.(int) != 20 {
        t.Error("Removed value is not 20")
    }

    _, err = queue.Remove()
    if err == nil {
        t.Error("Remove didn't fail on supposedly empty Queue")
    }
}

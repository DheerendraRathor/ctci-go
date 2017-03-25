// Min heap implementation for integers. It is not done for `interface{}` type as we can't compare these things
package heaps

import "github.com/DheerendraRathor/ctci-go/collections"

type MinHeap struct {
    array []int
    size int
}

func (mh *MinHeap) Push(val int) {
    if len(mh.array) > mh.size {
        mh.array[mh.size] = val
    } else {
        mh.array = append(mh.array, val)
    }
    mh.size++
    mh.heapifyBottomUp(mh.size - 1)
}

func (mh *MinHeap) Top() (int, *collections.DataStructureEmptyError) {
    if mh.size > 0 {
        return mh.array[0], nil
    } else {
        return 0, &collections.DataStructureEmptyError{}
    }
}

func (mh *MinHeap) Pop() (int, *collections.DataStructureEmptyError) {
    if mh.size <= 0 {
        return 0, &collections.DataStructureEmptyError{}
    }

    top := mh.array[0]
    mh.size--

    if mh.size > 0 {
        mh.array[0] = mh.array[mh.size]
        mh.heapifyTopDown(0)
    }

    return top, nil
}

// Bottom up heapify element at index i
func (mh *MinHeap) heapifyBottomUp(i int) {
    if i == 0 || i >= mh.size{
        return
    }

    parent := (i - 1) / 2
    if mh.array[parent] > mh.array[i] {
        mh.array[parent], mh.array[i] = mh.array[i], mh.array[parent]
        mh.heapifyBottomUp(parent)
    }
}

func (mh *MinHeap) heapifyTopDown(i int) {
    left := 2*i + 1
    right := 2*i + 2
    lowest := i

    if left < mh.size && mh.array[left] < mh.array[i] {
        lowest = left
    }
    if right < mh.size && mh.array[right] < mh.array[lowest] {
        lowest = right
    }

    if lowest != i {
        mh.array[i], mh.array[lowest] = mh.array[lowest], mh.array[i]
        mh.heapifyTopDown(lowest)
    }
}

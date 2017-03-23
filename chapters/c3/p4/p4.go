package p4

import "github.com/DheerendraRathor/ctci-go/collections"

type MyQueue struct {
    forward collections.Stack
    backward collections.Stack
}

func (mq *MyQueue) Push(val interface{}) {
    mq.forward.Push(val)
}

func (mq *MyQueue) Pop() (interface{}, *collections.DataStructureEmptyError) {
    if mq.backward.IsEmpty() {
        for !mq.forward.IsEmpty() {
            val, _ := mq.forward.Pop()
            mq.backward.Push(val)
        }
    }

    return mq.backward.Pop()
}

package p8

import (
    "testing"
    "github.com/DheerendraRathor/ctci-go/chapter_2"
)

func TestDetectCycleInList(t *testing.T) {
    list := chapter_2.SliceToLinkedList([]interface{}{})

    loopNode := DetectCycleInList(*list)
    if loopNode != nil {
        t.Error("Loop detected. Don't know why :/")
    }

    list.AddNodesBySlice([]interface{}{1, 2, 3, 4, 5})
    loopNode = DetectCycleInList(*list)
    if loopNode != nil {
        t.Error("Loop detected. Don't know why :/")
    }

    lastNode := list.Last()

    list.AddNodesBySlice([]interface{}{6, 7, 8, 9})
    list.Last().Next = lastNode
    loopNode = DetectCycleInList(*list)
    if loopNode == nil {
        t.Error("Loop not detected. I'm freaking stupid :-(")
    }

    loopNodeValue := loopNode.Value.(int)
    if loopNodeValue != 5 {
        t.Errorf("Loop node value is not 5. It is %d", loopNodeValue)
    }
}

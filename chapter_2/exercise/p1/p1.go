package p1

import (
    "github.com/DheerendraRathor/ctci-go/chapter_2"
)

func RemoveDuplicates(list *chapter_2.LinkedList)  {
    hashMap := make(map[interface{}]bool)

    current := list.Head
    for current != nil {
        hashMap[current.Value] = true
        current = current.Next
    }

    current = list.Head
    for current != nil {
        // After first detection, we'll make value to false and hence deleting further nodes
        if hashMap[current.Value] {
            hashMap[current.Value] = false
        } else {
            list.DeleteNodeByRef(current)
        }

        current = current.Next
    }
}

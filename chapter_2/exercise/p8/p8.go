package p8

import "github.com/DheerendraRathor/ctci-go/collections"

//Detect if linked list has loop and return the starting point of loop
// If there is no loop then it will return nil
func DetectCycleInList(list collections.LinkedList) *collections.Node {
    slowPtr := list.Head
    fastPtr := list.Head

    steps := 0

    for slowPtr != nil && fastPtr != nil && fastPtr.Next != nil {
        steps++
        slowPtr = slowPtr.Next
        fastPtr = fastPtr.Next.Next

        if slowPtr == fastPtr {
            break
        }
    }

    // No loop in list
    if slowPtr == nil || fastPtr == nil || fastPtr.Next == nil {
        return nil
    }

    slowPtr = list.Head
    for slowPtr != fastPtr {
        slowPtr = slowPtr.Next
        fastPtr = fastPtr.Next
    }

    return slowPtr
}

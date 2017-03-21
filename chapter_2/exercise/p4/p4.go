package p4

import "github.com/DheerendraRathor/ctci-go/collections"

func PartitionOfList(list *collections.LinkedList, pivot int) *collections.LinkedList {
    lessThanList := new(collections.LinkedList)
    greaterThanOrEqualList := new(collections.LinkedList)

    // This variable will store last node in lessThanList. So that we can merge greaterThanOrEqualList directly.
    var lastLessThanNode *collections.Node;

    currentNode := list.Head
    for currentNode != nil {
        intValueOfCurrentNode := currentNode.Value.(int)
        if intValueOfCurrentNode < pivot {
            lastLessThanNode = lessThanList.AddNode(intValueOfCurrentNode)
        } else {
            greaterThanOrEqualList.AddNode(intValueOfCurrentNode)
        }
        currentNode = currentNode.Next
    }

    if lastLessThanNode != nil {
        lastLessThanNode.Next = greaterThanOrEqualList.Head
    }
    return lessThanList
}

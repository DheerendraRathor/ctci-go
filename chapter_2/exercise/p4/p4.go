package p4

import "github.com/DheerendraRathor/ctci-go/chapter_2"

func PartitionOfList(list *chapter_2.LinkedList, pivot int) *chapter_2.LinkedList {
    lessThanList := new(chapter_2.LinkedList)
    greaterThanOrEqualList := new(chapter_2.LinkedList)

    // This variable will store last node in lessThanList. So that we can merge greaterThanOrEqualList directly.
    var lastLessThanNode *chapter_2.Node;

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
